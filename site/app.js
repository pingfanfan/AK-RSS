const byId = (id) => document.getElementById(id);
const SUBSCRIBE_URL_BASE = "https://github.com/pingfanfan/AK-RSS/issues/new";
const LANG_KEY = "akrss_lang";

const I18N = {
  en: {
    page_title: "AK Signal Deck",
    tag: "AK-RSS Signal Deck",
    hero_title: "Andrej K Recommended Intelligence Feed",
    subtitle: "Rolling updates from your OPML list, expanded with Andrej K recommendations, with AI briefs and platform-ready daily posts.",
    refresh: "Refresh Now",
    open_opml: "Open Source OPML",
    lang_zh: "中文",
    lang_en: "English",
    curation_title: "Curation Note",
    curation_text:
      "This watchlist uses your original feed universe with additional recommendations attributed to Andrej K, focused on high-signal engineering and AI reading.",
    subscribe_title: "2-Hour Email Subscription",
    subscribe_text: "Enter your email to receive new updates every 2 hours.",
    subscribe_placeholder: "you@example.com",
    subscribe_button: "Subscribe",
    subscribe_hint: "You will be redirected to a GitHub issue confirmation page.",
    subscribe_hint_opening: "Opening subscription confirmation page...",
    subscribe_hint_invalid: "Please enter a valid email address.",
    stats_updates: "Tracked Updates",
    stats_today: "Today",
    stats_feeds: "Active Feeds",
    stats_latest: "Latest Source",
    panel_daily: "Daily Capsule",
    panel_live: "Live Rolling Updates",
    empty_updates: "No fresh updates yet. The next run will append new entries here.",
    empty_daily: "Daily summary will appear after enough updates are collected.",
    last_sync: "Last sync",
    no_sync: "No sync yet",
    load_failed: "Load failed",
    updates_unit: "updates",
    days_unit: "days",
    label_tldr: "TLDR",
    label_why: "Why it matters",
    label_action: "Action",
    translation_prefix: "Translation:",
    source_label: "Original",
    copy_x: "Copy X Post",
    copy_linkedin: "Copy LinkedIn Post",
    copy_threads: "Copy Threads Post",
    copied: "Copied",
    copy_failed: "Copy failed",
    day_summary_label: "Daily Summary"
  },
  zh: {
    page_title: "AK 信号看板",
    tag: "AK-RSS 信号看板",
    hero_title: "Andrej K 推荐情报流",
    subtitle: "基于你的 OPML 信源滚动更新，并结合 Andrej K 的补全推荐，输出 AI 简报与可直接发布的多平台日报文案。",
    refresh: "立即刷新",
    open_opml: "查看开源 OPML",
    lang_zh: "中文",
    lang_en: "English",
    curation_title: "选源说明",
    curation_text: "这份 watchlist 以你的原始订阅为主，并加入 Andrej K 补全推荐，重点覆盖高信号的工程与 AI 内容。",
    subscribe_title: "每 2 小时邮件订阅",
    subscribe_text: "输入邮箱后，每 2 小时接收一次新增更新。",
    subscribe_placeholder: "you@example.com",
    subscribe_button: "订阅",
    subscribe_hint: "将跳转到 GitHub Issue 页面完成确认。",
    subscribe_hint_opening: "正在打开订阅确认页面...",
    subscribe_hint_invalid: "请输入有效邮箱地址。",
    stats_updates: "累计更新",
    stats_today: "今日条数",
    stats_feeds: "活跃信源",
    stats_latest: "最新来源",
    panel_daily: "每日摘要",
    panel_live: "滚动更新",
    empty_updates: "暂时没有新更新，下一轮任务会自动追加。",
    empty_daily: "积累足够更新后会自动生成每日摘要。",
    last_sync: "最近同步",
    no_sync: "暂无同步记录",
    load_failed: "加载失败",
    updates_unit: "条更新",
    days_unit: "天",
    label_tldr: "TLDR",
    label_why: "为什么重要",
    label_action: "建议动作",
    translation_prefix: "翻译：",
    source_label: "原文",
    copy_x: "复制 X 文案",
    copy_linkedin: "复制 LinkedIn 文案",
    copy_threads: "复制 Threads 文案",
    copied: "已复制",
    copy_failed: "复制失败",
    day_summary_label: "当日总结"
  }
};

function detectLang() {
  const saved = localStorage.getItem(LANG_KEY);
  if (saved === "zh" || saved === "en") return saved;
  return (navigator.language || "").toLowerCase().startsWith("zh") ? "zh" : "en";
}

let currentLang = detectLang();
let latestPayload = {
  updates: { updates: [], count: 0 },
  daily: { days: [] }
};

const formatters = {
  en: {
    full: new Intl.DateTimeFormat("en-GB", {
      year: "numeric",
      month: "short",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
      timeZone: "Europe/London"
    }),
    short: new Intl.DateTimeFormat("en-GB", {
      month: "short",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
      timeZone: "Europe/London"
    })
  },
  zh: {
    full: new Intl.DateTimeFormat("zh-CN", {
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
      timeZone: "Europe/London"
    }),
    short: new Intl.DateTimeFormat("zh-CN", {
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      hour12: false,
      timeZone: "Europe/London"
    })
  }
};

