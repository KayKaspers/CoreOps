# CoreOps Concept v3.0 – Platform Foundation

> Document Status: Accepted Product Vision
> Technical Decision Status: Unconfirmed
> ADR Status: No ADR accepted by this document
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`)
> Source: Human-Maintainer-provided CoreOps Concept v3.0
> Source Form: Local read-only Markdown handoff

**Dieses Dokument beschreibt das akzeptierte Zielbild. Technologie-, Architektur- und Roadmapdetails gelten nur dort als verbindlich, wo sie durch spätere Foundation-Entscheidungen oder akzeptierte ADRs bestätigt wurden.**

Registriert durch Work Package `CO-WP-002 – Concept v3.0 Registration and Decision Classification` (docs-only). Redaktionelle Anpassungen beschränken sich auf Public Neutrality, konsistente Markdown-Formatierung, Entfernung der reinen Chat-Anrede und klare Statuskennzeichnung. Es wurden keine Features, Architekturentscheidungen oder technischen Festlegungen ergänzt oder verändert. Die begleitende Klassifikation liegt in [CONCEPT_DECISION_CLASSIFICATION.md](CONCEPT_DECISION_CLASSIFICATION.md).

## Ausgangskontext

CoreOps wird als neues Softwareprojekt vollständig nach dem **Nova Development Framework – NDF v1.0.0** entwickelt.

Verbindliche NDF-Quelle: `https://github.com/KayKaspers/Nova-Development-Framework`

Für dieses Projekt ist der veröffentlichte Tag `v1.0.0` die normative NDF-Basis. Das veränderliche `main` darf ergänzend geprüft werden, aber Änderungen gegenüber dem finalen Tag dürfen nicht ungeprüft als verbindlich übernommen werden.

> Registrierungs-Note: Die verbindliche CoreOps-Projektgovernance (CO-WP-Serie) behandelt `main` als **nicht normativ**. Diese Formulierung des Concepts ist als Quellenaussage erhalten und wird in der Decision Classification als Klärungspunkt geführt.

---

# 1. Rollen und Zusammenarbeit

## Nova

Nova ist verantwortlich für:

* Produktplanung
* Architektur
* Sicherheitskonzepte
* Scope-Kontrolle
* Roadmap
* ADR-Vorbereitung
* Work-Package-Planung
* Erstellung der Claude-Prompts
* Review der Claude-Rückmeldungen
* Bewertung mit:
  * GO
  * GO WITH NOTES
  * REWORK
  * SPLIT
  * STOP
* Fortschreibung von Project Brain, Risiken, Roadmap und Context Packs

## Implementation Agent

Claude beziehungsweise ein anderer Implementation Agent:

* bearbeitet genau ein freigegebenes Work Package,
* hält sich strikt an den erlaubten Scope,
* verändert keine nicht freigegebenen Dateien,
* führt geforderte Tests aus,
* dokumentiert Abweichungen,
* erstellt keinen Commit,
* führt keinen Push aus,
* erstellt keinen Tag oder Release,
* liefert immer eine strukturierte **Rückmeldung an Nova**,
* liefert zusätzlich eine **Compact Context Summary**.

## Human Maintainer

Der Human Maintainer entscheidet allein über:

* finale Freigaben,
* Staging,
* Commit,
* Push,
* Merge,
* Tags,
* Releases,
* produktive Deployments,
* irreversible oder privilegierte Aktionen.

## Verbindlicher Ablauf

```text
Intake
→ Klassifizierung
→ Blueprint beziehungsweise Prompt
→ Umsetzung
→ Rückmeldung an Nova
→ Nova Review
→ Human-Maintainer-Entscheidung
→ Commit
→ CI und Validierung
→ Follow-up
```

Kein Schritt darf übersprungen werden.

---

# 2. NDF-Arbeitsmodus

CoreOps wird als **neues Projekt** aufgebaut.

Der NDF Project Adapter für bestehende Projekte ist deshalb nicht der primäre Einstieg. Zu verwenden ist der NDF New Project Flow.

## Verbindliche NDF-Prinzipien

* Documentation First
* Architecture First
* Security First
* Privacy by Design
* Fail Closed
* AI creates, Humans approve
* kleine und typisierte Work Packages
* ein kohärentes Work Package pro Commit
* keine verdeckten Scope-Erweiterungen
* keine irreversible Aktion ohne menschliche Freigabe
* Read-only vor Write
* Preview vor Execute
* Plan vor Deployment
* Backup vor gefährlicher Änderung
* Verifikation nach jeder Änderung
* Auditierbarkeit
* Context Economy
* Skills-first Operating Mode
* Public Neutrality für öffentliche Repository-Inhalte

## Work-Package-Typen

Jedes Work Package besitzt genau einen primären NDF-Typ, beispielsweise:

* `review-only`
* `docs-only`
* `code-fix`
* `feature`
* `test-only`
* `security-baseline`
* `security-code-fix`
* `health-score-update`
* `ci-diagnostic`
* `release-prep`
* `destructive-blueprint`
* `destructive-implementation`

Jeder Claude-Prompt muss mindestens enthalten:

```text
Work Package:
Work Package Type:
Prompt Mode:
Ziel:
Scope:
Allowed Files:
Forbidden Files:
Tests Required:
Akzeptanzkriterien:
Sicherheitsgrenzen:
Rückmeldung an Nova:
Compact Context Summary:
```

## Prompt-Modi

### Full Prompt

Verbindlich für:

* Scope Lock
* Architekturentscheidungen
* ADRs
* Sicherheitsrichtlinien
* Deployment-Ausführung
* privilegierte Änderungen
* destructive Actions
* Releases
* unklare oder risikoreiche Aufgaben

### Standard Prompt

Für:

* klar abgegrenzte normale Work Packages
* dokumentierte Features
* überschaubare Integrationen
* gezielte Reviews
* kleinere kontrollierte Änderungen

### Short Prompt

Nur für:

* wiederkehrende Folgearbeiten,
* kleine Updates,
* klar standardisierte Aufgaben,
* vorhandene und aktuelle Context Packs.

Short Prompts dürfen nicht für Security, ADRs, Scope Lock, Releases, destructive Actions oder unklare Anforderungen verwendet werden.

## Skills-first Operating Mode

Vor jedem Work Package ist zu prüfen:

1. Welche lokalen NDF-Skills sind relevant?
2. Sind diese Skills für den Work-Package-Typ zulässig?
3. Welche Sicherheitsgrenzen besitzen sie?
4. Sind sie docs-only und fail-closed?
5. Welche projektlokalen Skills werden später für CoreOps benötigt?

Skills dürfen niemals eigenständig:

* Secrets lesen,
* private Daten übertragen,
* Netzwerkzugriffe durchführen,
* Code ausführen,
* Commits erzeugen,
* pushen,
* Tags setzen,
* Releases veröffentlichen,
* produktive Systeme verändern.

## Context Economy

Für CoreOps sind Context Packs verbindlich.

Ein Context Pack enthält nur:

* aktuellen Phasenstatus,
* letztes abgeschlossenes Work Package,
* nächstes Work Package,
* aktive Blocker,
* relevante akzeptierte ADRs,
* Sicherheitsgrenzen,
* bekannte Risiken,
* notwendige Quelldateien,
* kompakte Historie.

Nicht enthalten sein dürfen:

* Roh-Chatverläufe,
* Secrets,
* private Daten,
* vollständige Log-Dumps,
* Chain-of-Thought,
* irrelevante abgeschlossene Detailhistorie.

Jede Rückmeldung an Nova endet mit einer Compact Context Summary von ungefähr 5 bis 10 Zeilen.

---

# 3. Projektidentität

## Projektname

**CoreOps**

## Slogan

**One Dashboard. Controlled Operations.**

## Projektart

* universelle Operations Control Plane
* vollständig self-hosted
* modular
* herstellerunabhängig
* offline- und Air-Gap-fähig
* deutsch und englisch
* Docker-first
* API-first
* Security-first
* NDF-native

## Kurzbeschreibung

> CoreOps ist eine universelle, self-hosted und offline-fähige Operations Control Plane für Anwendungen, Deployments, Server, Virtualisierung, Container, Netzwerke, Drucker und verteilte Infrastruktur. Die Plattform verbindet Source of Truth, Monitoring, Topologie, Desired State, Policy, vertrauenswürdige Automatisierung und nachvollziehbare Änderungen in einer modularen Oberfläche.

---

# 4. Vision

CoreOps soll zum zentralen Cockpit für heterogene IT-Infrastrukturen werden.

Die Plattform soll erkennen und nachvollziehbar darstellen:

