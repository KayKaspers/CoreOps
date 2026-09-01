# CoreOps – Foundation Capability Matrix

> Document Status: Accepted
> Technical Architecture Status: Unconfirmed
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004 – Foundation Capability Matrix and Initial Support Boundary` (docs-only)

Diese Matrix ist eine **Planungs- und Governance-Landkarte**, keine Implementierung und kein Supportnachweis. Begleitdokument: [INITIAL_SUPPORT_BOUNDARY.md](../integrations/INITIAL_SUPPORT_BOUNDARY.md).

## Zentrale Invarianten

```text
Capability target ≠ implemented capability
Implemented capability ≠ verified capability
Verified capability ≠ supported integration
Vendor priority ≠ vendor support
Detected device ≠ managed device
Read capability ≠ write capability
Generic protocol support ≠ support for every device using that protocol
```

## Aktueller Projektstatus

```text
CoreOps besitzt noch keine implementierte Runtime-Capability.
Keine Integration besitzt aktuell den Supportstatus Supported.
```

Entsprechend gilt in dieser Matrix für **jede** Zeile: `Implementation Status = not-implemented` und `Support Status = not-supported`. Keine Zeile behauptet etwas anderes.

## Drei getrennte Statusdimensionen

Die drei Dimensionen sind voneinander unabhängig.

**Roadmap Status** (nur Planung): `foundation-defined` · `target-observe` · `target-map` · `target-plan` · `target-deploy` · `target-automate` · `target-extend` · `target-scale` · `deferred` · `non-goal`

**Implementation Status**: `not-implemented` · `prototype` · `implemented-unverified` · `verified` — aktuell für alle produktiven Capabilities: `not-implemented`.

**Support Status**: `not-applicable` · `not-supported` · `experimental` · `preview` · `supported` · `deprecated` · `blocked` · `revoked`
- `supported` nur mit dokumentierter Evidenz.
- `experimental`/`preview` setzen mindestens eine Implementierung voraus.
- Eine nur geplante Capability ist `not-supported`.
- `blocked`/`revoked` benötigen eine dokumentierte Begründung.

## Erweitertes Status- und Governance-Modell (CO-WP-004E)

`CO-WP-004E` erweitert dieses Statusmodell additiv um zwei weitere Dimensionen (**Evidence Status**, **Security/Governance Status**) sowie **Profile Relevance**, **PSR-Domänen-Zuordnung** und **Responsibility-Codes**. Die bestehenden drei Dimensionen (Roadmap/Implementation/Support) bleiben unverändert autoritativ. Die vollständige Definition steht in [CAPABILITY_MATRIX_SPEC.md](../project-system/CAPABILITY_MATRIX_SPEC.md) und [CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md](../security/CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md). Die per-Capability-Zuordnung steht im Abschnitt „Security and Governance Alignment" unten. **Keine** Capability wird durch CO-WP-004E hochgestuft, gelöscht oder umbenannt; PSR-Zuordnung bedeutet Readiness-Relevanz, **nicht** BSI-Compliance.

## Funktionale Integrationslevel

```text
Level 0 – Detected
Level 1 – Inventoried
Level 2 – Observed
Level 3 – Protected
Level 4 – Managed
Level 5 – Automated
```

- Level gelten **pro Integration und Capability-Gruppe**, nicht pauschal pro Hersteller.
- Ein Gerät kann bei Inventory Level 1, bei Monitoring Level 2 und bei Configuration weiterhin Level 0/unsupported sein.
- Das Vorhandensein eines Adapters allein verleiht **kein** Level.
- Der erste funktionale **Observe**-Meilenstein reicht höchstens bis **Level 2**.
- Level 3–5 sind **nicht** Teil der initialen Observe-Supportgrenze.

## Spalten und Legende

Jede Matrixzeile enthält: `Capability ID`, `Domain`, `Capability`, `Description`, `Roadmap Status`, `Earliest Milestone`, `Read or Write`, `Required Trust Boundary`, `Offline Expectation`, `Dependency or Prerequisite`, `Implementation Status`, `Support Status`, `Evidence Required`, `Notes`.

Zur Kompaktheit werden folgende Kürzel verwendet:
- **R/W:** `read` (Beobachtung eines Zielsystems) · `write-target` (schreibende Aktion auf ein Zielsystem) · `platform` (CoreOps-interne Funktion/Selbstkonfiguration, keine Zielsystem-Schreibaktion).
- **Trust Boundary (Plane):** `Exp` Experience · `Ctrl` Control · `Data` Data · `Exec` Execution · `Trust` Trust.
- **Offline:** `off-cap` offline-fähig · `local` rein lokal · `degraded` mit letztem bekannten Stand · `online-pref` bevorzugt online.
- **Impl:** `not-impl` = not-implemented (gilt für alle Zeilen).
- **Support:** `not-supp` = not-supported (gilt für alle Zeilen).
- **Evidence Required:** `SES` = vollständiger Support-Evidence-Satz (siehe Support Boundary) · `impl+test` = Implementierung + dokumentierte Tests · `impl+test+sec` = zusätzlich Security-Review.

> **Read/Write-Hinweis:** Die Observe-Read-only-Grenze bezieht sich auf **Zielsysteme**. `platform`-Capabilities (Setup, Owner, Dashboards) konfigurieren CoreOps selbst und sind kein Zielsystem-Write; sie sind mit der Observe-Read-only-Grenze vereinbar. `write-target` ist im Observe-Meilenstein ausgeschlossen.

---

## Platform and Experience

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-PLATFORM-001 | Platform | Setup Wizard | Geführte lokale Ersteinrichtung | target-observe | Observe | platform | Exp | local | — | not-impl | not-supp | impl+test | Selbstkonfiguration, kein Zielsystem-Write |
| CAP-PLATFORM-002 | Platform | DE/EN Bilingual UI | Zweisprachige Oberfläche, DE als möglicher Default | target-observe | Observe | platform | off-cap | — | i18n-Modell | not-impl | not-supp | impl+test | APC (Produktanforderung) |
| CAP-PLATFORM-003 | Platform | Native Dashboards | Rollen-/standortabhängige Übersichtsansichten | target-observe | Observe | read | Exp | degraded | Inventory, Monitoring | not-impl | not-supp | impl+test | — |
| CAP-PLATFORM-004 | Platform | Integrated Documentation | In-App-Dokumentation | target-observe | Observe | platform | off-cap | — | — | not-impl | not-supp | impl+test | — |
| CAP-PLATFORM-005 | Platform | Global Search | Übergreifende Suche über Assets/Ansichten | target-observe | Observe | read | Exp | degraded | Inventory | not-impl | not-supp | impl+test | — |
| CAP-PLATFORM-006 | Platform | Audit Explorer | Read-only Ansicht der Auditdaten | target-observe | Observe | read | Data | off-cap | Audit-Modell | not-impl | not-supp | impl+test | — |
| CAP-PLATFORM-007 | Platform | Mobile Core Views (PWA) | Mobile Kernansichten | target-observe | Observe | read | Exp | degraded | Native Dashboards | not-impl | not-supp | impl+test | — |

## Identity and Governance

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-IDENTITY-001 | Identity | Local Owner & Basic Roles | Lokaler Owner und grundlegende Rollen | target-observe | Observe | platform | Ctrl | local | — | not-impl | not-supp | impl+test+sec | Break-Glass separat |
| CAP-IDENTITY-002 | Identity | RBAC | Rollenbasierte Rechte, mehrere Begrenzungsebenen | target-plan | Plan | platform | Ctrl | local | Local Owner | not-impl | not-supp | impl+test+sec | Basisrollen früh, Feinrechte später |
| CAP-IDENTITY-003 | Identity | Policy Evaluation | Kontextabhängige fail-closed Policy-Bewertung | target-plan | Plan | platform | Ctrl | local | Policy-Modell | not-impl | not-supp | impl+test+sec | ADR-pflichtig |
| CAP-IDENTITY-004 | Identity | Approval Engine | Freigaben/Approval Inbox | target-plan | Plan | platform | Ctrl | online-pref | Policy | not-impl | not-supp | impl+test+sec | — |
| CAP-IDENTITY-005 | Identity | Break Glass | Auditierter lokaler Notzugang | target-plan | Plan | platform | Trust | local | Local Owner | not-impl | not-supp | impl+test+sec | Nicht für Normalbetrieb |
| CAP-IDENTITY-006 | Identity | Machine Identity | Interne PKI/mTLS/Enrollment kurzlebiger Identitäten | target-extend | Extend | platform | Trust | degraded | interne PKI | not-impl | not-supp | impl+test+sec | Air-Gap-Spannung (CCR-06) |

## Inventory and Source of Truth

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-INVENTORY-001 | Inventory | Asset Inventory (read-only) | Read-only Bestandsführung | target-observe | Observe | read | Data | degraded | Discovery/Import | not-impl | not-supp | impl+test | — |
| CAP-INVENTORY-002 | Inventory | Sites | Standortmodell | target-observe | Observe | platform | Data | off-cap | — | not-impl | not-supp | impl+test | — |
| CAP-INVENTORY-003 | Inventory | Groups | Gruppierung von Assets | target-observe | Observe | platform | Data | off-cap | — | not-impl | not-supp | impl+test | — |
| CAP-INVENTORY-004 | Inventory | Field Provenance | Feldbasierte Herkunft/Vertrauen/Status | target-observe | Observe | read | Data | off-cap | SoT-Modell | not-impl | not-supp | impl+test | — |
| CAP-INVENTORY-005 | Inventory | Conflict Detection | Konflikt-/Prioritätsregeln je Feldtyp | target-map | Map | read | Data | off-cap | Field Provenance | not-impl | not-supp | impl+test | ADR-0009-Kandidat |
| CAP-INVENTORY-006 | Inventory | Historical States | Historisierung von Zuständen | target-map | Map | read | Data | off-cap | Inventory | not-impl | not-supp | impl+test | — |
| CAP-INVENTORY-007 | Inventory | Basic IPAM | Minimales IP-Adressmanagement | target-map | Map | read | Data | off-cap | Inventory | not-impl | not-supp | impl+test | Explizit Map |

## Discovery

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-DISCOVERY-001 | Discovery | ICMP Reachability | Erreichbarkeit in freigegebenen Bereichen | target-observe | Observe | read | Exec | online-pref | Range-Governance | not-impl | not-supp | SES | Rate Limits |
| CAP-DISCOVERY-002 | Discovery | SNMPv3 Discovery/Inventory | Sichere SNMPv3-Ermittlung | target-observe | Observe | read | Exec | online-pref | Credential-Handling | not-impl | not-supp | SES | SNMPv1/v2c kein sicherer Default |
| CAP-DISCOVERY-003 | Discovery | Generic HTTP/HTTPS Check | Statuscode/Antwortzeit/TLS-Metadaten | target-observe | Observe | read | Exec | online-pref | — | not-impl | not-supp | SES | Keine sensiblen Bodies speichern |
| CAP-DISCOVERY-004 | Discovery | Linux Discovery (read-only) | Basisinventar/OS/Version | target-observe | Observe | read | Exec | online-pref | Transport offen | not-impl | not-supp | SES | Agent/Agentless offen |
| CAP-DISCOVERY-005 | Discovery | Proxmox Discovery (read-only) | Cluster/Node/VM/LXC-Inventar | target-observe | Observe | read | Exec | online-pref | Proxmox-API | not-impl | not-supp | SES | Keine VM-Aktion |
| CAP-DISCOVERY-006 | Discovery | Docker Discovery (read-only) | Hosts/Container/Images/Compose | target-observe | Observe | read | Exec | online-pref | sicherer Zugriffspfad offen | not-impl | not-supp | SES | Kein Exec/Shell |
| CAP-DISCOVERY-007 | Discovery | Printer MIB/IPP Discovery | Druckerermittlung | target-observe | Observe | read | Exec | online-pref | — | not-impl | not-supp | SES | — |
| CAP-DISCOVERY-008 | Discovery | Controlled Network Ranges | Scan-Governance (Ranges/Ausschlüsse/Limits) | target-observe | Observe | platform | Ctrl | local | — | not-impl | not-supp | impl+test+sec | Kein ungefragter Vollscan |

## Monitoring and Observability

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-MONITORING-001 | Monitoring | Reachability | Erreichbarkeitsüberwachung | target-observe | Observe | read | Data | degraded | Discovery | not-impl | not-supp | SES | — |
| CAP-MONITORING-002 | Monitoring | Health | Zustandsüberwachung | target-observe | Observe | read | Data | degraded | Discovery | not-impl | not-supp | SES | — |
| CAP-MONITORING-003 | Monitoring | Metrics | Standardisierte Metriken | target-observe | Observe | read | Data | degraded | Telemetry Norm. | not-impl | not-supp | SES | — |
| CAP-MONITORING-004 | Monitoring | Alerts | Regelbasierte Warnungen | target-observe | Observe | read | Data | degraded | Metrics | not-impl | not-supp | impl+test | — |
| CAP-MONITORING-005 | Monitoring | Grafana Workspace | Analyse-Workspace | target-observe | Observe | read | Exp | degraded | Metrics/Logs | not-impl | not-supp | impl+test | Analyse, nicht einzige UI |
| CAP-MONITORING-006 | Monitoring | Logs | Logaggregation | target-observe | Observe | read | Data | degraded | — | not-impl | not-supp | impl+test | — |
| CAP-MONITORING-007 | Monitoring | Telemetry Normalization | Semantisches CoreOps-Schema | target-observe | Observe | read | Data | off-cap | — | not-impl | not-supp | impl+test | Rohdaten bleiben verfügbar |
| CAP-MONITORING-008 | Monitoring | CoreScore | Zusammengesetzter Gesundheitswert | target-map | Map | read | Data | degraded | Monitoring | not-impl | not-supp | impl+test | Fehlende Daten ≠ gesund |

## Network

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-NETWORK-001 | Network | Generic Device Status | Erreichbarkeit/Firmware/Uptime, sofern standardisiert | target-observe | Observe | read | Data | online-pref | SNMPv3 | not-impl | not-supp | SES | Keine Config-Änderung |
| CAP-NETWORK-002 | Network | Interface Status | Interface-Basisdaten | target-observe | Observe | read | Data | online-pref | SNMPv3 | not-impl | not-supp | SES | — |
| CAP-NETWORK-003 | Network | LLDP/CDP Neighbors | Nachbarschaftsermittlung | target-map | Map | read | Data | online-pref | SNMPv3 | not-impl | not-supp | SES | Map, nicht Observe |
| CAP-NETWORK-004 | Network | MAC/ARP Tables | Adress-/Nachbarschaftstabellen | target-map | Map | read | Data | online-pref | SNMPv3 | not-impl | not-supp | SES | Map |
| CAP-NETWORK-005 | Network | VLAN Data | VLAN-Informationen | target-map | Map | read | Data | online-pref | SNMPv3 | not-impl | not-supp | SES | Map |
| CAP-NETWORK-006 | Network | Path Explorer | Pfaddarstellung | target-map | Map | read | Data | degraded | Topology | not-impl | not-supp | impl+test | Map |
| CAP-NETWORK-007 | Network | Impact Analysis | Auswirkungsanalyse | target-map | Map | read | Data | degraded | Topology | not-impl | not-supp | impl+test | Map |
| CAP-NETWORK-008 | Network | Network Write Actions | Port/PoE/VLAN/Config/Firmware | target-extend | Extend | write-target | Exec | online-pref | Policy/Approval/Backup | not-impl | not-supp | impl+test+sec | Nicht Observe |

## Topology

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-TOPOLOGY-001 | Topology | Topology Graph & Evidence | Evidenzbasierter Topologiegraph mit Confidence | target-map | Map | read | Data | degraded | Network reads | not-impl | not-supp | impl+test | Keine Verbindung als Fakt |
| CAP-TOPOLOGY-002 | Topology | Manual Topology Authority | Manuelle Overrides/gesperrte Verbindungen | target-map | Map | platform | Data | off-cap | Topology Graph | not-impl | not-supp | impl+test | Invariante: keine Auto-Änderung |

## Print

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-PRINT-001 | Print | Printer Status | Gesamtstatus/Warnungen | target-observe | Observe | read | Data | online-pref | SNMP/IPP | not-impl | not-supp | SES | — |
| CAP-PRINT-002 | Print | Supplies | Toner/Tinte/Trommel/Kits | target-observe | Observe | read | Data | online-pref | SNMP/IPP | not-impl | not-supp | SES | — |
| CAP-PRINT-003 | Print | Counters | Seitenzähler | target-observe | Observe | read | Data | online-pref | SNMP/IPP | not-impl | not-supp | SES | — |
| CAP-PRINT-004 | Print | Alerts | Druckerwarnungen | target-observe | Observe | read | Data | online-pref | SNMP/IPP | not-impl | not-supp | SES | — |
| CAP-PRINT-005 | Print | Queue Read | Warteschlangenanzeige | target-map | Map | read | Data | online-pref | Printserver | not-impl | not-supp | impl+test+sec | Keine Auftragsnamen als Default-Telemetrie |
| CAP-PRINT-006 | Print | Queue Management | Pause/Resume/kontrolliertes Abbrechen | target-extend | Extend | write-target | Exec | online-pref | Policy/Approval | not-impl | not-supp | impl+test+sec | Nicht Observe |
| CAP-PRINT-007 | Print | Driver & Port Documentation | Treiber-/Portdokumentation | target-extend | Extend | read | Data | online-pref | Printserver | not-impl | not-supp | impl+test | — |
| CAP-PRINT-008 | Print | Printer Deployment | Druckerbereitstellung | target-extend | Extend | write-target | Exec | online-pref | Deploy-Modell | not-impl | not-supp | impl+test+sec | Nicht Observe |

## Virtualization and Containers

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-VIRT-001 | Virtualization | Proxmox Inventory (read) | Node/VM/LXC-Inventar & Health | target-observe | Observe | read | Data | online-pref | Proxmox-API | not-impl | not-supp | SES | Keine VM-Aktion |
| CAP-VIRT-002 | Virtualization | Docker Inventory (read) | Container/Images/Compose (read) | target-observe | Observe | read | Data | online-pref | sicherer Zugriff offen | not-impl | not-supp | SES | Kein Exec/Restart |
| CAP-VIRT-003 | Virtualization | Podman Inventory (read) | Podman-Bestand | target-extend | Extend | read | Data | online-pref | — | not-impl | not-supp | SES | — |
| CAP-VIRT-004 | Virtualization | Hyper-V Management | Hyper-V-Inventar/Verwaltung | target-extend | Extend | write-target | Exec | online-pref | Policy/Approval | not-impl | not-supp | impl+test+sec | Nicht Observe |
| CAP-VIRT-005 | Virtualization | VMware Management | VMware-Inventar/Verwaltung | target-extend | Extend | read | Data | online-pref | — | not-impl | not-supp | SES | Management nicht Observe |
| CAP-VIRT-006 | Virtualization | Kubernetes Management | K8s-Inventar/Verwaltung/Helm | target-extend | Extend | write-target | Exec | online-pref | Policy/Approval | not-impl | not-supp | impl+test+sec | Nicht Observe |

## Deployments and Change

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-DEPLOY-001 | Deployments | Artifact Registry Link | Verweis auf verifizierte Artefakte | target-deploy | Deploy | read | Trust | degraded | Artifact Trust | not-impl | not-supp | impl+test+sec | — |
| CAP-DEPLOY-002 | Deployments | Deployment Blueprint | Deklaratives Blueprint-Modell | target-plan | Plan | platform | Ctrl | off-cap | — | not-impl | not-supp | impl+test | Modell in Plan |
| CAP-DEPLOY-003 | Deployments | Preview | Vorschau vor Ausführung | target-plan | Plan | read | Ctrl | degraded | Blueprint | not-impl | not-supp | impl+test | — |
| CAP-DEPLOY-004 | Deployments | Approval (deploy) | Freigabe vor Deployment | target-plan | Plan | platform | Ctrl | online-pref | Policy | not-impl | not-supp | impl+test+sec | — |
| CAP-DEPLOY-005 | Deployments | Health Check | Post-Change-Verifikation | target-deploy | Deploy | read | Data | online-pref | Monitoring | not-impl | not-supp | impl+test | — |
| CAP-DEPLOY-006 | Deployments | Rollback | Rücknahme auf Last-Known-Good | target-deploy | Deploy | write-target | Exec | online-pref | Backup/LKG | not-impl | not-supp | impl+test+sec | — |
| CAP-DEPLOY-007 | Deployments | Docker Compose Deployment | Kontrollierte Compose-Deployments | target-deploy | Deploy | write-target | Exec | online-pref | Policy/Approval/Health | not-impl | not-supp | impl+test+sec | — |
| CAP-DEPLOY-008 | Deployments | Core Products Deployment | Deployment der Core-Produkte | target-deploy | Deploy | write-target | Exec | online-pref | Policy/Approval | not-impl | not-supp | impl+test+sec | — |
| CAP-DEPLOY-009 | Deployments | Linux Services Deployment | Ausgewählte Linux-Dienste | target-deploy | Deploy | write-target | Exec | online-pref | Policy/Approval | not-impl | not-supp | impl+test+sec | — |
| CAP-DEPLOY-010 | Deployments | Network Config Deployment | Netzwerkkonfiguration ausrollen | target-extend | Extend | write-target | Exec | online-pref | Backup/Diff/Preview | not-impl | not-supp | impl+test+sec | — |
| CAP-DEPLOY-011 | Deployments | Firmware Deployment | Firmware-Verteilung | target-extend | Extend | write-target | Exec | online-pref | Policy/Approval | not-impl | not-supp | impl+test+sec | — |

## Automation and Response

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-AUTOMATION-001 | Automation | Scripts | Signierter Skript-Lebenszyklus | target-automate | Automate | write-target | Exec | degraded | Trust/Policy | not-impl | not-supp | impl+test+sec | Kein Ausführen ungeprüfter Skripte |
| CAP-AUTOMATION-002 | Automation | Workflows | Persistente, wiederaufnehmbare Workflows | target-automate | Automate | write-target | Exec | degraded | Policy/Approval | not-impl | not-supp | impl+test+sec | Idempotenz erforderlich |
| CAP-AUTOMATION-003 | Automation | Runbooks | Doku/Checkliste/kontrollierter Workflow | target-automate | Automate | read | Ctrl | off-cap | — | not-impl | not-supp | impl+test | Doku nicht auto-ausführbar |
| CAP-AUTOMATION-004 | Automation | Maintenance Campaigns | Wartungskampagnen/Wellen | target-automate | Automate | write-target | Exec | online-pref | Policy/Approval | not-impl | not-supp | impl+test+sec | — |
| CAP-AUTOMATION-005 | Automation | Incidents | Incident-Koordination | target-automate | Automate | read | Ctrl | degraded | Events | not-impl | not-supp | impl+test | — |
| CAP-AUTOMATION-006 | Automation | Event Correlation | Korrelation zusammenhängender Symptome | target-automate | Automate | read | Data | degraded | Eventmodell | not-impl | not-supp | impl+test | — |

## Trust and Offline

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-TRUST-001 | Trust | Secrets Management | Verschlüsselter Secret Store, Rotation | target-plan | Plan | platform | Trust | local | — | not-impl | not-supp | impl+test+sec | Offene Frage: minimales Credential-Handling für Read-Integrationen |
| CAP-TRUST-002 | Trust | Internal PKI | Interne Zertifizierungsstelle | target-extend | Extend | platform | Trust | local | — | not-impl | not-supp | impl+test+sec | — |
| CAP-TRUST-003 | Trust | mTLS | Gegenseitige TLS-Authentifizierung | target-extend | Extend | platform | Trust | degraded | interne PKI | not-impl | not-supp | impl+test+sec | — |
| CAP-TRUST-004 | Trust | Artifact Verification | Signatur-/Hash-/Provenance-Prüfung | target-deploy | Deploy | read | Trust | off-cap | — | not-impl | not-supp | impl+test+sec | — |
| CAP-TRUST-005 | Trust | SBOM | Stücklistenerfassung | target-deploy | Deploy | read | Trust | off-cap | — | not-impl | not-supp | impl+test | — |
| CAP-TRUST-006 | Trust | Provenance | Build-Provenance | target-deploy | Deploy | read | Trust | off-cap | — | not-impl | not-supp | impl+test+sec | — |
| CAP-TRUST-007 | Trust | CorePack | Signiertes Offline-Transferpaket | target-deploy | Deploy | platform | Trust | off-cap | Verification | not-impl | not-supp | impl+test+sec | Import mit Human-Freigabe |
| CAP-TRUST-008 | Trust | Offline Queue | Verschlüsselte lokale Queue | target-automate | Automate | platform | Exec | off-cap | — | not-impl | not-supp | impl+test+sec | — |
| CAP-TRUST-009 | Trust | Store-and-Forward | Gepufferte spätere Übertragung | target-extend | Extend | platform | Exec | off-cap | Site Relay | not-impl | not-supp | impl+test+sec | — |

## Protection and Recovery

| ID | Domain | Capability | Description | Roadmap | Earliest MS | R/W | Trust | Offline | Dependency | Impl | Support | Evidence | Notes |
|----|--------|-----------|-------------|---------|-------------|-----|-------|---------|-----------|------|---------|----------|-------|
| CAP-PROTECT-001 | Protection | Configuration Vault | Versionierte Konfigurationen | target-deploy | Deploy | read | Data | off-cap | — | not-impl | not-supp | impl+test+sec | Redaction |
| CAP-PROTECT-002 | Protection | Backup | Konfigurations-/Zustandssicherung | target-deploy | Deploy | read | Data | off-cap | Config Vault | not-impl | not-supp | impl+test | — |
| CAP-PROTECT-003 | Protection | Restore Readiness | Wiederherstellbarkeitsprüfung | target-deploy | Deploy | read | Data | off-cap | Backup | not-impl | not-supp | impl+test+sec | Backup ≠ wiederherstellbar |
| CAP-PROTECT-004 | Protection | Last-Known-Good | Bekannter guter Zustand | target-deploy | Deploy | read | Data | off-cap | Config Vault | not-impl | not-supp | impl+test | — |
| CAP-PROTECT-005 | Protection | Self-Dependency Protection | Schutz der eigenen Betriebsgrundlage | target-deploy | Deploy | platform | Ctrl | local | — | not-impl | not-supp | impl+test+sec | Quorum/Exclusion |
| CAP-PROTECT-006 | Protection | Degraded Modes | Kontrollierte Teilstörungsmodi | target-observe | Observe | platform | Ctrl | degraded | — | not-impl | not-supp | impl+test | Fail-closed bei Policy-Ausfall |
| CAP-PROTECT-007 | Protection | Support Bundle | Redigierter Diagnoseexport | target-observe | Observe | read | Data | off-cap | Redaction | not-impl | not-supp | impl+test+sec | Keine Secrets/PII |
| CAP-PROTECT-008 | Protection | Recovery Mode | Reduzierte Notfalloberfläche | target-observe | Observe | platform | Ctrl | local | — | not-impl | not-supp | impl+test | Keine normalen Deployments |

---

## Security and Governance Alignment (CO-WP-004E)

Additive per-Capability-Zuordnung (gleiche IDs, keine parallele Liste). Spalten: **Evidence** (Evidence Status) · **Sec/Gov** (Security/Governance Status) · **Profiles** (Readiness: S=Standard/H=Hardened/G=Government; Governance: SO=Service Operations/MDP=Major Deployment Project/PSD=Public-Sector Delivery) · **PSR** (Readiness-Domänen PSR-01…18, Relevanz — **keine** BSI-Compliance) · **Resp** (P=Product/O=Operator/S=Shared).

> Evidence Status ist projektweit `not-assessed` (keine per-Capability-Evidenz in diesem WP bewertet). Kein `implemented`/`supported`/`validated`/`compliant`. `controls-required` markiert sicherheitskritische oder schreibende Capabilities, die noch Controls-Design benötigen.

### Platform and Experience
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-PLATFORM-001 | not-assessed | baseline-mapped | S/H/G | PSR-01,05 | P/O |
| CAP-PLATFORM-002 | not-assessed | not-assessed | S/H/G | mapping-review-required | P |
| CAP-PLATFORM-003 | not-assessed | baseline-mapped | S/H/G; SO | PSR-16 | P/O |
| CAP-PLATFORM-004 | not-assessed | baseline-mapped | S/H/G | PSR-16 | P/O |
| CAP-PLATFORM-005 | not-assessed | baseline-mapped | S/H/G | PSR-16 | P/O |
| CAP-PLATFORM-006 | not-assessed | baseline-mapped | S/H/G; PSD | PSR-07,16 | P/O |
| CAP-PLATFORM-007 | not-assessed | baseline-mapped | S/H/G | PSR-16 | P/O |

### Identity and Governance
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-IDENTITY-001 | not-assessed | controls-required | S/H/G; PSD | PSR-03,01 | P/O/S |
| CAP-IDENTITY-002 | not-assessed | controls-required | S/H/G; PSD | PSR-03,04 | P/O/S |
| CAP-IDENTITY-003 | not-assessed | controls-required | H/G; PSD | PSR-04,13 | P/O/S |
| CAP-IDENTITY-004 | not-assessed | controls-required | H/G; MDP,PSD | PSR-04,13 | P/O/S |
| CAP-IDENTITY-005 | not-assessed | controls-required | H/G; PSD | PSR-04 | P/O/S |
| CAP-IDENTITY-006 | not-assessed | controls-required | H/G | PSR-03,12 | P/O/S |

### Inventory and Source of Truth
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-INVENTORY-001 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O |
| CAP-INVENTORY-002 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O |
| CAP-INVENTORY-003 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O |
| CAP-INVENTORY-004 | not-assessed | baseline-mapped | S/H/G | PSR-02,14 | P/O |
| CAP-INVENTORY-005 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O |
| CAP-INVENTORY-006 | not-assessed | baseline-mapped | S/H/G | PSR-02,16 | P/O |
| CAP-INVENTORY-007 | not-assessed | baseline-mapped | S/H/G | PSR-02,11 | P/O |

### Discovery
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-DISCOVERY-001 | not-assessed | baseline-mapped | S/H/G | PSR-02,11 | P/O/S |
| CAP-DISCOVERY-002 | not-assessed | baseline-mapped | S/H/G | PSR-02,05 | P/O/S |
| CAP-DISCOVERY-003 | not-assessed | baseline-mapped | S/H/G | PSR-02,11 | P/O/S |
| CAP-DISCOVERY-004 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O/S |
| CAP-DISCOVERY-005 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O/S |
| CAP-DISCOVERY-006 | not-assessed | baseline-mapped | S/H/G | PSR-02,14 | P/O/S |
| CAP-DISCOVERY-007 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O/S |
| CAP-DISCOVERY-008 | not-assessed | controls-required | S/H/G | PSR-11,13,05 | P/O/S |

### Monitoring and Observability
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-MONITORING-001 | not-assessed | baseline-mapped | S/H/G; SO | PSR-07,08 | P/O/S |
| CAP-MONITORING-002 | not-assessed | baseline-mapped | S/H/G; SO | PSR-07,08 | P/O/S |
| CAP-MONITORING-003 | not-assessed | baseline-mapped | S/H/G; SO | PSR-07 | P/O/S |
| CAP-MONITORING-004 | not-assessed | baseline-mapped | S/H/G; SO | PSR-08 | P/O/S |
| CAP-MONITORING-005 | not-assessed | baseline-mapped | S/H/G; SO | PSR-07,16 | P/O/S |
| CAP-MONITORING-006 | not-assessed | baseline-mapped | S/H/G; SO,PSD | PSR-07 | P/O/S |
| CAP-MONITORING-007 | not-assessed | baseline-mapped | S/H/G; SO | PSR-07 | P/O/S |
| CAP-MONITORING-008 | not-assessed | baseline-mapped | S/H/G; SO | PSR-07 | P/O/S |

### Network
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-NETWORK-001 | not-assessed | baseline-mapped | S/H/G; SO | PSR-11,02 | P/O/S |
| CAP-NETWORK-002 | not-assessed | baseline-mapped | S/H/G; SO | PSR-11,02 | P/O/S |
| CAP-NETWORK-003 | not-assessed | baseline-mapped | S/H/G | PSR-11 | P/O/S |
| CAP-NETWORK-004 | not-assessed | baseline-mapped | S/H/G | PSR-11 | P/O/S |
| CAP-NETWORK-005 | not-assessed | baseline-mapped | S/H/G | PSR-11 | P/O/S |
| CAP-NETWORK-006 | not-assessed | baseline-mapped | S/H/G | PSR-11 | P/O/S |
| CAP-NETWORK-007 | not-assessed | baseline-mapped | S/H/G; SO | PSR-11,10 | P/O/S |
| CAP-NETWORK-008 | not-assessed | controls-required | H/G; MDP | PSR-13,11 | P/O/S |

### Topology
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-TOPOLOGY-001 | not-assessed | baseline-mapped | S/H/G | PSR-02,11 | P/O |
| CAP-TOPOLOGY-002 | not-assessed | baseline-mapped | S/H/G | PSR-02,11 | P/O |

### Print
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-PRINT-001 | not-assessed | baseline-mapped | S/H/G; SO | PSR-02 | P/O/S |
| CAP-PRINT-002 | not-assessed | baseline-mapped | S/H/G; SO | PSR-02 | P/O/S |
| CAP-PRINT-003 | not-assessed | baseline-mapped | S/H/G | PSR-02,15 | P/O/S |
| CAP-PRINT-004 | not-assessed | baseline-mapped | S/H/G; SO | PSR-08,02 | P/O/S |
| CAP-PRINT-005 | not-assessed | controls-required | H/G; PSD | PSR-15,07 | P/O/S |
| CAP-PRINT-006 | not-assessed | controls-required | H/G; MDP | PSR-13,15 | P/O/S |
| CAP-PRINT-007 | not-assessed | baseline-mapped | S/H/G | PSR-02,16 | P/O |
| CAP-PRINT-008 | not-assessed | controls-required | H/G; MDP | PSR-13 | P/O/S |

### Virtualization and Containers
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-VIRT-001 | not-assessed | baseline-mapped | S/H/G | PSR-02 | P/O/S |
| CAP-VIRT-002 | not-assessed | baseline-mapped | S/H/G | PSR-02,14 | P/O/S |
| CAP-VIRT-003 | not-assessed | baseline-mapped | S/H/G | PSR-02,14 | P/O/S |
| CAP-VIRT-004 | not-assessed | controls-required | H/G; MDP | PSR-13 | P/O/S |
| CAP-VIRT-005 | not-assessed | baseline-mapped | S/H/G | PSR-02,13 | P/O/S |
| CAP-VIRT-006 | not-assessed | controls-required | H/G; MDP | PSR-13,14 | P/O/S |

### Deployments and Change
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-DEPLOY-001 | not-assessed | baseline-mapped | S/H/G; MDP | PSR-14 | P/O/S |
| CAP-DEPLOY-002 | not-assessed | baseline-mapped | S/H/G; MDP | PSR-13 | P/O/S |
| CAP-DEPLOY-003 | not-assessed | baseline-mapped | S/H/G; MDP | PSR-13 | P/O/S |
| CAP-DEPLOY-004 | not-assessed | controls-required | H/G; MDP,PSD | PSR-13,04 | P/O/S |
| CAP-DEPLOY-005 | not-assessed | baseline-mapped | S/H/G; MDP,SO | PSR-13,08 | P/O/S |
| CAP-DEPLOY-006 | not-assessed | controls-required | H/G; MDP | PSR-09,13 | P/O/S |
| CAP-DEPLOY-007 | not-assessed | controls-required | H/G; MDP | PSR-13,14 | P/O/S |
| CAP-DEPLOY-008 | not-assessed | controls-required | H/G; MDP | PSR-13 | P/O/S |
| CAP-DEPLOY-009 | not-assessed | controls-required | H/G; MDP | PSR-13 | P/O/S |
| CAP-DEPLOY-010 | not-assessed | controls-required | H/G; MDP | PSR-13,11 | P/O/S |
| CAP-DEPLOY-011 | not-assessed | controls-required | H/G; MDP | PSR-13,06 | P/O/S |

### Automation and Response
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-AUTOMATION-001 | not-assessed | controls-required | H/G; MDP | PSR-13,14 | P/O/S |
| CAP-AUTOMATION-002 | not-assessed | controls-required | H/G; MDP | PSR-13 | P/O/S |
| CAP-AUTOMATION-003 | not-assessed | baseline-mapped | S/H/G; SO | PSR-08,16 | P/O |
| CAP-AUTOMATION-004 | not-assessed | controls-required | H/G; MDP | PSR-13,06 | P/O/S |
| CAP-AUTOMATION-005 | not-assessed | baseline-mapped | S/H/G; SO | PSR-08 | P/O/S |
| CAP-AUTOMATION-006 | not-assessed | baseline-mapped | S/H/G; SO | PSR-08,07 | P/O/S |

### Trust and Offline
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-TRUST-001 | not-assessed | controls-required | H/G; PSD | PSR-12 | P/O/S |
| CAP-TRUST-002 | not-assessed | controls-required | H/G | PSR-12 | P/O/S |
| CAP-TRUST-003 | not-assessed | controls-required | H/G | PSR-12,11 | P/O/S |
| CAP-TRUST-004 | not-assessed | controls-required | H/G; MDP | PSR-14 | P/O/S |
| CAP-TRUST-005 | not-assessed | baseline-mapped | S/H/G; MDP | PSR-14 | P/O/S |
| CAP-TRUST-006 | not-assessed | controls-required | H/G; MDP | PSR-14 | P/O/S |
| CAP-TRUST-007 | not-assessed | controls-required | H/G; PSD | PSR-14,10,17 | P/O/S |
| CAP-TRUST-008 | not-assessed | baseline-mapped | H/G | PSR-10 | P/O/S |
| CAP-TRUST-009 | not-assessed | baseline-mapped | H/G | PSR-10 | P/O/S |

### Protection and Recovery
| ID | Evidence | Sec/Gov | Profiles | PSR | Resp |
|----|----------|---------|----------|-----|------|
| CAP-PROTECT-001 | not-assessed | controls-required | H/G; PSD | PSR-05,15 | P/O/S |
| CAP-PROTECT-002 | not-assessed | baseline-mapped | S/H/G; PSD | PSR-09 | P/O/S |
| CAP-PROTECT-003 | not-assessed | baseline-mapped | S/H/G; PSD | PSR-09 | P/O/S |
| CAP-PROTECT-004 | not-assessed | baseline-mapped | S/H/G | PSR-09,13 | P/O/S |
| CAP-PROTECT-005 | not-assessed | controls-required | H/G | PSR-10 | P/O/S |
| CAP-PROTECT-006 | not-assessed | baseline-mapped | S/H/G | PSR-10 | P/O/S |
| CAP-PROTECT-007 | not-assessed | controls-required | H/G; PSD | PSR-15,16 | P/O/S |
| CAP-PROTECT-008 | not-assessed | baseline-mapped | S/H/G | PSR-10 | P/O |

> **Alignment-Bestätigung:** 94 Capabilities zugeordnet; keine Capability gelöscht/umbenannt/hochgestuft; Evidence durchgehend `not-assessed`; kein `compliant`/`requirement-satisfied`/`government-approved`. PSR-Zuordnung = Readiness-Relevanz, **nicht** BSI-Control-Erfüllung.

## Zusammenfassung

- **Domains:** 13 (Platform · Identity · Inventory · Discovery · Monitoring · Network · Topology · Print · Virtualization · Deployments · Automation · Trust · Protection — Topology als eigenständige Gruppe innerhalb der Network-nahen Domänen geführt).
  > **Domain-Zählkorrektur (CO-WP-029):** Die Summenangabe lautete „12", während dieselbe Aufzählung **13** Domänen benennt und die Capability-IDs **13** eindeutige Domänen-Präfixe tragen (`AUTOMATION`, `DEPLOY`, `DISCOVERY`, `IDENTITY`, `INVENTORY`, `MONITORING`, `NETWORK`, `PLATFORM`, `PRINT`, `PROTECT`, `TOPOLOGY`, `TRUST`, `VIRT`). Rein mechanische Korrektur der Summenangabe: **keine** Domäne hinzugefügt, entfernt, umbenannt oder neu zugeordnet; **keine** Capability verändert.
- **Capabilities gesamt:** 94.
  > **Zählkorrektur (CO-WP-004E):** Die frühere Zusammenfassung nannte „74". Eine deterministische Zeilenzählung (`grep -cE "^\| CAP-"`) ergibt **94** eindeutige Capability-Zeilen mit 94 eindeutigen IDs und ohne Duplikate. Keine Capability wurde hinzugefügt oder entfernt; nur die Summenangabe wurde auf den tatsächlichen Bestand korrigiert. Der Abgleich der übrigen „74"-Referenzen ist in `CO-WP-029` **abgeschlossen** (Commit `6afa7ab`) und in `CO-WP-031` deterministisch nachgeprüft; kein Dokument führt „74" mehr als autoritative Summe. Verbleibende „74"-Nennungen sind ausschließlich historische Review- und Lessons-Aufzeichnungen und bleiben als solche erhalten (`RISK-66` `closed`, HM-8).
- **Implementation Status:** 94× `not-implemented` (100 %).
- **Support Status:** 94× `not-supported` (100 %). Keine Integration `supported`.
- **Read/Write-Verteilung:** `read` und `platform` dominieren die frühen Meilensteine; `write-target` ausschließlich ab `Deploy`/`Extend`/`Automate` und **nie** im Observe-Meilenstein.
- **Roadmap-Verteilung (grob):** `target-observe` ~30 · `target-map` ~10 · `target-plan` ~6 · `target-deploy` ~14 · `target-automate` ~6 · `target-extend` ~8. (Näherung; verbindlich ist die Zeilenangabe.)

**Bestätigung:** Keine Runtime-Capability ist implementiert. Keine Integration ist `supported`. Keine technische Architektur oder Technologie wird durch diese Matrix ausgewählt. Es wurde keine ADR erzeugt.
