const byId = (id) => document.getElementById(id);
const fmt = new Intl.DateTimeFormat("en-GB", {
  year: "numeric",
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

function renderStats(updates, daily) {
  const days = daily.days || [];
  const total = updates.count || 0;
  const today = days[0]?.total || 0;
  const feeds = new Set((updates.updates || []).map((x) => x.feed_title).filter(Boolean)).size;

  byId("stats").innerHTML = [
    ["tracked updates", total],
    ["today", today],
    ["active feeds", feeds]
  ].map(([k, v]) => `<div class="card"><div class="k">${k}</div><div class="v">${v}</div></div>`).join("");
}

function renderTimeline(updates) {
  const list = byId("timeline");
  const tmpl = byId("timelineItemTmpl");
  list.innerHTML = "";

  for (const item of (updates.updates || [])) {
    const node = tmpl.content.firstElementChild.cloneNode(true);
    const a = node.querySelector(".item-title");
    a.textContent = item.title || "(untitled)";
    a.href = item.link || "#";
    node.querySelector(".time").textContent = item.published ? fmt.format(new Date(item.published)) : "";
    node.querySelector(".feed").textContent = item.feed_title || "";
    node.querySelector(".tldr").textContent = `TLDR: ${item.tldr || item.what || ""}`;
    node.querySelector(".action").textContent = `Action: ${item.action || ""}`;

    const btn = node.querySelector(".tweet-btn");
    const tweet = item.tweet || `${item.title || ""} ${item.link || ""}`;
    btn.addEventListener("click", async () => {
      await navigator.clipboard.writeText(tweet);
      btn.textContent = "Copied";
      setTimeout(() => (btn.textContent = "Copy Tweet Draft"), 1200);
    });

    list.appendChild(node);
  }

  byId("updateCount").textContent = `${updates.count || 0} updates`;
}

function renderDaily(daily) {
  const list = byId("dailyList");
  const tmpl = byId("dailyItemTmpl");
  list.innerHTML = "";

  for (const d of (daily.days || [])) {
    const node = tmpl.content.firstElementChild.cloneNode(true);
    node.querySelector(".item-title").textContent = `${d.date} · ${d.total} updates`;
    node.querySelector(".time").textContent = "Daily";
    node.querySelector(".feed").textContent = (d.feeds || []).slice(0, 3).join(" · ");

    const ul = node.querySelector(".mini-list");
    for (const p of (d.key_points || []).slice(0, 3)) {
      const li = document.createElement("li");
      li.textContent = p;
      ul.appendChild(li);
    }
    list.appendChild(node);
  }

  byId("dayCount").textContent = `${(daily.days || []).length} days`;
}

async function load() {
  try {
    const [updates, daily] = await Promise.all([
      fetchJSON("./data/updates.json"),
      fetchJSON("./data/daily.json")
    ]);
    renderStats(updates, daily);
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
setInterval(load, 60000);
load();