* welche Systeme vorhanden sind,
* wo sich diese Systeme befinden,
* wem sie zugeordnet sind,
* wie sie miteinander verbunden sind,
* welche Anwendungen darauf betrieben werden,
* welche Versionen installiert sind,
* wie gesund die Systeme sind,
* welche Abhängigkeiten bestehen,
* welcher Zustand gewünscht ist,
* welche Abweichungen existieren,
* welche Änderungen geplant sind,
* ob eine Änderung zulässig ist,
* wer sie freigegeben hat,
* wie sie ausgeführt wurde,
* ob sie erfolgreich war,
* wie der vorherige Zustand wiederhergestellt werden kann.

CoreOps soll den Administrator entlasten, ohne ihm die Kontrolle zu entziehen.

---

# 5. Allgemeine Einsetzbarkeit

CoreOps ist ausdrücklich **nicht exklusiv für Core-Produkte**.

Die Plattform soll allgemein einsetzbar sein für:

* Homelabs
* Selfhoster
* Vereine
* kleine und mittlere Unternehmen
* Bildungseinrichtungen
* Entwickler
* Content Creator
* IT-Abteilungen
* verteilte Standorte
* eingeschränkte Netzwerke
* Offline-Umgebungen
* Air-Gap-Umgebungen
* später MSPs und Rechenzentren

Die Core-Produkte bilden hochwertige First-Party-Integrationen, sind aber weder Voraussetzung noch alleiniger Mittelpunkt des Projekts.

Beispiele:

* SpeakCore
* CastCore
* AirCore
* OrgaCore beziehungsweise SC-OrgaBase
* zukünftige Core-Produkte

---

# 6. Primäre Ziele

CoreOps soll langfristig folgende Aufgaben erfüllen:

1. Infrastruktur erkennen und inventarisieren.
2. Anwendungen und Dienste katalogisieren.
3. Systeme, Netzwerkgeräte und Drucker überwachen.
4. Netzwerktopologien automatisch erkennen und visualisieren.
5. physische und logische Abhängigkeiten abbilden.
6. Ist- und Sollzustände vergleichen.
7. Konfigurations- und Versionsabweichungen erkennen.
8. Deployments kontrolliert planen und durchführen.
9. Updates und Wartungen koordinieren.
10. Skripte und Workflows sicher ausführen.
11. Änderungen durch Policies und Genehmigungen absichern.
12. Artefakte, Plugins und Skripte kryptografisch prüfen.
13. Ereignisse korrelieren und Incidents unterstützen.
14. Konfigurationen versioniert sichern.
15. vollständig ohne externe Cloudpflicht arbeiten.
16. zukünftige Hersteller und Produkte modular integrieren.

---

# 7. Nicht-Ziele

CoreOps soll zumindest in den frühen Phasen nicht versuchen, folgende Produktklassen vollständig zu ersetzen:

* vollständiges CI-Build-System
* vollständiges SIEM
* vollständiges ITSM- oder Helpdesk-System
* Endpoint Detection and Response
* Antivirus
* allgemeine Remote-Desktop-Plattform
* vollständige Public-Key-Infrastruktur
* vollständiger Low-Code-App-Builder
* universelle Cloud-Management-Suite
* autonomer KI-Administrator
* vollständiger Ersatz für jedes Herstellerinterface
* vollständige ERP-, Lizenz- oder Vertragsverwaltung

CoreOps darf mit solchen Systemen integriert werden, ohne deren gesamten Funktionsumfang nachzubauen.

---

# 8. Zentrale Produktformel

```text
Source of Truth
+ Observed State
+ Desired State
+ Policy
+ Trusted Execution
+ Verification
+ Audit
= CoreOps
```

CoreOps soll nicht einfach Befehle an entfernte Systeme senden.

Jede relevante Änderung folgt grundsätzlich:

```text
Erkennen
→ Ist-Zustand erfassen
→ Soll-Zustand bestimmen
→ Abweichung bewerten
→ Plan erstellen
→ Policy prüfen
→ Preview anzeigen
→ Genehmigung einholen
→ Aktion ausführen
→ Health Check durchführen
→ Ergebnis verifizieren
→ Audit und Evidence Pack erzeugen
```

---

# 9. Architekturprinzipien

## 9.1 Modularer Monolith

CoreOps beginnt als modularer Monolith mit klaren internen Modulgrenzen.

Separat betrieben werden dürfen:

* Web-Frontend
* Control-Plane-Backend
* Background Worker
* CoreOps Agent
* CoreOps Site Relay
* Observability-Komponenten

Microservices werden nicht vorsorglich eingeführt.

Eine spätere Auslagerung erfolgt nur bei:

* nachgewiesenem Skalierungsbedarf,
* klarer Sicherheitsgrenze,
* unabhängigem Releasebedarf,
* abweichendem Ressourcenprofil,
* technischer Notwendigkeit.

## 9.2 API-first

Alle wesentlichen Funktionen werden über versionierte APIs modelliert.

Die Benutzeroberfläche darf keine privilegierten internen Abkürzungen verwenden, die für API-Nutzer nicht gelten.

## 9.3 Read-only First

Jede neue Integration beginnt mit:

* Discovery
* Inventar
* Health
* Monitoring
* Versionen
* Capabilities
* Topologieinformationen

Schreibende Aktionen werden in separaten Work Packages geplant.

## 9.4 Standards First

Offene Standards und dokumentierte APIs werden bevorzugt.

Herstellerspezifische CLI-Automatisierung ist nur zulässig, wenn:

* keine stabile API verfügbar ist,
* Befehle strikt validiert werden,
* unterstützte Versionen dokumentiert sind,
* Tests und Fixtures vorhanden sind,
* Fehler- und Rollbackpfade geklärt sind.

## 9.5 Offline First

Keine Kernfunktion darf zwingend einen externen Cloudservice benötigen.

## 9.6 Fail Closed

Bei folgenden Problemen werden privilegierte Aktionen blockiert:

* Policy nicht erreichbar
* Policy nicht eindeutig
* Identität nicht verifiziert
* Artefakt nicht verifiziert
* Signatur ungültig
* Zielzustand unklar
* Backupvoraussetzung nicht erfüllt
* Health-Zustand unbekannt
* Integration nicht kompatibel
* notwendige Genehmigung fehlt

---

# 10. Plane-Architektur

CoreOps wird konzeptionell in sechs Ebenen gegliedert.

```text
CoreOps
├── Experience Plane
├── Control Plane
├── Data Plane
├── Execution Plane
├── Trust Plane
└── Domain Packs
```

---

# 11. Experience Plane

Die Experience Plane bildet die sichtbare Benutzeroberfläche.

## Bestandteile

* zentrales Dashboard
* anpassbare Widgets
* globale Suche
* Asset Explorer
* Servicekatalog
* Deployment Center
* Maintenance Center
* Network Center
* Print Center
* Topology Explorer
* Incident Center
* Approval Inbox
* Audit Explorer
* Grafana Workspace
* integrierte Dokumentation
* Setup Wizard
* Recovery Mode
* mobile PWA-Ansichten

## Dashboard-Profile

Beispiele:

* Global Overview
* Infrastructure Operations
* Network Operations
* Print Operations
* Deployments
* Updates and Maintenance
* Security and Compliance
* Backup and Recovery
* Core Products
* Standortansicht
* Mobile Quick Actions
* NDF Governance

## Widget-Typen

* Statuskarte
* Ampel
* Einzelwert
* Zeitreihe
* Kapazitätsanzeige
* Tabelle
* Ereignisliste
* Warnung
* Topologieausschnitt
* Deploymentstatus
* Jobstatus
* Approval
* Schnellzugriff
* Grafana-Panel
* Dokumentationshinweis
* NDF-Status
* Versionsübersicht
* Drift-Anzeige

Widgets müssen:

* verschiebbar,
* skalierbar,
* filterbar,
* rollenabhängig,
* standortabhängig,
* benutzerbezogen speicherbar,
* DE/EN-fähig,
* barrierearm

sein.

---

# 12. Control Plane

Die Control Plane entscheidet, was CoreOps tun darf und welcher Zustand angestrebt wird.

## Bestandteile

* Identity
* RBAC
* Policy Engine
* Desired State
* Deployment Management
* Workflow Management
* Approval Engine
* Maintenance Windows
* Change Management
* Incident Coordination
* Audit
* Configuration Management
* Integration Registry
* Plugin Governance

Die Control Plane darf nicht durch Plugins oder Agents umgangen werden.

---

# 13. Data Plane

Die Data Plane speichert und normalisiert die ermittelten Informationen.

## Bestandteile

* Asset Inventory
* Source-of-Truth-Daten
* Feld-Provenance
* Observed State
* Telemetrie
* Logs
* Events
* Topologiegraph
* historische Zustände
* Deploymenthistorie
* Configuration Vault
* Auditdaten
* Evidence Packs
* IPAM- und DDI-Daten
* Zertifikatsinventar

## Grundregel

Rohdaten und normalisierte Daten müssen unterscheidbar bleiben.

