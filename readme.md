# SyncMind – Plan your energy, not just your time

Smart assistant that syncs your calendar and wearable data to give personalized suggestions for sleep, meals, supplements, and training.

SyncMind helps you stay in your optimal flow state by analyzing your calendar, sleep, and activity data.  
It suggests the best timing for meals, caffeine, supplements, and recovery — so you can plan your energy, not just your time.

➡️ [Watch demo video](https://youtu.be/U4Xq8WQw0Gs)

## ✨ Features
- Integration with Google Calendar (via ical file)
- Daily summary based on sleep quality and planned events
- Personalized caffeine timing
- Meal timing & protein intake suggestions
- Supplement stack suggestions (magnesium, L-theanine, Rhodiola, etc.)
- Context-based day analysis (morning / noon / evening insights)


## ⚙️ Tech Stack
- **Frontend:** Svelte (Vite)
- **Backend:** Go
- **Infra:** Docker + Makefile (one-command build)
- **APIs:** runalyze.com API (for smartwatch integration)


## 📁 Monorepo structure
```
.
├── apps/         
│   ├── core/                 # Go API server
│   │   └── go.mod
│   └── webapp/               # Svelte web app
│       ├── src/
│       └── package.json
│
├── Makefile                  # Unified build commands
│
├── .docker/    
│   └── core/ 
│       └── Dockerfile
│
└── README.md
```

## 👤 Author
Created by **Piotr Sarna** – solo project for HackYeah2025.  
