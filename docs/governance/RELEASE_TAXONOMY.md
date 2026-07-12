# CoreOps – Release and Version Taxonomy

> Status: **Proposed for acceptance**
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-003` (docs-only)

Bis zum Nova Review und Human-Maintainer-Commit ist diese Taxonomie ein Vorschlag. Nach dem Human-Maintainer-Commit gilt sie als bindende Projektgovernance. **Dieses Dokument erzeugt keinen Tag und keinen Release.**

## Purpose

Auflösung der Mehrdeutigkeit zwischen der Foundation-Phase und dem ersten funktionalen Produktmeilenstein (Auflösung von CCR-02) sowie Festlegung einer eindeutigen Versions- und Tag-Taxonomie.

## Milestone Names

Phasen- und Produktmeilensteine (Namen; Nummern/Versionen siehe unten):

- **Foundation** – dokumentations- und governanceorientierte Phase
- **Observe** – erster funktionaler Produktmeilenstein
- **Map**
- **Plan**
- **Deploy**
- **Automate**
- **Extend**
- **Scale**

Trennung der zuvor kollidierenden Bezeichnungen:

- `Foundation 0.1` = **interner Name** der Foundation-Phase (Dokumentation/Governance), **kein** funktionaler Produktrelease.
- `Observe` = **Name** des ersten funktionalen Produktmeilensteins (nicht mehr als „0.1" nummeriert, um die Kollision zu vermeiden).

## Foundation Versioning

```text
Foundation 0.1
Tag candidate: v0.0.1-foundation
```

- `v0.0.1-foundation` enthält ausschließlich Foundation-Dokumentation und Governance.
- Es ist **kein** funktionaler CoreOps-Produktrelease und behauptet **keine** funktionale Produktreife.

## Product Versioning

```text
Observe first prerelease candidate:
v0.1.0-alpha.1
```

- `v0.1.0-alpha.1` darf erst **nach** funktionaler Implementierung und einer **separaten Readiness-Prüfung** entstehen.
- Spätere Observe-Prereleases folgen SemVer-Prerelease-Konventionen (`v0.1.0-alpha.2`, `v0.1.0-beta.1`, …).

## SemVer Rules

- Tags sind **unveränderlich**.
- **Keine Wiederverwendung** eines Tags.
- Ein Foundation-Tag behauptet **keine** funktionale Produktreife.
- Instabile funktionale Versionen tragen eine **Prerelease-Kennzeichnung** (`-alpha.N`, `-beta.N`).
- Finale/stabile Releases entstehen **nur** nach einem eigenen Readiness- und Release-Prep-Work-Package.

## Human-Maintainer Gates

Erstellung, Veröffentlichung und Signierung von Tags und Releases sind ausschließlich dem Human Maintainer vorbehalten. Kein Implementation Agent und kein Skill erzeugt Tags oder Releases.

## No Automatic Release

Es erfolgt keine automatische Tag- oder Release-Erzeugung. Dieses Work Package legt die Taxonomie fest, führt aber **keine** Release-Aktion aus.

## Relationship to Roadmap

Die [ROADMAP.md](../../ROADMAP.md) nennt die Meilensteinnamen und die vorläufigen Versionskandidaten (`v0.0.1-foundation`, `v0.1.0-alpha.1`). Die Roadmap bleibt ein Planungsartefakt ohne Termine und ohne Fertigstellungsversprechen; diese Taxonomie ist die maßgebliche Versions-Referenz.

## Open Questions

- Endgültige Versionsnummerierung der Meilensteine nach `Observe` (Map/Plan/Deploy/…) — offen, ohne SemVer-Zusage.
- Zeitpunkt des Übergangs von `-alpha` zu `-beta` und zu einem stabilen Release — an spätere Readiness-/Release-Prep-Work-Packages gebunden.