Beispiel:

```text
Hersteller-Rohwert:
deviceCpuLoad = 72

Normalisierter CoreOps-Wert:
coreops.system.cpu.utilization = 0.72
```

---

# 14. Execution Plane

Die Execution Plane führt freigegebene Aufgaben aus.

## Bestandteile

* CoreOps Agents
* Site Relays
* Connectors
* Runner
* Worker
* Deployment Targets
* Script Runner
* Workflow Runtime
* Artifact Transfer
* Offline Queue
* Store-and-Forward

## Grundregel

Die Execution Plane entscheidet nicht selbst, ob eine privilegierte Aktion zulässig ist.

Sie führt nur Aktionen aus, für die eine gültige, signierte und noch nicht abgelaufene Ausführungsfreigabe der Control Plane vorliegt.

---

# 15. Trust Plane

Die Trust Plane ist eine eigene, übergreifende Sicherheitsgrenze.

## Bestandteile

* Secrets Management
* interne PKI
* Machine Identity
* Agent Enrollment
* mTLS
* Zertifikatsrotation
* Artifact Verification
* Signaturen
* Hashes
* SBOM
* Build Provenance
* Publisher Trust
* Widerruf
* Plugin Trust
* Policy Trust
* CorePack Verification

Sicherheit darf nicht nur als einzelnes Modul behandelt werden.

---

# 16. Domain Packs

Geräte- und Plattformunterstützung wird in Domain Packs gegliedert.

## Geplante Domain Packs

* Generic Linux
* Windows
* Proxmox
* Docker und Podman
* Kubernetes
* Hyper-V
* VMware
* Network
* Print
* Storage und NAS
* DDI
* Certificates
* UPS und Power
* IoT und Edge
* Core Products
* Generic HTTP/API
* Generic SNMP
* Generic MQTT

Ein Domain Pack kann enthalten:

* Discovery Provider
* Inventory Provider
* Monitoring Provider
* Topology Provider
* Deployment Provider
* Action Provider
* Backup Provider
* Widget Pack
* Grafana Pack
* Alert Pack
* Workflow Pack
* Dokumentation
* Test Fixtures

---

# 17. CoreOps Integration Contract – OIC

Der **CoreOps Integration Contract**, kurz OIC, ist der allgemeine Integrationsstandard der Plattform.

Er ist nicht auf Core-Produkte begrenzt.

## Der OIC definiert

* Integrationsidentität
* Hersteller
* Produktfamilie
* unterstützte Modelle
* unterstützte Versionen
* Capability-Modell
* Authentifizierungsverfahren
* Discovery-Verhalten
* Inventory-Schema
* Health-Schema
* Monitoring-Schema
* Topology-Schema
* Action-Schema
* Deployment-Schema
* Fehlerformate
* Timeout-Verhalten
* Idempotenz
* Audit-Korrelation
* Offline-Fähigkeiten
* Sicherheitsgrenzen
* Kompatibilitätsregeln

## Beispiel-Capabilities

```text
asset.discover
asset.inventory.read
health.read
metrics.read
logs.read

topology.neighbors.read
topology.interfaces.read
topology.routes.read

configuration.read
configuration.backup
configuration.validate
configuration.deploy

deployment.plan
deployment.execute
deployment.verify
deployment.rollback

update.scan
update.plan
update.execute

printer.status.read
printer.supplies.read
printer.counters.read
printer.queue.read
printer.queue.manage

network.interface.read
network.interface.configure
network.poe.cycle
network.vlan.configure
```

Eine Capability darf nur angeboten werden, wenn:

1. das Zielsystem sie unterstützt,
2. der Adapter sie implementiert,
3. die Integration für diese Version getestet wurde,
4. die Benutzerrolle sie erlaubt,
5. die Policy sie freigibt.

---

# 18. Integrations- und Qualitätsstufen

## Funktionale Integrationslevel

### Level 0 – Detected

* erreichbar
* grob klassifiziert

### Level 1 – Inventoried

* Hersteller
* Modell
* Version
* Basisinventar

### Level 2 – Observed

* Monitoring
* Health
* Alerts
* Metriken

### Level 3 – Protected

* Konfigurationsbackup
* Zustandsbackup
* Last-Known-Good

### Level 4 – Managed

* klar begrenzte Verwaltungsaktionen

### Level 5 – Automated

* Deployments
* Firmware
* Workflows
* Rollback

## Supportstatus

* Experimental
* Preview
* Supported
* Deprecated
* Blocked
* Revoked

## Herausgeberklassen

* Built-in
* First-Party Pack
* Verified Partner Pack
* Community Pack
* Local Pack

Die Benutzeroberfläche muss klar anzeigen, welches Level und welcher Supportstatus tatsächlich gilt.

---

# 19. Source of Truth und Provenance

CoreOps erhält ein feldbasiertes Source-of-Truth-Modell.

Ein Asset kann Informationen aus mehreren Quellen erhalten:

* Agent
* SNMP
* Hersteller-API
* DNS
* DHCP
* Active Directory
* Hypervisor
* NetBox
* manueller Eintrag
* CSV- oder JSON-Import
* Discovery Scan

## Feldstatus

* discovered
* imported
* declared
* calculated
* manually confirmed
* inherited
* overridden
* stale
* conflicting
* locked

## Beispiel

```yaml
hostname:
  value: switch-core-01
  source: dns
  confidence: high
  authoritative: false
  discoveredAt: 2026-07-12T12:00:00Z

displayName:
  value: Core Switch Site A
  source: manual
  confirmedBy: human-maintainer
  locked: true
```

> Public-Neutrality-Note: Der ursprüngliche private Standortname im Beispiel `displayName` wurde öffentlich neutral durch `Core Switch Site A` ersetzt. Bedeutung und technische Aussage bleiben unverändert.

## Konfliktregeln

CoreOps benötigt pro Feldtyp definierte Prioritäten.

Beispiel:

```text
manuell gesperrter Wert
→ externe autoritative Source of Truth
→ authentifizierte Hersteller-API
→ vertrauenswürdiger Agent
→ SNMP
→ Discovery-Heuristik
```

Diese Reihenfolge darf je Datenklasse unterschiedlich sein.

## NetBox

CoreOps soll ein integriertes Basisinventar sowie ein minimales IPAM besitzen.

Für größere Installationen ist eine optionale NetBox-Integration vorzusehen.

Ein bidirektionaler Sync darf nur mit:

* festgelegter Datenhoheit,
* Preview,
* Konfliktanzeige,
* Audit,
* begrenzten Objektklassen

erfolgen.

---

# 20. Observed State, Desired State und Drift

CoreOps unterscheidet strikt zwischen:

## Observed State

Der tatsächlich ermittelte Zustand.

## Desired State

Der deklarierte gewünschte Zustand.

## Effective State

Der nach Policies, Ausnahmen und Overrides aktuell anzustrebende Zustand.

## Drift

Die Abweichung zwischen beobachtetem und gewünschtem Zustand.

## Drift-Arten

* Version Drift
* Configuration Drift
* Infrastructure Drift
* Policy Drift
* Secret Drift
* Firmware Drift
* Network Drift
* Topology Drift
* Inventory Drift
* Certificate Drift

## Betriebsmodi

### Observe

Drift nur erkennen und anzeigen.

### Recommend

Korrekturplan erzeugen.

### Approve

Korrektur nach menschlicher Freigabe ausführen.

### Enforce

Automatische Korrektur in engen, vorab genehmigten Grenzen.

### Frozen

Keine automatische Änderung zulässig.

Für frühe CoreOps-Versionen sind nur `Observe`, später `Recommend` und anschließend `Approve` vorgesehen.

`Enforce` ist kein MVP-Ziel.

---

# 21. Policy Engine

RBAC allein reicht nicht aus.

Eine Aktion wird kontextabhängig bewertet.

## Mögliche Policy-Eingaben

* Benutzer
* Rolle
* Zielsystem
* Standort
* Umgebung
* Kritikalität
* Wartungsfenster
* Backupstatus
* Health Score
* Anzahl der Ziele
* Deploymentstrategie
* Artefaktvertrauen
* Signaturstatus
* bekannte Incidents
* Rollback-Verfügbarkeit
* Tageszeit
* notwendige Genehmigungen
* Integration Support Level

## Beispiel

```text
Ein Operator darf ein Docker-Deployment ausführen, wenn:

- das Ziel nicht Production ist,
- das Artefakt verifiziert ist,
- ein Health Check definiert ist,
- maximal zwei Ziele gleichzeitig geändert werden,
- kein kritischer Incident aktiv ist.
```

Für Production zusätzlich:

```text
- genehmigtes Wartungsfenster
- aktuelles und geprüftes Backup
- zweiter Genehmiger
- verfügbare Last-Known-Good-Version
- erfolgreiches Staging-Deployment
```

