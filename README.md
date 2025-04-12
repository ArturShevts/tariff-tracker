# Trump Tariff Tracker - Product Requirements Document

## 1. Elevator Pitch
The Trump Tariff Tracker is a responsive web app that visualizes the current state of tariffs imposed by the United States under Donald Trump's trade policies, along with retaliation tariffs imposed by other countries. It provides a fun, informative, and politically-themed interface that allows users to explore the global impact of Trump's trade wars in real-time, with a bold red-white-blue Trump-flare.

## 2. Who is this app for
- Journalists and researchers covering international trade and Trump-era policies
- Politically-engaged citizens and policy enthusiasts
- Students and educators seeking up-to-date trade data
- Satirical and meme communities with an interest in global economics

## 3. Functional Requirements
- Display a real-time, sortable leaderboard of countries the US has tariffs against
- Show corresponding retaliation tariffs from those countries
- Automatically surface the country with the highest tariff at the top
- Clicking on a country opens a detailed breakdown (by goods, dates, percentages)
- Mobile-responsive layout for phones and tablets
- Tariff data is fetched from an internal API powered by AI aggregation from public sources (e.g. USTR, WTO, UN Comtrade)

## 4. User Stories
- As a user, I want to see which country has the highest tariffs against the US so I can understand the retaliation landscape
- As a user, I want to click on a country to view a detailed list of tariffs imposed by and against the US
- As a user, I want to quickly scan a leaderboard of tariff levels for all countries affected
- As a user, I want the interface to be fun, bold, and politically themed to reflect the spirit of the Trump era

