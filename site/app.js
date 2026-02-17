const byId = (id) => document.getElementById(id);
const SUBSCRIBE_URL_BASE = "https://github.com/pingfanfan/AK-RSS/issues/new";

const fmt = new Intl.DateTimeFormat("en-GB", {
  year: "numeric",
  month: "short",
  day: "2-digit",
  hour: "2-digit",
  minute: "2-digit",
  hour12: false,
  timeZone: "Europe/London"
});

const shortFmt = new Intl.DateTimeFormat("en-GB", {
  month: "short",
  day: "2-digit",
  hour: "2-digit",
  minute: "2-digit",
  hour12: false,
  timeZone: "Europe/London"
});

async function fetchJSON(path) {
  const res = await fetch(`${path}?t=${Date.now()}`, { cache: "no-store" });
  if (!res.ok) throw new Error(`${path} ${res.status}`);
  return res.json();
}

function statNode(k, v) {
  const card = document.createElement("article");
  card.className = "stat";

  const key = document.createElement("p");
  key.className = "stat-k";
  key.textContent = k;

  const val = document.createElement("p");
  val.className = "stat-v";
  val.textContent = String(v);

  card.appendChild(key);
  card.appendChild(val);
  return card;
}

function renderStats(updates, daily) {
  const el = byId("stats");
  el.innerHTML = "";

  const days = daily.days || [];
  const items = updates.updates || [];
  const activeFeeds = new Set(items.map((x) => x.feed_title).filter(Boolean)).size;
  const today = days[0]?.total || 0;
  const top = items[0]?.feed_title || "-";

  [
    ["Tracked Updates", updates.count || 0],
    ["Today", today],
    ["Active Feeds", activeFeeds],
    ["Latest Source", top]
  ].forEach(([k, v]) => el.appendChild(statNode(k, v)));
}

function renderMarquee(updates) {
  const root = byId("feedMarquee");
  const feeds = Array.from(new Set((updates.updates || []).map((u) => u.feed_title).filter(Boolean))).slice(0, 24);
  if (feeds.length === 0) {
    root.innerHTML = '<div class="marquee-track"><span class="feed-chip">Waiting for fresh updates...</span></div>';
    return;
  }

  const row = [...feeds, ...feeds];
  const track = document.createElement("div");
  track.className = "marquee-track";
  row.forEach((feed) => {
    const chip = document.createElement("span");
    chip.className = "feed-chip";
    chip.textContent = feed;
    track.appendChild(chip);
  });

  root.innerHTML = "";
  root.appendChild(track);
}

function renderTimeline(updates) {
  const list = byId("timeline");
  const tmpl = byId("timelineItemTmpl");
  const items = updates.updates || [];
  list.innerHTML = "";

  if (items.length === 0) {
    const empty = document.createElement("p");
    empty.className = "empty";
    empty.textContent = "No fresh updates yet. The next half-hour run will append new entries here.";
    list.appendChild(empty);
  }

  items.forEach((item) => {
    const node = tmpl.content.firstElementChild.cloneNode(true);
    node.querySelector(".entry-feed").textContent = item.feed_title || "Unknown feed";
    node.querySelector(".entry-time").textContent = item.published ? shortFmt.format(new Date(item.published)) : "";

    const a = node.querySelector(".entry-title");
    a.textContent = item.title || "(untitled)";
    a.href = item.link || "#";

    node.querySelector(".entry-tldr").textContent = `TLDR: ${item.tldr || item.what || ""}`;
    node.querySelector(".entry-why").textContent = `Why it matters: ${item.why || ""}`;
    node.querySelector(".entry-action").textContent = `Action: ${item.action || ""}`;

    const tweetText = item.tweet || `${item.title || ""} ${item.link || ""}`.trim();
    node.querySelector(".tweet-preview").textContent = tweetText;
    const btn = node.querySelector(".tweet-btn");
    btn.addEventListener("click", async () => {
      try {
        await navigator.clipboard.writeText(tweetText);
        btn.textContent = "Copied";
      } catch {
        btn.textContent = "Copy failed";
      }
      setTimeout(() => (btn.textContent = "Copy Tweet Draft"), 1200);
    });

    list.appendChild(node);
  });

  byId("updateCount").textContent = `${items.length} updates`;
}

function renderDaily(daily) {
  const list = byId("dailyList");
  const tmpl = byId("dailyItemTmpl");
  const days = daily.days || [];
  list.innerHTML = "";

  if (days.length === 0) {
    const empty = document.createElement("p");
    empty.className = "empty";
    empty.textContent = "Daily summary will appear after enough updates are collected.";
    list.appendChild(empty);
  }

  days.forEach((d) => {
    const node = tmpl.content.firstElementChild.cloneNode(true);
    node.querySelector(".day-date").textContent = d.date;
    node.querySelector(".day-total").textContent = `${d.total} updates`;
    node.querySelector(".day-feeds").textContent = (d.feeds || []).slice(0, 4).join(" · ");

    const ul = node.querySelector(".day-points");
    (d.key_points || []).slice(0, 4).forEach((point) => {
      const li = document.createElement("li");
      li.textContent = point;
      ul.appendChild(li);
    });
    list.appendChild(node);
  });

  byId("dayCount").textContent = `${days.length} days`;
}

async function load() {
  try {
    const [updates, daily] = await Promise.all([
      fetchJSON("./data/updates.json"),
      fetchJSON("./data/daily.json")
    ]);

    renderStats(updates, daily);
    renderMarquee(updates);
    renderTimeline(updates);
    renderDaily(daily);

    byId("lastUpdated").textContent = updates.generated_at
      ? `Last sync: ${fmt.format(new Date(updates.generated_at))}`
      : "No sync yet";
  } catch (err) {
    byId("lastUpdated").textContent = `Load failed: ${err.message}`;
  }
}

byId("refreshBtn").addEventListener("click", load);

function validEmail(email) {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
}

function wireSubscribeForm() {
  const form = byId("subscribeForm");
  const emailInput = byId("subscribeEmail");
  const hint = byId("subscribeHint");
  if (!form || !emailInput || !hint) return;

  form.addEventListener("submit", (e) => {
    e.preventDefault();
    const email = (emailInput.value || "").trim().toLowerCase();
    if (!validEmail(email)) {
      hint.textContent = "Please enter a valid email address.";
      return;
    }

    const title = `[subscribe] ${email}`;
    const body = [
      `EMAIL: ${email}`,
      "",
      "SOURCE: github-pages",
      "",
      "Keep this issue OPEN to stay subscribed.",
      "Close this issue to unsubscribe."
    ].join("\n");

    const url = `${SUBSCRIBE_URL_BASE}?title=${encodeURIComponent(title)}&body=${encodeURIComponent(body)}`;
    hint.textContent = "Opening subscription confirmation page...";
    window.open(url, "_blank", "noopener");
  });
}

wireSubscribeForm();
setInterval(load, 60000);
load();