Die Policy Engine arbeitet fail-closed.

Eine mögliche Technologie ist Open Policy Agent. Die endgültige Entscheidung erfolgt durch ADR und Technologievergleich.

---

# 22. Deployment Management

CoreOps soll kontrollierte Deployments durchführen können.

## Deployment-Arten

* Docker Compose
* Podman
* Kubernetes Manifeste
* Helm
* Linux-Dienste
* Windows-Pakete
* Webanwendungen
* Backend-Dienste
* Core-Produkte
* Proxmox VMs
* Proxmox LXC
* Cloud-Init
* Hyper-V Templates
* VMware Templates
* Terraform beziehungsweise OpenTofu
* Ansible
* Netzwerkkonfigurationen
* Druckerwarteschlangen
* Firmware

## Build und Deployment bleiben getrennt

CoreOps soll zunächst kein vollständiges CI-Build-System werden.

Bevorzugter Ablauf:

```text
CI baut und testet
→ Registry speichert unveränderliches Artefakt
→ CoreOps prüft Artefakt und Provenance
→ CoreOps plant Deployment
→ Human Maintainer genehmigt
→ CoreOps verteilt
→ CoreOps verifiziert
```

## Verbotene Produktionsmuster

* unbestimmter `latest`-Tag
* veränderlicher Git-Branch als Produktivquelle
* nicht verifiziertes Skript aus dem Internet
* Artefakt ohne Digest
* Deployment ohne Health Check
* Deployment ohne bekannten Zielzustand
* stiller Fallback auf eine andere Version

## Deployment Blueprint

Ein Blueprint definiert:

* Name
* Version
* Artefakt
* Digest
* Herausgeber
* Zielumgebung
* Zielsysteme
* Abhängigkeiten
* Konfiguration
* Secret-Referenzen
* Vorbedingungen
* Backupanforderung
* Deploymentstrategie
* maximale Parallelität
* Health Checks
* Smoke Tests
* Rollback
* Wartungsfenster
* Policy-Anforderungen
* Genehmigungen

## Strategien

* Recreate
* Rolling
* Canary
* Blue/Green
* Standortwellen
* Gruppenwellen
* manuell freigegebene Wellen

## Promotion

```text
Development
→ Test
→ Staging
→ Production
```

Zwischen Umgebungen wird dasselbe unveränderliche Artefakt promoviert.

---

# 23. Artifact Trust und Supply Chain

CoreOps wird Teil der Software-Lieferkette und benötigt daher ein eigenes Vertrauensmodell.

## Für Artefakte zu speichern

* Version
* Digest
* Signatur
* Herausgeber
* Build-Provenance
* SBOM
* unterstützte Plattformen
* bekannte Abhängigkeiten
* Security-Status
* Widerrufsstatus
* Importquelle
* Importzeitpunkt
* verifizierende Instanz

## CoreOps Trust

```text
CoreOps Trust
├── Artifact Verification
├── Signatures
├── Provenance
├── SBOM
├── Publisher Trust
├── Revocation
├── Vulnerability Status
└── Trust Policies
```

Mögliche Standards und Technologien werden später per ADR bewertet, beispielsweise:

* Sigstore beziehungsweise Cosign
* CycloneDX
* SLSA-Provenance
* TUF-ähnliches Repository-Vertrauen

Keine Technologie wird ohne dokumentierten Vergleich verbindlich festgelegt.

---

# 24. Monitoring und Observability

## Empfohlener Stack

* Prometheus
* Grafana
* OpenTelemetry Collector
* Loki
* Alertmanager oder eigenes Notification Routing
* optional später Tracing

## Zwei Darstellungsebenen

### Native CoreOps-Widgets

Für:

* tägliche Übersicht
* rollenbasierte Darstellung
* mobile Ansicht
* einheitliches Design
* schnelle Bedienung

### Grafana Workspace

Für:

* tiefe technische Analyse
* flexible Dashboards
* Zeitreihen
* Logs
* historische Untersuchungen

Grafana ist Analysewerkzeug und nicht die einzige Oberfläche von CoreOps.

## Telemetrie-Normalisierung

CoreOps erhält ein eigenes semantisches Schema.

Beispiele:

```text
coreops.system.cpu.utilization
coreops.system.memory.utilization
coreops.network.interface.errors
coreops.printer.supplies.level
coreops.deployment.duration
coreops.agent.heartbeat.age
```

Hersteller-Rohdaten bleiben zusätzlich verfügbar.

---

# 25. CoreScore

Assets, Anwendungen und Instanzen können einen zusammengesetzten Gesundheitswert erhalten.

Mögliche Teilbereiche:

* Erreichbarkeit
* Dienstzustand
* Ressourcen
* Updates
* Backups
* Sicherheit
* Integration
* Zertifikate
* Drift
* Datenaktualität

Der CoreScore muss zusätzlich anzeigen:

* Datenabdeckung
* letzte Aktualisierung
* nicht unterstützte Prüfungen
* ignorierte Findings
* manuelle Overrides
* Datenquellen

Fehlende Daten dürfen nicht automatisch als gesund gewertet werden.

---

# 26. Netzwerkmanagement

## Geräteklassen

* Managed Switches
* Router
* Firewalls
* Access Points
* WLAN Controller
* VPN Gateways
* Load Balancer
* virtuelle Switches
* virtuelle Router
* Proxmox Bridges
* Netzwerk-Appliances
* USV und Environmental Monitoring

## Standards und Protokolle

* ICMP
* SNMPv3
* LLDP
* CDP
* Syslog
* SSH
* NETCONF
* RESTCONF
* gNMI
* REST APIs
* sFlow
* NetFlow
* IPFIX

## Geplante Herstellerfamilien

Frühe Priorität:

* Ubiquiti UniFi
* MikroTik
* Cisco
* HPE Aruba
* TP-Link Omada
* OPNsense
* pfSense
* AVM
* Netgear

Spätere Erweiterung:

* Juniper
* Fortinet
* Palo Alto Networks
* Sophos
* Extreme Networks
* Dell Networking
* Huawei
* Ruckus
* Arista

Eine Herstellerangabe allein reicht nicht als Supportversprechen.

Jede Integration dokumentiert:

* getestete Modelle
* getestete Firmware
* unterstützte Capabilities
* benötigte Rechte
* bekannte Grenzen
* letztes Testdatum

## Netzwerkmonitoring

* Erreichbarkeit
* Latenz
* Paketverlust
* CPU
* RAM
* Temperatur
* Lüfter
* Netzteile
* Uptime
* Firmware
* Linkstatus
* Geschwindigkeit
* Duplex
* Fehler
* Drops
* Interfaceauslastung
* SFP-Werte
* PoE
* VLANs
* Trunks
* STP
* LACP
* MAC-Tabellen
* Routing
* VPN
* WLAN
* Access Points
* SSIDs
* Clientzahlen

## Schreibende Netzwerkaktionen

Später:

* Port aktivieren oder deaktivieren
* PoE neu starten
* Portbeschreibung setzen
* VLAN-Zuordnung ändern
* Konfiguration sichern
* Konfigurationsvorlage ausrollen
* Firmware deployen
* VPN neu aufbauen

Jede Änderung benötigt:

```text
Ist-Zustand
→ Konfigurationsbackup
→ Diff
→ Validierung
→ Impact Analysis
→ Preview
→ Genehmigung
→ Ausführung
→ Erreichbarkeitsprüfung
→ erneutes Auslesen
→ Verifikation
```

---

# 27. Druckermanagement

## Geplante Hersteller

* HP
* Brother
* Canon
* Epson
* Kyocera
* Xerox
* Lexmark
* Ricoh
* Konica Minolta
* Sharp
* OKI

## Standards

* SNMP
* Printer MIB
* IPP
* IPPS
* mDNS
* DNS-SD
* CUPS
* Windows Print Management
* Hersteller-APIs

## Monitoring

* Erreichbarkeit
* Gesamtstatus
* Papierstau
* Papierfachstatus
* Toner
* Tinte
* Trommel
* Wartungskit
* Resttonerbehälter
* Seitenzähler
* Farbseiten
* Schwarz-Weiß-Seiten
* Duplex
* Seriennummer
* Modell
* Firmware
* Warnungen

## Printserver

* Windows Print Server
* CUPS
* Linux Printserver

## Funktionen

* Warteschlangen anzeigen
* blockierte Aufträge erkennen
* Warteschlange pausieren
* Warteschlange fortsetzen
* Aufträge kontrolliert abbrechen
* Treiberstand anzeigen
* Ports dokumentieren
* Testseite auslösen
* Druckerzuordnungen bereitstellen

Druckauftragsnamen und Benutzerdaten sind als potenziell personenbezogen zu behandeln.

---

# 28. Netzwerktopologie