function t(key) {
  return I18N[currentLang][key] || I18N.en[key] || key;
}

function pickLocalized(item, base) {
  if (!item) return "";
  if (currentLang === "zh") {
    return item[`${base}_zh`] || item[base] || item[`${base}_en`] || "";
  }
  return item[`${base}_en`] || item[base] || item[`${base}_zh`] || "";
}

function normalizeText(s) {
  return (s || "").trim().replace(/\s+/g, " ").toLowerCase();
}

function getTitleTranslation(item) {
  const original = (item.title || "").trim();
  const translated = currentLang === "zh" ? (item.title_zh || "") : (item.title_en || "");
  if (!translated) return "";
  if (normalizeText(original) === normalizeText(translated)) return "";
  return translated.trim();
}

function setupStaticText() {
  document.documentElement.lang = currentLang === "zh" ? "zh-CN" : "en";
  document.title = t("page_title");

  byId("tagText").textContent = t("tag");
  byId("heroTitle").textContent = t("hero_title");
  byId("subtitleText").textContent = t("subtitle");
  byId("refreshBtn").textContent = t("refresh");
  byId("openOpmlLink").textContent = t("open_opml");

  byId("langBtnZh").textContent = t("lang_zh");
  byId("langBtnEn").textContent = t("lang_en");

  byId("curationTitle").textContent = t("curation_title");
  byId("curationText").textContent = t("curation_text");

  byId("subscribeTitle").textContent = t("subscribe_title");
  byId("subscribeText").textContent = t("subscribe_text");
  byId("subscribeEmail").placeholder = t("subscribe_placeholder");
  byId("subscribeSubmit").textContent = t("subscribe_button");

  const hint = byId("subscribeHint");
  if (!hint.dataset.userEdited) {
    hint.textContent = t("subscribe_hint");
  }

  byId("dailyHeading").textContent = t("panel_daily");
  byId("liveHeading").textContent = t("panel_live");

  byId("langBtnZh").classList.toggle("active", currentLang === "zh");
  byId("langBtnEn").classList.toggle("active", currentLang === "en");
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
    [t("stats_updates"), updates.count || 0],
    [t("stats_today"), today],
    [t("stats_feeds"), activeFeeds],
    [t("stats_latest"), top]
  ].forEach(([k, v]) => el.appendChild(statNode(k, v)));
}

