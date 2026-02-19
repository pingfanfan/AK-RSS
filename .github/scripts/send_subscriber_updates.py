#!/usr/bin/env python3
import json
import os
import re
import smtplib
import sys
from dataclasses import dataclass
from datetime import datetime, timezone
from email.message import EmailMessage
from pathlib import Path
from typing import List
from urllib import request, parse

EMAIL_RE = re.compile(r"[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}", re.I)


@dataclass
class Update:
    feed_title: str
    title: str
    link: str
    run_at: datetime
    tldr: str
    action: str


def parse_time(s: str) -> datetime:
    return datetime.fromisoformat(s.replace("Z", "+00:00")).astimezone(timezone.utc)


def load_last_sent(path: Path) -> datetime:
    if not path.exists():
        return datetime(1970, 1, 1, tzinfo=timezone.utc)
    txt = path.read_text(encoding="utf-8").strip()
    if not txt:
        return datetime(1970, 1, 1, tzinfo=timezone.utc)
    return parse_time(txt)


def save_last_sent(path: Path, now: datetime) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(now.astimezone(timezone.utc).isoformat().replace("+00:00", "Z") + "\n", encoding="utf-8")


def fetch_subscribers(repo: str, token: str, api_url: str) -> List[str]:
    owner, name = repo.split("/", 1)
    page = 1
    emails = set()

    while True:
        q = parse.urlencode({"state": "open", "per_page": 100, "page": page})
        url = f"{api_url}/repos/{owner}/{name}/issues?{q}"
        req = request.Request(url)
        req.add_header("Authorization", f"Bearer {token}")
        req.add_header("Accept", "application/vnd.github+json")
        req.add_header("X-GitHub-Api-Version", "2022-11-28")

        with request.urlopen(req, timeout=30) as resp:
            issues = json.loads(resp.read().decode("utf-8"))

        if not issues:
            break

        for issue in issues:
            if "pull_request" in issue:
                continue
            title = (issue.get("title") or "").lower()
            body = issue.get("body") or ""
            if "subscribe" not in title and "email:" not in body.lower():
                continue

            match = EMAIL_RE.search(body)
            if match:
                emails.add(match.group(0).lower())

        page += 1

    return sorted(emails)


def load_updates(path: Path, after: datetime) -> List[Update]:
    if not path.exists():
        return []

    data = json.loads(path.read_text(encoding="utf-8"))
    out: List[Update] = []
    for item in data.get("updates", []):
        run_at_raw = item.get("run_at")
        if not run_at_raw:
            continue
        run_at = parse_time(run_at_raw)
        if run_at <= after:
            continue

        out.append(
            Update(
                feed_title=item.get("feed_title") or "Unknown Feed",
                title=item.get("title") or "(untitled)",
                link=item.get("link") or "",
                run_at=run_at,
                tldr=item.get("tldr") or item.get("what") or "",
                action=item.get("action") or "",
            )
        )

    out.sort(key=lambda x: x.run_at, reverse=True)
    return out


def build_message(from_addr: str, to_addr: str, updates: List[Update]) -> EmailMessage:
    msg = EmailMessage()
    msg["From"] = from_addr
    msg["To"] = to_addr
    msg["Subject"] = f"[AK-RSS] {len(updates)} new updates in the last 2 hours"

    lines = []
    lines.append("AK-RSS 2-hour update")
    lines.append("")
    lines.append("You are receiving this because you subscribed on the AK-RSS page.")
    lines.append("To unsubscribe, close your subscription issue in the repo.")
    lines.append("")

    for i, u in enumerate(updates, start=1):
        lines.append(f"{i}. {u.title}")
        lines.append(f"   Feed: {u.feed_title}")
        if u.tldr:
            lines.append(f"   TLDR: {u.tldr}")
        if u.action:
            lines.append(f"   Action: {u.action}")
        if u.link:
            lines.append(f"   Link: {u.link}")
        lines.append("")

    msg.set_content("\n".join(lines))
    return msg


def send_all(subscribers: List[str], updates: List[Update]) -> int:
    host = os.environ["SMTP_HOST"]
    port = int(os.environ.get("SMTP_PORT", "587"))
    username = os.environ.get("SMTP_USERNAME", "")
    password = os.environ.get("SMTP_PASSWORD", "")
    from_addr = os.environ["SMTP_FROM"]

    if not subscribers or not updates:
        return 0

    sent = 0
    with smtplib.SMTP(host, port, timeout=30) as server:
        server.ehlo()
        server.starttls()
        server.ehlo()
        if username and password:
            server.login(username, password)

        for to_addr in subscribers:
            msg = build_message(from_addr, to_addr, updates)
            server.send_message(msg)
            sent += 1

    return sent


def main() -> int:
    repo = os.environ["GITHUB_REPOSITORY"]
    token = os.environ["GITHUB_TOKEN"]
    api_url = os.environ.get("GITHUB_API_URL", "https://api.github.com")

    updates_path = Path("site/data/updates.json")
    state_path = Path(".opmlwatch/subscriber_last_sent_at.txt")

    now = datetime.now(timezone.utc)
    last_sent = load_last_sent(state_path)
    subscribers = fetch_subscribers(repo, token, api_url)
    updates = load_updates(updates_path, last_sent)

    print(f"subscribers={len(subscribers)}")
    print(f"updates_since_last_send={len(updates)}")

    sent = send_all(subscribers, updates)
    print(f"emails_sent={sent}")

    # Only advance cursor when at least one email was delivered.
    # This avoids empty commits and preserves unsent updates for retry.
    if sent > 0:
        save_last_sent(state_path, now)
        print("cursor_updated=1")
    else:
        print("cursor_updated=0")
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except Exception as exc:
        print(f"subscriber_update_failed: {exc}", file=sys.stderr)
        raise