CoreOps soll Netzwerktopologien automatisch erkennen, berechnen, anzeigen und historisieren.

## Datenquellen

* LLDP
* CDP
* SNMP-Interfaces
* MAC-Tabellen
* ARP
* IPv6 Neighbor Discovery
* Routingtabellen
* VLANs
* Spanning Tree
* LACP
* WLAN Controller
* DHCP
* DNS
* Hypervisor APIs
* virtuelle Switches
* Container-Netzwerke
* Kubernetes-Netzwerke
* Agents
* manuelle Verbindungen
* Herstellercontroller

## Topologieobjekte

* Standort
* Gebäude
* Raum
* Rack
* Gerät
* Interface
* physische Verbindung
* logische Verbindung
* VLAN
* VRF
* Subnetz
* Route
* WLAN
* VPN
* Dienst
* Anwendung
* Abhängigkeit

## Evidence Model

Jede erkannte Verbindung speichert:

* Quelle
* Zeitpunkt
* Quellgerät
* Zielgerät
* Quellinterface
* Zielinterface
* Protokoll
* Confidence Score
* Bestätigung
* widersprüchliche Evidenz
* manuelle Overrides

## Confidence-Level

* bestätigt
* hoch
* mittel
* niedrig
* manuell
* widersprüchlich

CoreOps darf eine abgeleitete Verbindung nicht als zweifelsfreie Tatsache darstellen.

## Ansichten

* physische Topologie
* Layer 2
* Layer 3
* VLAN
* VRF
* Virtualisierung
* Container
* Kubernetes
* Services
* Anwendungen
* Standorte
* WAN
* VPN
* WLAN

## Interaktion

* Zoom
* Verschieben
* Filter
* Ebenen
* gespeicherte Layouts
* automatische Layouts
* Live-Status
* Metrik-Overlay
* Alert-Overlay
* historische Ansicht
* Export als SVG, PNG und PDF

## Path Explorer

Der Benutzer wählt Quelle und Ziel.

CoreOps zeigt den bekannten oder vermuteten Pfad.

## Impact Analysis

Vor Wartung oder Deployment zeigt CoreOps:

* betroffene Geräte
* betroffene Ports
* betroffene VLANs
* betroffene Dienste
* abhängige Anwendungen
* alternative Pfade
* geschätzte Auswirkungen

---

# 29. Discovery Engine

## Passive Discovery

* LLDP
* CDP
* SNMP Traps
* Syslog
* DHCP
* Controllerdaten
* Agentmeldungen

## Aktive Discovery

* ICMP
* TCP-Prüfungen
* SNMP
* API-Abfragen
* SSH-Banner
* IPP
* mDNS
* DNS

## Import Discovery

* CSV
* JSON
* YAML
* NetBox
* Proxmox
* VMware
* Hyper-V
* UniFi
* Omada
* Active Directory
* DHCP
* DNS

## Sicherheitsregeln

Scans benötigen:

* definierte Netzbereiche
* Ausschlusslisten
* Rate Limits
* maximale Parallelität
* Timeouts
* Benutzerberechtigung
* Audit
* optional Wartungsfenster

CoreOps darf nicht ungefragt alle erreichbaren Netze scannen.

---

# 30. CoreOps Agent und Site Relay

## Agentless

Geeignet für:

* SSH
* WinRM
* SNMP
* REST
* Proxmox API
* Docker API
* Kubernetes API

## Agent-basiert

Geeignet für:

* Systeme hinter NAT
* dynamische IPs
* lokale Skriptausführung
* Offline Queues
* präzise Inventarisierung
* Windows
* entfernte Standorte
* eingeschränkte Netze

## Site Relay

Ein Site Relay bündelt:

* lokale Agents
* Discovery
* Telemetrie
* Artifact Cache
* Jobausführung
* Store-and-Forward
* lokale Policies
* Offlinebetrieb

## Verbindung

Agents und Relays bauen bevorzugt eine ausgehende mTLS-Verbindung zur Control Plane auf.

## Sicherheitsgrenzen

* eindeutige Maschinenidentität
* kurzlebige Zertifikate
* Capability-Allowlist
* lokale Policy
* Widerruf
* verschlüsselte Queue
* keine allgemeine Remote-Root-Shell
* keine ungeprüften Kommandos
* keine automatische Privilegienausweitung

Als Agent-Technologie ist Go ein möglicher Kandidat. Die Entscheidung erfolgt erst nach ADR und Prototypvergleich.

---

# 31. Maschinenidentitäten

Technische Komponenten sollen keine dauerhaft gültigen universellen API-Tokens verwenden.

## Identitätsobjekte

* Agent
* Relay
* Worker
* Connector
* Plugin
* Collector
* Deployment Target

## Anforderungen

* eindeutige Identität
* kurze Laufzeit
* Rotation
* Widerruf
* minimale Capabilities
* Standortbindung
* Assetbindung
* Audit-Korrelation

## MVP

* interne CoreOps-PKI
* mTLS
* Enrollment
* Zertifikatsrotation

## Später

* externe PKI
* TPM
* Hardware Attestation
* optional SPIFFE-kompatibles Modell

---

# 32. Skript-System

Unterstützte Artefakte:

* Bash
* Shell
* Python
* PowerShell
* Node.js
* Ansible
* Terraform
* OpenTofu
* Container Runner
* signierte Binärdateien

## Skriptmetadaten

* ID
* Name
* Beschreibung
* Version
* Autor
* Herausgeber
* Kategorie
* Tags
* Plattformen
* benötigte Rechte
* Capabilities
* Timeout
* Parameter
* Changelog
* Hash
* Signatur
* Herkunft
* Freigabestatus
* Reviewstatus

## Lebenszyklus

```text
Imported
→ Quarantine
→ Reviewed
→ Approved
→ Active
→ Deprecated
→ Revoked
```

Externe Skripte dürfen nicht direkt nach Download ausgeführt werden.

Produktive Skripte müssen an:

* Commit
* Tag
* Digest
* Hash
* signierten Release

gebunden sein.

---

# 33. Workflow Engine

Workflows können langfristig über Stunden oder Tage laufen.

## Anforderungen

* persistenter Zustand
* Wiederaufnahme nach Neustart
* idempotente Schritte
* Retries
* Timeouts
* manuelle Gates
* Genehmigungen
* Kompensation
* Rollback
* externe Signale
* versionierte Definitionen

## Knotentypen

* Condition
* Approval
* Delay
* Script
* API Request
* SSH Task
* WinRM Task
* Backup
* Update
* Deployment
* Restart
* Health Check
* Notification
* Manual Gate
* Rollback
* Parallel Group
* Sequence
* For Each Target

Ein visueller Drag-and-Drop-Editor wird erst nach Stabilisierung des deklarativen Workflow-Schemas entwickelt.

Vor einer eigenen Workflow Engine sind Alternativen wie Temporal zu bewerten.

---

# 34. Einheitliches Eventmodell

CoreOps benötigt ein versioniertes Eventmodell.

## Eventfelder

* Event-ID
* Event-Typ
* Quelle
* Zeitpunkt
* Empfangszeitpunkt
* Asset
* Standort
* Severity
* Correlation-ID
* Causation-ID
* Payload-Schema
* Schema-Version
* Workspace
* Vertrauensstufe
* Datenschutzklasse

Das Modell soll sich an etablierten Eventkonventionen orientieren.

## Eventquellen

* Agent
* SNMP Trap
* Syslog
* Webhook
* Deployment
* Workflow
* Backup
* Printer
* Topology
* Benutzer
* Core-Produkt
* Herstellercontroller

## Korrelation

CoreOps soll zusammenhängende Symptome erkennen können.

Beispiel:

```text
Storage-Latenz steigt
→ Datenbank-Timeout
→ CastCore Stream degraded
→ SpeakCore Backup fehlerhaft
```

---

# 35. Incident Response und Runbooks

## CoreOps Response

* Incidents
* Alert Correlation
* Runbooks
* Ownership
* Escalation
* Maintenance Mode
* Timeline
* Post-Incident Review

## Runbook-Arten

* reine Dokumentation
* interaktive Checkliste
* kontrollierter Workflow

Dokumentation allein darf nicht automatisch ausführbar sein.

## Maintenance Suppression

Während genehmigter Wartungen werden erwartete Alerts markiert oder unterdrückt.

Sie bleiben im Audit erhalten.

---

# 36. DDI und Zertifikate

## CoreOps DDI

* Subnetze
* IP-Adressen
* Reservierungen
* DHCP-Leases
* DNS-Zonen
* DNS-Records
* Reverse DNS
* freie Bereiche
* Konflikterkennung
* VLAN-Zuordnung

Frühe Versionen arbeiten primär read-only.

## CoreOps Certificates