## 5. User Interface
- Homepage features a bold red-white-blue theme with MAGA-style icons and Trump caricatures (e.g. leaderboard called "The Trade War Winners")
- Leaderboard format similar to [covid19tracker.ca](https://covid19tracker.ca) with sortable columns (Country, US Tariffs, Retaliation Tariffs)
- Each country name is clickable, expanding into a breakdown section showing product categories, percentage rates, and dates
- Header banner with “TRUMP TARIFF TRACKER” in bold fonts with flags, eagles, and playful design elements
- Responsive design ensures clean display on desktops, tablets, and mobile devices

# Trump Tariff Tracker - User Interface Design Document

## Layout Structure
- **Header**: Full-width banner featuring “Trump Tariff Tracker” in bold MAGA-style font with animated flags, bald eagles, and a waving Trump.
- **Main View**:
    - **Leaderboard Table** (sticky, scrollable): Centerpiece of the layout with sortable columns for Country, US Tariffs, Retaliation Tariffs.
    - **Modal Expansion Panel**: Clicking a country opens a modal card with that country's flag, breakdown of key tariffs, and trade summary.
- **Sidebar (optional)**: Floating Trump quotes or MAGA tweets for added humor.
- **Footer**: Disclaimer about satire, sources of data, and credits.

## Core Components
- **Leaderboard Table**:
    - Country Name (with flag icon)
    - US Tariffs %
    - Retaliation Tariffs %
    - Trump badge overlays for fun: e.g. “SAD!” for weak tariffs, “TREMENDOUS!” for high ones
- **Country Modal**:
    - Title bar with country name and trade balance emoji
    - Tariff breakdown list (categories, percentages, date of last update)
    - “Trump Rating” for each country (meme-style ranking)
- **Tooltip Microcopy**:
    - Hovering over badges shows Trump-style comments (e.g. “The best tariffs, believe me.”)

## Interaction Patterns
- **Sortable Table Columns**: Clicking headers sorts by tariff % values
- **Row Click**: Opens modal card with more data on the selected country
- **Tooltips on Hover**: Tooltip bubbles with Trump-style quips
- **Sticky Headers**: Header and table headers stick while scrolling
- **Optional Fun**:
    - Trump airplane cursor
    - Sound effect: “You’re fired!” on modal close (optional toggle)

## Visual Design Elements & Color Scheme
- **Color Scheme**:
    - Background: White or light parchment-style
    - Primary: Bold red (#FF0000)
    - Secondary: Deep blue (#002868)
    - Accent: Gold highlights (for maximum Trump energy)
- **Graphics**:
    - Trump caricatures (GIFs/static)
    - Country flags as icons
    - “MAGA-styled” badges, borders, buttons

## Mobile, Web App, Desktop Considerations
- **Responsive Web App**:
    - Collapsible leaderboard into cards for mobile view
    - Sticky headers become floating toolbars on scroll
    - Modals slide from bottom on mobile instead of pop-up
- **No dedicated desktop app**; optimized for browser use

## Typography
- **Header Font**: Chunky, campaign-style serif (e.g. "Bebas Neue", "Anton")
- **Body Font**: Clean sans-serif like “Open Sans” or “Roboto”
- **Font Styling**:
    - Bold and uppercase for headers
    - Occasional exaggerated styling (caps, color) for comedic emphasis

## Accessibility
- Sufficient contrast between text and background
- Alt text for all icons, images, and buttons
- Keyboard-navigable modal windows and tables
- Sound effects are optional and toggleable
- All tooltip data also accessible as in-page text for screen readers


# Trump Tariff Tracker - Software Requirements Specification (SRS)

## System Design
- **Monorepo**: Shared codebase for frontend (Vue + TS), backend (Go), shared types, and infrastructure configs (e.g. Docker, DB migrations).
- Central database stores structured tariff data.
- Frontend queries an internal API and renders sorted and categorized tariff views.
- Tariff details exist at the product/industry level, and each country view is a dynamic breakdown of these.

## Architecture Pattern
- **Monorepo with Workspaces**:
  - `apps/frontend`: Vue 3 SPA
  - `apps/backend`: Go API service
  - `libs/types`: Shared TS/Go type definitions
  - `infra/`: Docker, migrations, CI config

## State Management
- Pinia (Vue) for reactive global state:
  - Leaderboard data (highest tariffs)
  - Selected country view (modal/card)
  - Sound/UX preferences

## Data Flow
1. Frontend loads list of countries with max tariff against or from the US.
2. Clicking a country triggers a fetch:
   - Tariffs from that country → US
   - Tariffs from US → that country
   - Grouped by `product`
3. Server queries DB and returns structured response.
4. Frontend renders it in infographic/modal view.

## Technical Stack
- **Frontend**: Vue 3, TypeScript, Vite, Pinia, Axios
- **Backend**: Go, Gorilla Mux, GORM or sqlc
- **Database**: PostgreSQL
- **Infra**: Docker, Docker Compose, GitHub Actions (CI), possibly Fly.io or Render for deployment

## Authentication Process
- No user authentication required for read-only access.
- Optional environment-token protected `/refresh` endpoint.

## Route Design
### Frontend
- `/` — Tariff leaderboard (countries with highest tariffs)
- `/country/:code` — Expanded view of a country (modal or separate page on mobile)

### Backend API
- `GET /api/leaderboard` — Returns all countries with max tariff %
- `GET /api/country/:code` — Tariffs to/from this country grouped by product
- `POST /api/refresh` — (Optional) Admin-triggered scrape and cache update

## API Design

### `GET /api/leaderboard`
Returns:
```json
[
  {
    "country": "India",
    "country_code": "IN",
    "max_tariff": 72.0,
    "direction": "retaliation" // or "us_to_them"
  },
  ...
]
```

### `GET /api/country/:code`
Returns:
```json
{
  "country": "India",
  "tariffs_from_country": [
    {
      "product": "Steel",
      "tariff": 40.0,
      "type": "standard",
      "target_country": "USA"
    },
    ...
  ],
  "tariffs_from_us": [
    {
      "product": "Pharmaceuticals",
      "tariff": 30.0,
      "type": "embargo",
      "target_country": "India"
    },
    ...
  ]
}
```

## Database Design ERD

### `countries`
- `id` (PK)
- `name`
- `code` (ISO Alpha-2)
- `flag_url`

### `tariffs`
- `id` (PK)
- `country_id` (FK → countries.id)
- `target_country` (FK → countries.id)
- `product` (string)
- `type` (enum: `'standard' | 'embargo' | 'legislation'`)
- `tariff` (float)
- `last_updated` (timestamp)

**Relationships**:
- One `country` can have many outbound tariffs
- Each tariff has both a `source country` and a `target country`

---

## Initial SQL Migration (PostgreSQL)

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
---

Let me know if you want:
- A Makefile or Docker setup for running these migrations
- A Go DB access layer with sqlc
- Seed data to test the frontend quickly

Or ready to generate the initial backend + API scaffolding?