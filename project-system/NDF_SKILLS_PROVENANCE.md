# NDF v1.0.0 Skills – Provenance

Menschlich lesbare Herkunfts- und Verifikationsdokumentation für den lokal bereitgestellten NDF-Skills-Bestand. Maschinenlesbares Gegenstück: [ndf-skills-lock.json](ndf-skills-lock.json).

Erzeugt durch Work Package `CO-WP-001A – Complete NDF v1.0.0 Skills Bootstrap` (docs-only).

## Zweck

Der vollständige Claude-Skills-Bestand des normativen NDF-Tags `v1.0.0` wurde lokal im CoreOps-Projekt verfügbar gemacht, damit spätere Work Packages die passenden NDF-Skills **lokal, versionsgebunden und offline** auswählen können. Der Bestand wird **vorsorglich vollständig** bereitgestellt; er wird **nicht** automatisch aktiviert oder vollständig in jeden Promptkontext geladen.

## Normative Quelle

- **Repository:** https://github.com/KayKaspers/Nova-Development-Framework
- **Tag:** `v1.0.0`
- **Commit:** `9dcadc1` (verifiziert: `HEAD` = `9dcadc12fb960914b9a5baeff2ab1aee75912b57`)
- **Quellpfad:** `.claude/skills/`

Der Branch `main` ist **nicht** normativ und wurde **nicht** verwendet. Es wurden keine Inhalte aus `main`, späteren Releases, anderen Branches, Forks, Mirrors, Skill-Marktplätzen, Package Managern, der GitHub API oder Raw-Dateien anderer Refs übernommen.

## Importmodell

- Vollständiger Skills-Bestand des normativen Tags (keine Auswahl einzelner Skills).
- **Byte-identische** Übernahme (keine Änderung von Inhalt, Line-Endings oder Frontmatter).
- Keine zusätzlichen lokalen Skills ergänzt.
- Keine Übernahme aus `main`.
- Kein Marketplace, kein Package Manager.
- Bezug über einen temporären Clone außerhalb des CoreOps-Repositorys (`git clone --branch v1.0.0 --depth 1`), der nach der Verifikation vollständig entfernt wurde.

## Bestand

- **Skill-Verzeichnisse:** 38
- **Reguläre Dateien:** 39 (38 × `SKILL.md` + 1 × `README.md` im Skills-Wurzelverzeichnis)
- **Gesamtgröße:** 102.761 Bytes

Enthaltene Skills:

- ndf-accessibility-reviewer
- ndf-adr-governance-review
- ndf-architecture-blueprint-runner
- ndf-behavioral-adoption-reviewer
- ndf-branding-kit-runner
- ndf-changelog-writer
- ndf-compact-context-summary-runner
- ndf-content-tone-reviewer
- ndf-context-pack-maintainer
- ndf-creative-direction-runner
- ndf-debugging-root-cause-reviewer
- ndf-docs-polish-runner
- ndf-ethical-growth-reviewer
- ndf-existing-project-analysis-runner
- ndf-feature-scope-runner
- ndf-feedback-triage-runner
- ndf-implementation-review-runner
- ndf-landing-page-concept-runner
- ndf-naming-runner
- ndf-onboarding-friction-reviewer
- ndf-privacy-data-minimization-reviewer
- ndf-product-discovery-runner
- ndf-project-adapter-quality-reviewer
- ndf-project-brief-runner
- ndf-public-neutrality-guard
- ndf-public-release-body-reviewer
- ndf-readme-quality-reviewer
- ndf-release-notes-runner
- ndf-release-safety
- ndf-skill-quality-reviewer
- ndf-skill-supply-chain-risk-reviewer
- ndf-skill-trigger-quality-reviewer
- ndf-test-strategy-runner
- ndf-ui-style-system-runner
- ndf-ux-flow-reviewer
- ndf-v1-readiness-review
- ndf-validation-evidence-reviewer
- ndf-work-package-runner

Die vollständige Datei-zu-Hash-Liste ist in [ndf-skills-lock.json](ndf-skills-lock.json) hinterlegt.

## Verifikation

- **Commit-Prüfung:** `git rev-parse HEAD` = `9dcadc1…` ✅
- **Tag-Prüfung:** `git tag --points-at HEAD` enthält `v1.0.0` ✅
- **Anzahl Skills:** 38 ✅
- **Anzahl Dateien:** 39 ✅
- **SHA-256 je Datei:** in [ndf-skills-lock.json](ndf-skills-lock.json) (lowercase, alphabetisch sortiert) ✅
- **Pfadgleichheit** Quelle ↔ Ziel ✅
- **Hashgleichheit** Quelle ↔ Ziel ✅
- **Bytegleichheit** Quelle ↔ Ziel (102.761 Bytes) ✅
- **Source Safety:** keine symbolischen Links, keine Special Files, kein verschachteltes Repository innerhalb `.claude/skills/` ✅

## Sicherheitsgrenzen

- Keine Skill-Datei wurde ausgeführt, interpretiert oder als Script geladen.
- Keine Skill-Datei wurde verändert.
- Keine automatische Aktivierung aller Skills.
- Skills ersetzen **keine** Human-Maintainer-Freigabe.
- Skills dürfen **keine** autonomen Commits, Pushes, Releases oder produktiven Änderungen durchführen.
- Skills sind **keine** technische Sandbox; sie ersetzen nicht die Scope- und Sicherheitsgrenzen eines Work-Package-Prompts.

## Nutzungsmodell

Alle Skills sind lokal verfügbar. Pro Work Package werden nur passende Skills ausgewählt. Die Auswahl muss berücksichtigen:

- Work-Package-Typ
- Prompt Mode
- Scope
- Allowed Files
- Sicherheitsgrenzen
- Context Economy

Vollständige lokale Verfügbarkeit bedeutet **nicht** vollständiges Laden in jeden Kontext.

## Update-Regel

Jede spätere Aktualisierung, Ergänzung oder Entfernung des NDF-Skills-Bestands benötigt:

- ein eigenes Work Package,
- eine neue Quellenprüfung,
- eine neue Hashprüfung,
- einen aktualisierten Lock,
- Nova Review,
- Human-Maintainer-Freigabe.