* Zertifikatsinventar
* Ablaufdatum
* Aussteller
* SANs
* Kettenprüfung
* Zuordnung zu Diensten
* Erneuerungsstatus
* schwache Algorithmen
* interne und externe Zertifikate

CoreOps soll Zertifikatsworkflows unterstützen, aber nicht zwingend selbst Certificate Authority werden.

---

# 37. Configuration Vault

CoreOps speichert versionierte Konfigurationen von:

* Switches
* Routern
* Firewalls
* Druckern
* Docker Compose
* Kubernetes
* Anwendungen
* Betriebssystemen
* Workflows
* Deployment Blueprints

## Funktionen

* Snapshot
* Diff
* Redaction
* Last-Known-Good
* Wiederherstellungsplan
* Integritätsprüfung
* Zuordnung zu Change und Deployment

Ein vorhandenes Backup gilt nicht automatisch als wiederherstellbar.

## Recovery Readiness

* vorhanden
* Integrität geprüft
* entschlüsselbar
* kompatibel
* Test-Restore erfolgreich
* Wiederherstellungsanleitung vorhanden

---

# 38. Offline- und Air-Gap-Betrieb

Offline-Fähigkeit ist Foundation-Anforderung.

## Betriebsmodi

### Connected

Normale Internetverbindung.

### Restricted

Proxy, Allowlist oder zeitweise Verbindung.

### Isolated

Keine Internetverbindung, interne Repositories.

### Air-Gapped

Kontrollierter Datenträgertransfer.

## CorePack

Ein CorePack ist ein signiertes Offline-Transferpaket.

Es kann enthalten:

* CoreOps-Updates
* Agent-Versionen
* Plugins
* Skripte
* Container Images
* Helm Charts
* VM Templates
* Pakete
* Druckertreiber
* Firmware
* Dashboards
* Alert Rules
* SBOM
* Signaturen
* Prüfsummen
* Dokumentation

## Importprozess

```text
Transfer
→ Quarantäne
→ Signaturprüfung
→ Hashprüfung
→ Kompatibilitätsprüfung
→ Inhaltsvorschau
→ Human-Maintainer-Freigabe
→ lokaler Import
→ Audit
```

## Store-and-Forward

Agents und Relays puffern:

* Metriken
* Events
* Auditdaten
* Jobresultate
* Warnungen

Die spätere Übertragung benötigt:

* Zeitstempel
* Sequenznummer
* Deduplizierung
* Integritätsprüfung

---

# 39. Benutzer, Rollen und Rechte

## Standardrollen

* Owner
* Administrator
* Operator
* Maintainer
* Viewer
* Auditor
* Guest

## Begrenzungsebenen

Rechte müssen begrenzbar sein auf:

* Workspace
* Tenant
* Standort
* Gruppe
* Asset
* Produkt
* Capability
* Aktion
* Umgebung
* Zeitfenster

## Authentifizierung

* lokaler Owner
* OpenID Connect
* OAuth
* LDAP
* Active Directory
* 2FA
* Passkeys
* Recovery Codes
* Break-Glass Account

Der Break-Glass-Zugang:

* funktioniert lokal,
* wird besonders geschützt,
* wird vollständig auditiert,
* ist nicht für den normalen Betrieb vorgesehen.

---

# 40. Secrets Management

Secrets umfassen:

* SSH Keys
* Passwörter
* API Tokens
* Zertifikate
* Agent Credentials
* Webhooks
* Registry-Zugänge

## Grundregeln

* keine Secrets in Logs
* keine Secrets in Skriptdateien
* keine Secrets in Exporten
* keine Secrets in Jobparametern
* verschlüsselte Speicherung
* Rotation
* Ablauf
* Zugriffsprotokollierung
* minimale Sichtbarkeit
* Secret-Referenzen statt Klartext

## Backends

CoreOps benötigt einen integrierten verschlüsselten Secret Store.

Später optional:

* HashiCorp Vault
* Bitwarden Secrets Manager
* Kubernetes Secrets
* Enterprise Vaults

---

# 41. Datenschutz und Datenklassifizierung

## Datenklassen

* Public
* Internal
* Confidential
* Secret
* Personal Data
* Security Sensitive

## Potenziell sensible Daten

* IP-Adressen
* Hostnamen
* Domains
* Topologien
* Benutzernamen
* Seriennummern
* Standorte
* Druckauftragsnamen
* Auditdaten
* Logs
* Softwareversionen

## Redaction

Support Bundles und Exporte müssen abhängig von der Datenklasse:

* Secrets entfernen
* IP-Adressen anonymisieren
* Benutzernamen pseudonymisieren
* interne Domains ersetzen
* Seriennummern redigieren
* Druckauftragsnamen entfernen

## Retention

Separate Aufbewahrungsregeln für:

* Metriken
* Logs
* Audit
* Events
* Topologie-Snapshots
* Jobausgaben
* Artefakte
* Backups
* Incidents

---

# 42. CoreOps Self-Protection

CoreOps muss auch bei eigenen Teilstörungen kontrollierbar bleiben.

## Control-Plane-Health

* API
* Datenbank
* Redis
* Worker
* Agents
* Relays
* Grafana
* Prometheus
* Object Storage
* Secret Store
* Policy Engine
* Event Pipeline

## Degraded Modes

### Grafana ausgefallen

Native Übersicht bleibt verfügbar.

### Redis ausgefallen

Keine neuen Jobs, read-only Funktionen bleiben verfügbar.

### Worker ausgefallen

Laufende Aufgaben werden sicher pausiert oder wiederaufgenommen.

### Policy Engine ausgefallen

Privilegierte Aktionen werden blockiert.

### Prometheus ausgefallen

Letzter bekannter Zustand wird mit Zeitstempel angezeigt.

### Agent offline

Jobs werden nicht blind mehrfach verteilt.

### Internet ausgefallen

Lokale Funktionen bleiben verfügbar.

## Self-Dependency Protection

CoreOps darf nicht unkontrolliert die Infrastruktur abschalten, auf der es selbst läuft.

Benötigt werden:

* Self-Dependency-Erkennung
* Control-Plane-Schutz
* Management-Exclusion-Tags
* Quorum-Regeln
* Standortschutz
* Wartungswellen

---

# 43. Recovery Mode

Eine reduzierte Notfalloberfläche soll verfügbar bleiben.

## Funktionen

* Control-Plane-Status
* Audit
* Backup
* Restore
* Plugin deaktivieren
* Workerstatus
* Diagnoseexport
* Policy-Status
* Agentstatus
* Datenbankstatus

Keine normalen Deployments oder Automationen im Recovery Mode.

---

# 44. Integration Lab

Für Hersteller- und Gerätetests wird ein Integration Lab vorgesehen.

## Bestandteile

* simulierte SNMP-Geräte
* anonymisierte SNMP-Walks
* IPP-Testserver
* API-Mocks
* virtuelle Netzwerkgeräte
* LLDP/CDP-Fixtures
* aufgezeichnete Telemetrie
* fehlerhafte Antworten
* Timeouts
* verschiedene Firmwareprofile

## Golden Fixtures

Beispiel:

```text
fixtures/
├── generic-snmp/
├── hp-printer/
├── brother-printer/
├── unifi-switch/
├── mikrotik-router/
├── proxmox/
└── core-products/
```

---

# 45. KI-Unterstützung

KI ist langfristig nur beratend.

## Erlaubte Funktionen

* Incident-Zusammenfassungen
* Alert-Gruppierung
* Topologieerklärungen
* Trendanalyse
* Wartungsplanvorschläge
* Erklärung von Konfigurationsdiffs
* Suche in Dokumentation
* Workflow-Entwürfe
* Risiko-Hinweise

## Nicht ohne deterministische Prüfung und Freigabe

* Root-Kommandos
* Firewalländerungen
* Firmware-Deployment
* Löschung
* Restore
* produktive Deployments
* ungeprüfte Skriptausführung

Grundregel:

```text
KI darf analysieren, erklären, priorisieren und vorschlagen.

KI darf nicht autonom ändern, deployen, löschen oder wiederherstellen.
```

---

# 46. Vorgeschlagener Technologie-Stack

Alle Angaben sind zunächst Architekturhypothesen und müssen durch ADRs bestätigt werden.

## Frontend

* Next.js
* React
* TypeScript
* Tailwind CSS
* responsive Design
* PWA
* Light/Dark Mode
* DE/EN
* Accessibility Baseline

## Backend

* Node.js
* TypeScript
* Fastify
* REST API
* OpenAPI
* WebSocket oder Server-Sent Events
* modularer Monolith

## Datenhaltung

* PostgreSQL
* Redis
* optional S3-kompatibler Objektspeicher

SQLite höchstens für isolierte Tests, nicht als primäres Entwicklungsziel.

## Agent und Relay

