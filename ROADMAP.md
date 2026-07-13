# CoreOps – Roadmap (Gerüst)

**NDF-Basis:** `v1.0.0` (Tag `v1.0.0`, Commit `9dcadc1`) — `main` ist informativ, **nicht** normativ.

Dieses Dokument ist ein **kompaktes, vorläufiges Gerüst**. Es enthält bewusst **keine** finale SemVer-Zuordnung, keine Release-Termine, keine Fertigstellungsversprechen und keine detaillierten Technologieentscheidungen.

## Produktmeilensteine (vorläufige Namen und Nummern)

> Die Nummerierung ist **vorläufig** und dient nur der Ordnung. Sie ist **keine** verbindliche Versions- oder Release-Nummer.

1. Foundation – Platform Foundation *(vorläufig)* — nur Dokumentation, Architektur, Governance, Security.
2. Observe *(vorläufig)* — **vollständig read-only**; Inventory, generische Protokoll- und Read-only-Domain-Beobachtung (Details: [FOUNDATION_CAPABILITY_MATRIX.md](docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md), [INITIAL_SUPPORT_BOUNDARY.md](docs/integrations/INITIAL_SUPPORT_BOUNDARY.md)).
3. Map *(vorläufig)* — Topologie, LLDP/CDP, MAC/ARP, VLAN, Path Explorer und Basis-IPAM.
4. Plan *(vorläufig)* — Desired State, Drift, Deployment-Blueprints, Policy Evaluation und Preview.
5. Deploy *(vorläufig)* — erste kontrollierte Write-Aktionen (z. B. Docker Compose), Health Checks, Rollback, Evidence.
6. Automate *(vorläufig)* — Workflows, Runbooks, Skriptbibliothek, Event-Korrelation.
7. Extend *(vorläufig)* — Windows, Hyper-V, Kubernetes, Netzwerk-/Druckerbereitstellung, Firmware.
8. Scale *(vorläufig)* — Multi-Site, Site Relay, HA, Multi-Tenant.

## Release-Taxonomie (Proposed for acceptance)

Die frühere Kollision zwischen `Foundation 0.1` und `Release 0.1 – Observe` ist in CO-WP-003 vorgeschlagen aufgelöst (siehe [RELEASE_TAXONOMY.md](docs/governance/RELEASE_TAXONOMY.md)):

```text
Foundation documentation release:
v0.0.1-foundation candidate

First functional Observe prerelease:
v0.1.0-alpha.1 candidate
```

- `Foundation 0.1` ist die dokumentations-/governanceorientierte Phase (kein funktionaler Produktrelease).
- `Observe` ist der erste funktionale Produktmeilenstein; sein erster Prerelease-Kandidat ist `v0.1.0-alpha.1` (erst nach funktionaler Implementierung und separater Readiness-Prüfung).
- Status `Proposed for acceptance` bis zum Human-Maintainer-Commit. Es wird **kein** Tag und **kein** Release erzeugt.

## Strategische Richtungen (Accepted Product Direction, CO-WP-004A)

- CoreOps Core bleibt **unabhängig** von externen Managementprodukten (siehe [SOVEREIGNTY_AND_DEPENDENCY_POLICY.md](docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md)).
- **Hardened Profile** ist ein späterer Foundation-/Release-Kandidat.
- **Government Profile** ist ein späterer Readiness-Kandidat (keine Zertifizierung/Zulassung).
- **BSI-Nachweisartefakte** werden in späteren Work Packages definiert (BSI-orientiert; siehe [BSI_ALIGNMENT_POSITIONING.md](docs/security/BSI_ALIGNMENT_POSITIONING.md)) — keine Zertifizierung versprochen.
- **ITIL-/PRINCE2-Tailoring** bleibt offen (Kandidaten; Entscheidung in `CO-WP-004D`).
- **Lessons Learned** werden projektübergreifend kontrolliert ausgewertet; NDF-Rückfluss nur via eigenes NDF-WP (`CO-WP-004B`).

## Hinweise

- Keine detaillierten Technologieentscheidungen in diesem Dokument.
- Keine Release-Termine.
- Keine unrealistischen Fertigstellungsversprechen.
