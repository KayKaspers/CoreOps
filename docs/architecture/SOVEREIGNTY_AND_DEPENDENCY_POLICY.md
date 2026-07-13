# CoreOps – Sovereignty and Dependency Policy

> Status: **Accepted Product Direction** (technische Umsetzung offen)
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004A` (docs-only / governance-baseline)

Dieses Dokument regelt die Produktsouveränität von CoreOps und die kontrollierte Behandlung von Abhängigkeiten. Es akzeptiert **keine** konkrete Technologie und ersetzt keine Human-Maintainer-Entscheidung.

## 1. Purpose

Sicherstellen, dass CoreOps ein eigenständiges, souveränes, self-hosted und offline-/air-gap-fähiges Produkt bleibt, und Abhängigkeiten kontrolliert, nachvollziehbar und exit-fähig zu halten.

## 2. Core Product Independence

```text
CoreOps ist ein eigenständiges Produkt.

Die Kernplattform muss ohne verpflichtende externe Monitoring-,
CMDB-, ITSM-, Automation-, Container-Management- oder
GitOps-Plattform betrieben werden können.
```

Unabhängigkeit bedeutet **nicht** vollständige Eigenentwicklung jeder technischen Basisfunktion. Etablierte, kontrolliert eingebundene Basiskomponenten bleiben grundsätzlich zulässig (Abschnitt 5–6).

## 3. Mandatory External Product Prohibition

Folgende externen Produktklassen dürfen **keine verpflichtende Laufzeitabhängigkeit** des CoreOps-Kerns sein:

- externe Monitoring-Suiten
- externe CMDBs
- externe ITSM-Systeme
- externe Automation Control Planes
- externe Container-Management-Plattformen
- externe GitOps-Plattformen
- verpflichtende SaaS- oder Cloudservices

Externe Produkte dürfen als Markt-, Architektur-, Bedienungs- und Sicherheitsreferenz **analysiert** werden. Ableitungen erfolgen als allgemeine Muster/Ideen, ohne Code, geschützte Inhalte oder Produktabhängigkeiten ungeprüft zu übernehmen.

## 4. Optional Integrations

Die genannten Produktklassen dürfen als **optionale** Integrationen unterstützt werden, sofern der Kern ohne sie voll funktionsfähig bleibt. Eine optionale Integration darf nie zur stillen Kernvoraussetzung werden.

## 5. Technical Foundation Dependencies

Nach späterer, dokumentierter Bewertung grundsätzlich zulässig:

- Datenbanken
- etablierte Frameworks
- Kryptobibliotheken
- Protokollbibliotheken
- Telemetriekomponenten
- Build- und Packaging-Komponenten

> In diesem Work Package wird **keine** konkrete Technologie ausgewählt oder akzeptiert.

## 6. Dependency Admission Criteria

Jede spätere technische Basisabhängigkeit benötigt mindestens:

1. dokumentierten Zweck
2. Herkunft
3. Versionskontrolle
4. Lizenzprüfung
5. SBOM-Erfassung
6. Schwachstellenprozess
7. Offline-Verfügbarkeit
8. Updatepfad
9. Migrations- oder Exit-Strategie
10. Human-Maintainer-Freigabe

## 7. Offline Availability

Jede zugelassene Basisabhängigkeit muss für den self-hosted, offline-/air-gap-fähigen Betrieb verfügbar sein (z. B. über kontrollierte lokale Repositories / CorePack). Eine Abhängigkeit, die zwingend Online-Konnektivität voraussetzt, ist für Kernfunktionen unzulässig.

## 8. Supply-Chain Expectations

Zugelassene Abhängigkeiten unterliegen den Erwartungen des CoreOps-Trust-Modells (Concept v3.0 §23): Digest, Signatur/Provenance sofern verfügbar, SBOM, Schwachstellen- und Widerrufsstatus, kontrollierter Import. Details werden in den späteren Supply-Chain-/Security-Baseline-Work-Packages ausgearbeitet.

## 9. Exit and Replacement Strategy

Für jede kritische Basisabhängigkeit ist eine dokumentierte Exit-/Ersatzstrategie vorzusehen, damit CoreOps bei Lizenz-, Sicherheits- oder Verfügbarkeitsproblemen souverän bleibt. Diese Strategie wird pro Abhängigkeit im jeweiligen Zulassungs-Work-Package konkretisiert.

## 10. Exceptions and Human Gates

Ausnahmen von dieser Policy benötigen ein eigenes Work Package, Nova Review und Human-Maintainer-Freigabe sowie ein Project-Brain-/Context-Pack-Update. Zu Beginn gibt es **keine** stillen Ausnahmen.

## 11. Current Status

- Produktsouveränität: **akzeptierte Produktrichtung**.
- Verpflichtende externe Managementprodukte: **verboten** als Kernabhängigkeit.
- Technische Basisabhängigkeiten: **offen**, spätere kontrollierte Bewertung.
- Es ist **keine** konkrete Dependency akzeptiert.