* Kandidat: Go
* statisch paketierbar
* Linux und Windows
* geringer Ressourcenverbrauch
* ausgehende mTLS-Verbindung

## Observability

* Prometheus
* Grafana
* OpenTelemetry Collector
* Loki
* Alertmanager oder CoreOps Notification Router

## Policy

* Kandidat: Open Policy Agent

## Workflows

Zu bewerten:

* eigene persistente Job Engine
* Temporal
* andere geeignete Engines

## Topologie

Zunächst:

* PostgreSQL
* relationale Daten
* vorberechnete Graphansichten

Keine Graphdatenbank im MVP ohne nachgewiesenen Bedarf.

---

# 47. Repository-Zielstruktur

Die genaue Struktur wird in einem NDF-Work-Package finalisiert.

Vorgeschlagene Basis:

```text
coreops/
├── .claude/
│   └── skills/
├── .github/
│   └── workflows/
├── apps/
│   ├── web/
│   ├── api/
│   └── worker/
├── agents/
│   ├── coreops-agent/
│   └── site-relay/
├── packages/
│   ├── contracts/
│   ├── ui/
│   ├── policy/
│   ├── telemetry/
│   ├── events/
│   ├── topology/
│   └── testing/
├── domain-packs/
│   ├── generic-linux/
│   ├── proxmox/
│   ├── docker/
│   ├── network/
│   ├── print/
│   └── core-products/
├── observability/
│   ├── grafana/
│   ├── prometheus/
│   ├── loki/
│   └── otel/
├── docs/
│   ├── architecture/
│   ├── security/
│   ├── operations/
│   ├── integrations/
│   └── user-guide/
├── adr/
├── project-brain/
├── project-system/
├── scripts/
├── tests/
├── ROADMAP.md
├── README.md
├── SECURITY.md
├── CONTRIBUTING.md
└── LICENSE
```

---

# 48. NDF-Projektdateien

Der New Project Flow soll mindestens folgende Dateien erzeugen:

```text
project-system/project-manifest.yaml
project-system/PROJECT_PROFILE.md
project-brain/PROJECT_BRAIN.md
project-system/WORK_PACKAGE_QUEUE.md
ROADMAP.md
```

Zusätzlich für CoreOps empfohlen:

```text
project-system/HEALTH_SCORE.md
project-system/NEXT_PHASE.md
project-system/RISK_REGISTER.md
project-system/DECISION_INDEX.md
project-brain/CONTEXT_PACK_FOUNDATION_0_1.md
docs/architecture/COREOPS_CONCEPT_V3.md
docs/architecture/PLATFORM_PLANES.md
docs/security/SECURITY_BASELINE.md
docs/security/THREAT_MODEL.md
adr/README.md
```

---

# 49. Roadmap

## Foundation 0.1 – Platform Foundation

Nur Dokumentation, Architektur, Governance und Sicherheitsgrundlagen.

* NDF-Projektstruktur
* Vision
* Scope Lock
* Plane-Architektur
* Source of Truth
* Desired State
* Policy
* Trust
* OIC
* Events
* Deploymentmodell
* Offline-Baseline
* Topologiemodell
* Domain-Pack-Governance
* UX-Informationsarchitektur
* Teststrategie
* Risiko- und ADR-Register

Kein produktiver Anwendungscode.

## Release 0.1 – Observe

* Setup Wizard
* lokaler Owner
* Inventory
* Standorte
* Gruppen
* ICMP
* SNMPv3
* generisches HTTP Monitoring
* Grafana
* Audit
* Netzwerkstatus
* Druckerstatus
* Linux Discovery
* Proxmox Discovery
* Docker Discovery
* Core-Produktadapter
* vollständig read-only

## Release 0.2 – Map

* LLDP
* CDP
* Interfaces
* MAC
* ARP
* VLAN
* Topologie v1
* Confidence Score
* manuelle Ergänzungen
* historische Snapshots
* Path Explorer
* erste Impact Analysis
* Basis-IPAM

## Release 0.3 – Plan

* Desired State
* Drift
* Deployment Registry
* Environments
* Deployment Blueprints
* Policy Evaluation
* Change Safety Score
* Maintenance Windows
* Approval Inbox
* Preview

Noch keine breite produktive Automation.

## Release 0.4 – Deploy

* Docker Compose
* ausgewählte Linux-Dienste
* Core-Produkte
* gestaffelte Deployments
* Health Checks
* Smoke Tests
* Last-Known-Good
* einfache Rollbacks
* Evidence Packs
* CorePack-Deployment

## Release 0.5 – Automate

* persistente Workflows
* Runbooks
* Skriptbibliothek
* Wartungskampagnen
* Event-Korrelation
* Incident Response
* Linux Updates
* LXC
* Proxmox Operations

## Release 0.6 – Extend

* Windows
* Hyper-V
* Kubernetes
* Helm
* Netzwerkkonfiguration
* Druckerbereitstellung
* Firmware
* Zertifikatsworkflows
* erweiterte DDI-Funktionen

## Release 0.7 – Scale

* Multi-Site
* Site Relay
* HA
* externe Vaults
* Active Directory
* Multi-Tenant
* Compliance Packs
* BSI/CIS
* Partner- und Community-Packs

> Registrierungs-Note: Die hier genannten Release- und Meilensteinnummern sind Quellenangaben aus dem Concept. Die verbindliche Release-Taxonomie (inklusive der Kollision zwischen `Foundation 0.1` und `Release 0.1 – Observe`) wird durch `CO-WP-003` festgelegt.

---

# 50. Foundation Work-Package Queue

Die Queue ist ein initialer Vorschlag und muss durch Nova geprüft und gegebenenfalls neu geschnitten werden.

| WP        | Typ                 | Inhalt                                        |
| --------- | ------------------- | --------------------------------------------- |
| CO-WP-001 | `docs-only`         | NDF Project Bootstrap                         |
| CO-WP-002 | `docs-only`         | Project Brief, Vision und Scope Lock          |
| CO-WP-003 | `docs-only`         | Plane-Architektur und Modulgrenzen            |
| CO-WP-004 | `security-baseline` | Threat Model und Trust Boundaries             |
| CO-WP-005 | `docs-only`         | Source-of-Truth- und Provenance-Modell        |
| CO-WP-006 | `docs-only`         | Observed State, Desired State und Drift       |
| CO-WP-007 | `security-baseline` | Policy-, Approval- und Fail-Closed-Modell     |
| CO-WP-008 | `security-baseline` | Agent Enrollment und Machine Identity         |
| CO-WP-009 | `docs-only`         | OIC v0.1                                      |
| CO-WP-010 | `docs-only`         | Eventmodell und Korrelation                   |
| CO-WP-011 | `docs-only`         | Deployment Control Plane und Blueprint        |
| CO-WP-012 | `security-baseline` | Artifact Trust und Supply Chain               |
| CO-WP-013 | `docs-only`         | Offline-, Isolated- und Air-Gap-Baseline      |
| CO-WP-014 | `docs-only`         | Topology Graph und Evidence Model             |
| CO-WP-015 | `docs-only`         | Domain-Pack- und Integrations-Governance      |
| CO-WP-016 | `docs-only`         | UX, Navigation und Dashboard-System           |
| CO-WP-017 | `docs-only`         | Telemetrie- und Normalisierungsschema         |
| CO-WP-018 | `docs-only`         | Datenklassifizierung, Retention und Redaction |
| CO-WP-019 | `docs-only`         | Teststrategie und Integration Lab             |
| CO-WP-020 | `review-only`       | Foundation Readiness Review                   |
| CO-WP-021 | `release-prep`      | Foundation 0.1 Release Preparation            |

Die Foundation-Phase soll keine funktionalen Features implementieren.

> Registrierungs-Note: Diese Concept-interne Queue (21 Work Packages) weicht von der aktuell verbindlichen CoreOps-Queue in `project-system/WORK_PACKAGE_QUEUE.md` (CO-WP-001 … CO-WP-031) ab. Die verbindliche Queue wurde in CO-WP-001 definiert; diese Concept-Fassung ist als Quellenvorschlag zu verstehen. Die Abweichung ist ein Klärungspunkt (siehe Decision Classification).

---

# 51. Initiale ADR-Kandidaten