function renderMarquee(updates) {
  const root = byId("feedMarquee");
  const feeds = Array.from(new Set((updates.updates || []).map((u) => u.feed_title).filter(Boolean))).slice(0, 24);
  if (feeds.length === 0) {
    root.innerHTML = `<div class="marquee-track"><span class="feed-chip">${t("empty_updates")}</span></div>`;
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

async function copyWithFeedback(btn, content, defaultLabel) {
  try {
    await navigator.clipboard.writeText(content || "");
    btn.textContent = t("copied");
  } catch {
    btn.textContent = t("copy_failed");
  }
  setTimeout(() => {
    btn.textContent = defaultLabel;
  }, 1200);
}

function renderTimeline(updates) {
  const list = byId("timeline");
  const tmpl = byId("timelineItemTmpl");
  const items = updates.updates || [];
  list.innerHTML = "";

  if (items.length === 0) {
    const empty = document.createElement("p");
    empty.className = "empty";
    empty.textContent = t("empty_updates");
    list.appendChild(empty);
  }

  const shortFmt = formatters[currentLang].short;

  items.forEach((item) => {
    const node = tmpl.content.firstElementChild.cloneNode(true);
    node.querySelector(".entry-feed").textContent = item.feed_title || "Unknown feed";
    node.querySelector(".entry-time").textContent = item.published ? shortFmt.format(new Date(item.published)) : "";

    const a = node.querySelector(".entry-title");
    a.textContent = item.title || "(untitled)";
    a.href = item.link || "#";

    const translated = getTitleTranslation(item);
    const tNode = node.querySelector(".entry-title-translation");
    if (translated) {
      tNode.textContent = `${t("translation_prefix")} ${translated}`;
    } else {
      tNode.remove();
    }

    const tldr = pickLocalized(item, "tldr") || pickLocalized(item, "what");
    const why = pickLocalized(item, "why");
    const action = pickLocalized(item, "action");

    node.querySelector(".entry-tldr").textContent = `${t("label_tldr")}: ${tldr || ""}`;
    node.querySelector(".entry-why").textContent = `${t("label_why")}: ${why || ""}`;
    node.querySelector(".entry-action").textContent = `${t("label_action")}: ${action || ""}`;

    const xText = pickLocalized(item, "x") || pickLocalized(item, "tweet") || `${item.title || ""} ${item.link || ""}`.trim();
    node.querySelector(".tweet-preview").textContent = xText;

    const btn = node.querySelector(".tweet-btn");
    const label = t("copy_x");
    btn.textContent = label;
    btn.addEventListener("click", () => copyWithFeedback(btn, xText, label));

    list.appendChild(node);
  });

  byId("updateCount").textContent = `${items.length} ${t("updates_unit")}`;
}

function dailyPlatformText(day, platform) {
  const postObj = day.social_posts?.[platform] || {};
  const localized = currentLang === "zh" ? (postObj.zh || postObj.en) : (postObj.en || postObj.zh);
  if (localized) return localized;

  const fallback = day.tweet_drafts || [];
  return fallback.slice(0, 2).join("\n") || "";
}

function buildBullets(points) {
  return (points || [])
    .slice(0, 4)
    .filter(Boolean)
    .map((point) => {
      const li = document.createElement("li");
      li.textContent = point;
      return li;
    });
}

function renderDaily(daily) {
  const list = byId("dailyList");
  const tmpl = byId("dailyItemTmpl");
  const days = daily.days || [];
  list.innerHTML = "";

  if (days.length === 0) {
    const empty = document.createElement("p");
    empty.className = "empty";
    empty.textContent = t("empty_daily");
    list.appendChild(empty);
  }

  days.forEach((d) => {
    const node = tmpl.content.firstElementChild.cloneNode(true);
    node.querySelector(".day-date").textContent = d.date;
    node.querySelector(".day-total").textContent = `${d.total} ${t("updates_unit")}`;
    node.querySelector(".day-feeds").textContent = (d.feeds || []).slice(0, 4).join(" · ");

    const summary = currentLang === "zh" ? (d.summary_zh || "") : (d.summary_en || "");
    const summaryFallback = (d.key_points || []).slice(0, 3).join(currentLang === "zh" ? "；" : "; ");
    node.querySelector(".day-summary-label").textContent = t("day_summary_label");
    node.querySelector(".day-summary").textContent = summary || summaryFallback;

    const ul = node.querySelector(".day-points");
    buildBullets(d.key_points).forEach((li) => ul.appendChild(li));

    const platformMap = [
      { key: "x", title: "X", copyKey: "copy_x" },
      { key: "linkedin", title: "LinkedIn", copyKey: "copy_linkedin" },
      { key: "threads", title: "Threads", copyKey: "copy_threads" }
    ];

    platformMap.forEach((p) => {
      const block = node.querySelector(`.social-${p.key}`);
      const textNode = block.querySelector(".social-post");
      const btn = block.querySelector(".social-copy");

      block.querySelector(".social-platform").textContent = p.title;
      const content = dailyPlatformText(d, p.key);
      textNode.textContent = content;

      const label = t(p.copyKey);
      btn.textContent = label;
      btn.addEventListener("click", () => copyWithFeedback(btn, content, label));
    });

    list.appendChild(node);
  });

  byId("dayCount").textContent = `${days.length} ${t("days_unit")}`;
}

function renderLastUpdated(updates) {
  const fullFmt = formatters[currentLang].full;
  const el = byId("lastUpdated");
  if (updates.generated_at) {
    el.textContent = `${t("last_sync")}: ${fullFmt.format(new Date(updates.generated_at))}`;
    return;
  }
  el.textContent = t("no_sync");
}

function renderAll() {
  setupStaticText();
  const updates = latestPayload.updates || { updates: [], count: 0 };
  const daily = latestPayload.daily || { days: [] };

  renderStats(updates, daily);
  renderMarquee(updates);
  renderTimeline(updates);
  renderDaily(daily);
  renderLastUpdated(updates);
}

async function fetchJSON(path) {
  const res = await fetch(`${path}?t=${Date.now()}`, { cache: "no-store" });
  if (!res.ok) throw new Error(`${path} ${res.status}`);
  return res.json();
}

async function load() {
  try {
    const [updates, daily] = await Promise.all([fetchJSON("./data/updates.json"), fetchJSON("./data/daily.json")]);
    latestPayload = { updates, daily };
    renderAll();
  } catch (err) {
    byId("lastUpdated").textContent = `${t("load_failed")}: ${err.message}`;
  }
}

function setLang(next) {
  if (next !== "zh" && next !== "en") return;
  currentLang = next;
  localStorage.setItem(LANG_KEY, next);
  renderAll();
}

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
      hint.textContent = t("subscribe_hint_invalid");
      hint.dataset.userEdited = "1";
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
    hint.textContent = t("subscribe_hint_opening");
    hint.dataset.userEdited = "1";
    window.open(url, "_blank", "noopener");
  });
}

function wireEvents() {
  byId("refreshBtn").addEventListener("click", load);
  byId("langBtnZh").addEventListener("click", () => setLang("zh"));
  byId("langBtnEn").addEventListener("click", () => setLang("en"));
  wireSubscribeForm();
}

wireEvents();
renderAll();
setInterval(load, 60000);
load();
