

# 🏗️ Trump Tariff Tracker - Developer Overview

## 🚀 Project Summary
A responsive web app that visualizes tariffs imposed during the Trump era, including retaliation tariffs from other countries. The frontend is a bold, MAGA-themed leaderboard interface that lets users explore tariff data in real time.

---

## 👥 Target Users
- Journalists, researchers, and educators
- Politically-engaged citizens
- Meme/satire communities interested in trade

---

## ✅ Core Features
- Leaderboard of countries based on tariff data (US ↔ Others)
- Detailed breakdown per country (products, rates, dates)
- Mobile-friendly layout
- Satirical MAGA-style theme

---

## 🔧 Architecture
- **Monorepo** with:
  - `apps/frontend`: Vue 3 + Vite + TypeScript
  - `apps/backend`: Go + Gorilla Mux
  - `libs/types`: Shared TS/Go types
  - `infra/`: Docker, DB migrations, CI (GitHub Actions)

- **DB**: PostgreSQL with `countries` and `tariffs` tables
- **Infra**: Docker Compose, Fly.io or Render

---

## 🧭 Routes and APIs

### 🔹 Frontend Routes
| Path               | Description                          |
|--------------------|--------------------------------------|
| `/`                | Leaderboard of tariffs               |
| `/country/:code`   | Detailed breakdown (modal or page)   |

### 🔹 API Routes

#### `GET /api/leaderboard`
Returns list of countries with max tariff %.
```json
[
  {
    "country": "India",
    "country_code": "IN",
    "max_tariff": 72.0,
    "direction": "retaliation" // or "us_to_them"
  }
]
```

#### `GET /api/country/:code`
Returns tariff breakdown for one country.
```json
{
  "country": "India",
  "tariffs_from_country": [
    { "product": "Steel", "tariff": 40.0, "type": "standard", "target_country": "USA" }
  ],
  "tariffs_from_us": [
    { "product": "Pharmaceuticals", "tariff": 30.0, "type": "embargo", "target_country": "India" }
  ]
}
```

#### `POST /api/refresh` *(optional, token-protected)*
Manually refreshes cached data from external sources.

---

## 🧱 Database Schema

### `countries`
- `id` (PK)
- `name`
- `code` (ISO Alpha-2)
- `flag_url`

### `tariffs`
- `id` (PK)
- `country_id` (FK → countries.id)
- `target_country` (FK → countries.id)
- `product`
- `type`: enum (`standard`, `embargo`, `legislation`)
- `tariff` (percentage)
- `last_updated`

Indexes:
- `country_id`, `target_country`, `product`

---

## 📦 Data Flow

1. **Frontend** fetches `/api/leaderboard` on load.
2. Clicking a country triggers `/api/country/:code`.
3. Backend queries DB for product-level tariffs in both directions.
4. Data is rendered in modals and tables.

---

## 🎨 UI Summary
- **Leaderboard**: Sortable table showing countries and tariff percentages.
- **Modal view**: Per-country tariff details with MAGA-style elements.
- **Extras**: Optional cursor effects, tooltips, sound effects.

---

## 🧪 Stack
- Frontend: Vue 3, Pinia, Axios
- Backend: Go (Gorilla Mux, GORM/sqlc)
- DB: PostgreSQL
- Infra: Docker, GitHub Actions, Fly.io or Render

---

## 🔒 Auth
- No login required.
- Optional token-based `/refresh` endpoint.

---

## 🌐 Free Public Data APIs (For scraping or sourcing)
- **[Federal Register API](https://www.federalregister.gov/developers/v1)** – US trade policy notices
- **[GDELT Project](https://blog.gdeltproject.org/gdelt-2-0-our-global-world-in-realtime/)** – Real-time events and export actions
- **[EU TARIC API](https://ec.europa.eu/taxation_customs/dds2/taric/taric_consultation.jsp?Lang=en)** – EU tariff data
- **[WTO API](https://data.wto.org/developers)** – Tariff and trade data (registration required)

Use scrapers or watchers for:
- **USTR**
- **USITC**
- **UN Comtrade**

---

## ✅ Developer Checklist

### 🔌 Backend/API Tasks
- [x] Set up Go API project with Gorilla Mux
- [x] Connect to PostgreSQL via GORM
- [x] Implement `/api/leaderboard` route
- [x] Implement `/api/country/:code` route
- [x] Implement `/api/refresh` endpoint 
- [ ] Run Federal Register API check on refresh call and save changes to the DB
- [ ] Run GDELT Project check every on refresh call and save changes to the DB
- [ ] Run EU TARIC API check every on refresh call and save changes to the DB
- [ ] Run WTO API check every on refresh call and save changes to the DB
- [ ] Seed initial data (sample countries + tariffs)

### 🗄️ Database Tasks
- [ ] Create `countries` and `tariffs` tables (see migration SQL)
- [ ] Create indexes for `product`, `country_id`, and `target_country`
- [ ] Add enum `tariff_type`

### 🧪 Frontend Tasks
- [ ] Vue 3 app scaffold with Pinia
- [ ] Fetch and display leaderboard
- [ ] Modal or page view for country breakdown
- [ ] Sorting and responsiveness
- [ ] Tooltip and meme UX elements

### ⚙️ Infra & Deployment
- [ ] Write Dockerfiles for backend and frontend
- [x] Set up Docker Compose
- [ ] GitHub Actions pipeline
- [ ] Deploy to Fly.io / Render

---






```sql
-- 001_init.sql

CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code CHAR(2) NOT NULL UNIQUE,
    flag_url TEXT
);

CREATE TYPE tariff_type AS ENUM ('standard', 'embargo', 'legislation');

CREATE TABLE tariffs (
    id SERIAL PRIMARY KEY,
    country_id INTEGER NOT NULL REFERENCES countries(id) ON DELETE CASCADE,
    target_country INTEGER NOT NULL REFERENCES countries(id) ON DELETE CASCADE,
    product TEXT NOT NULL,
    type tariff_type NOT NULL,
    tariff NUMERIC(5,2) NOT NULL,
    last_updated TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

CREATE INDEX idx_tariffs_country_id ON tariffs(country_id);
CREATE INDEX idx_tariffs_target_country ON tariffs(target_country);
CREATE INDEX idx_tariffs_product ON tariffs(product);
```