| ADR      | Thema                                            |
| -------- | ------------------------------------------------ |
| ADR-0001 | CoreOps als universelle Operations Control Plane |
| ADR-0002 | Modularer Monolith                               |
| ADR-0003 | Read-only First                                  |
| ADR-0004 | Offline First                                    |
| ADR-0005 | Plane-Architektur                                |
| ADR-0006 | CoreOps Integration Contract                     |
| ADR-0007 | Standards First                                  |
| ADR-0008 | Native Widgets vor allgemeinen iFrames           |
| ADR-0009 | Source of Truth und feldbasierte Provenance      |
| ADR-0010 | Desired State getrennt von Observed State        |
| ADR-0011 | Drift Detection vor Reconciliation               |
| ADR-0012 | zentrale fail-closed Policy Engine               |
| ADR-0013 | kurzlebige Machine Identities                    |
| ADR-0014 | ausgehende Agent-Verbindungen                    |
| ADR-0015 | Build und Deployment getrennt                    |
| ADR-0016 | unveränderliche Deployment-Artefakte             |
| ADR-0017 | Artifact Trust, SBOM und Provenance              |
| ADR-0018 | versioniertes Eventmodell                        |
| ADR-0019 | persistente und wiederaufnehmbare Workflows      |
| ADR-0020 | evidenzbasierte Topologie                        |
| ADR-0021 | manuelle Topologieautorität                      |
| ADR-0022 | Domain Packs umgehen die Control Plane nicht     |
| ADR-0023 | Configuration Vault                              |
| ADR-0024 | Data Classification und Redaction                |
| ADR-0025 | CoreOps Self-Protection                          |
| ADR-0026 | Integration Quality Levels                       |
| ADR-0027 | PostgreSQL vor Graphdatenbank im MVP             |
| ADR-0028 | Grafana als Analysewerkzeug                      |
| ADR-0029 | KI nur beratend                                  |
| ADR-0030 | CorePack und Offline Trust                       |

ADRs werden nicht allein aufgrund dieser Liste als akzeptiert betrachtet.

Jeder ADR benötigt:

* Kontext
* Problem
* Optionen
* Vor- und Nachteile
* Sicherheitsauswirkung
* Betriebswirkung
* Offline-Auswirkung
* Entscheidung
* Konsequenzen
* Status
* Human-Maintainer-Freigabe

> Registrierungs-Note: Diese 30 ADR-Kandidaten sind reine Kandidaten. Es wird durch dieses Dokument kein ADR akzeptiert. Es wurden keine ADR-Dateien erzeugt.

---

# 52. Sicherheitsinvarianten

Folgende Regeln gelten projektweit:

1. Keine allgemeine Remote-Root-Shell.
2. Keine ungeprüfte Ausführung externer Skripte.
3. Keine Secrets in Logs, Prompts oder Exports.
4. Keine privilegierte Aktion ohne Policy.
5. Keine destruktive Funktion ohne Blueprint.
6. Keine Ausführung ohne Preview.
7. Keine Bulk-Aktion ohne Zielvorschau und Begrenzung.
8. Kein Production-Deployment mit `latest`.
9. Kein Deployment ohne Health Check.
10. Kein stiller Fallback auf ein anderes Artefakt.
11. Kein Plugin darf die Control Plane umgehen.
12. Kein Agent erhält automatisch unbegrenzte Rechte.
13. Keine unbekannte Policy-Entscheidung wird als Freigabe interpretiert.
14. Kein Restore ohne Integritäts- und Kompatibilitätsprüfung.
15. Keine automatische Änderung manueller, gesperrter Topologiedaten.
16. Kein Netzwerkscan außerhalb freigegebener Bereiche.
17. Keine Cloudpflicht für Kernfunktionen.
18. KI darf keine autonome Infrastrukturänderung durchführen.
19. Jeder privilegierte Job benötigt Audit-Korrelation.
20. CoreOps muss seine eigene Betriebsabhängigkeit erkennen.

---

# 53. Qualitätsanforderungen

## Funktional

* nachvollziehbare Zustände
* klare Fehleranzeigen
* keine stillen Teilfehler
* idempotente Aktionen
* kontrollierte Wiederaufnahme
* nachvollziehbare Abhängigkeiten

## Security

* Least Privilege
* Defense in Depth
* Fail Closed
* mTLS
* Secret Redaction
* Signaturprüfung
* Audit
* Approval
* Rate Limits
* sichere Defaults

## UX

* Simple Mode
* Expert Mode
* klare Risikostufen
* klare Capability-Anzeige
* verständliche Previews
* keine irreführenden Erfolgsanzeigen
* mobile Kernfunktionen
* Accessibility Baseline

## Betrieb

* Docker-first
* lokale Installation
* Backup und Restore
* Offlinebetrieb
* Degraded Mode
* Support Bundle
* Health Checks
* integrierte Dokumentation

## Internationalisierung

* Deutsch und Englisch ab Foundation
* Deutsch darf Standardsprache sein
* Übersetzungen werden strukturiert gepflegt
* keine hartcodierten UI-Texte

---

# 54. Branding und Support

CoreOps soll ein professionelles Branding erhalten, das mit SpeakCore und CastCore mithalten kann.

Geplant:

* Logo
* Wortmarke
* Icon
* Farbwelt
* Typografie
* UI-Style-System
* Dark und Light Mode
* Diagrammstil
* Produkt-Screenshots
* Social Assets
* Repository Banner
* Favicon
* App Icons

Der projektübergreifende freiwillige Support-Hinweis soll professionell und unaufdringlich eingebunden werden:

* README
* About
* Dokumentation
* Website Footer

Keine Pop-ups, Paywalls oder aufdringlichen Hinweise.

---

# 55. Erste Aufgabe in diesem neuen Projektchat

> Registrierungs-Note: Dieser Abschnitt beschreibt die ursprüngliche Erstaufgabe an Nova aus dem Quelldokument. Er ist als historische Quellenaussage registriert; der tatsächliche Foundation-Start erfolgte bereits über CO-WP-001 und die aktuelle CoreOps-Work-Package-Queue.

Behandle dieses Dokument als **akzeptierte Produktvision und Ausgangsbasis**, aber noch nicht als automatisch akzeptierte technische Detailentscheidung.

Starte nicht mit Anwendungscode.

Führe zuerst ein NDF-konformes Foundation Intake durch.

## Die erste Antwort soll enthalten

1. eine kurze Bestätigung des verstandenen Zielbildes,
2. eine Prüfung, ob die vorgesehene Foundation-Queue NDF-konform geschnitten ist,
3. erkennbare Widersprüche oder fehlende Foundation-Entscheidungen,
4. eine Empfehlung für die genaue Foundation-0.1-Reihenfolge,
5. eine klare Definition von `CO-WP-001`,
6. den vollständigen Claude-Prompt für `CO-WP-001`,
7. die erwartete Rückmeldung an Nova,
8. die erwartete Compact Context Summary,
9. noch keinen Code und keinen Commit.

## CO-WP-001 soll voraussichtlich enthalten

**Titel:** NDF Project Bootstrap
**Typ:** `docs-only`
**Prompt Mode:** Full Prompt

Ziel:

* minimale NDF-Projektstruktur vorbereiten,
* Project Manifest anlegen,
* Project Profile anlegen,
* Project Brain anlegen,
* Roadmap als Foundation-Gerüst anlegen,
* Work-Package-Queue anlegen,
* aktuelle Phase dokumentieren,
* dieses Concept v3.0 als Referenz dokumentieren.

Dabei:

* keine Anwendung implementieren,
* keine Technologieentscheidung stillschweigend akzeptieren,
* keine ADRs als Accepted markieren,
* keine Git-Aktion ausführen,
* keine externen Dateien ungeprüft übernehmen,
* den Scope klein und überprüfbar halten.

Sollte sich zeigen, dass `CO-WP-001` dafür noch zu groß ist, ist es NDF-konform in kleinere Work Packages zu teilen und der Split zu begründen.

---

# 56. Gewünschter dauerhafter Projektworkflow

Nach jeder Claude-Rückmeldung:

1. Nova prüft die Rückmeldung.
2. Nova bewertet das Ergebnis.
3. Nova nennt offene Notes und Risiken.
4. Nova entscheidet GO, GO WITH NOTES, REWORK, SPLIT oder STOP.
5. Nova formuliert nur bei Freigabe den nächsten Schritt.
6. Der Human Maintainer entscheidet über Commit und Push.
7. Der Project Brain wird fortgeschrieben.
8. Der aktuelle Context Pack wird aktualisiert.
9. Relevante Erfahrungen werden als NDF-Learnings dokumentiert.

Bei produktiv relevanten Änderungen soll Nova zusätzlich Hinweise für ein Update und einen Test auf einer Proxmox-VM bereitstellen.

---

# 57. Erwartete Haltung von Nova

* ehrlich
* kritisch
* sicherheitsorientiert
* konstruktiv
* keine künstliche Zustimmung
* keine unnötige Komplexität
* keine versteckten Scope-Erweiterungen
* keine vorschnellen Architekturentscheidungen
* keine Behauptung ohne Evidenz
* keine Umsetzung außerhalb freigegebener Work Packages
* kleine Schritte
* nachvollziehbare Entscheidungen
* klare Empfehlungen

CoreOps soll professionell, langfristig wartbar und realistisch umsetzbar werden.

**CoreOps – One Dashboard. Controlled Operations.**
