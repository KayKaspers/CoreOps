# CoreOps – Initial Support Boundary (Observe Target)

> Status: **Accepted**
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004` (docs-only). Begleitdokument: [FOUNDATION_CAPABILITY_MATRIX.md](../architecture/FOUNDATION_CAPABILITY_MATRIX.md).

Dieses Dokument beschreibt den **geplanten** ersten funktionalen Zielumfang des Meilensteins **Observe**. Es ist keine Implementierung und kein Supportnachweis.

```text
CoreOps besitzt noch keine implementierte Runtime-Capability.
Keine Integration besitzt aktuell den Supportstatus Supported.
```

Für **alle** hier genannten Ziele gilt bis zu tatsächlicher Implementierung und Evidenz:

```text
Implementation Status: not-implemented
Support Status: not-supported
```

Der Observe-Meilenstein ist **vollständig read-only** gegenüber Zielsystemen und reicht höchstens bis **Integrationslevel 2 (Observed)**.

## 1. Core Platform – Observe Target

Zielkandidaten (alle `not-implemented` / `not-supported`):

- Setup Wizard
- lokaler Owner
- grundlegende Rollen
- Assets
- Standorte
- Gruppen
- read-only Inventory
- native Übersichtsansichten
- Audit
- integrierte Dokumentation
- DE/EN
- Grafana Workspace
- lokale self-hosted Installation
- Betrieb ohne aktive Cloudpflicht

## 2. Generic Protocol Targets

### ICMP
Geplanter Umfang: Erreichbarkeit · Latenz · Paketverlust · Zeitstempel · kontrollierte Zielbereiche · Rate Limits.
**Keine Discovery außerhalb freigegebener Netze.**

### SNMPv3
Geplanter Umfang: Discovery · Basisinventar · Gerätestatus · ausgewählte standardisierte Metriken · Printer MIB · Interface-Basisdaten.
**SNMPv1 und SNMPv2c sind nicht Bestandteil der sicheren Standardkonfiguration.** Eine spätere optionale Legacy-Unterstützung benötigt eine gesonderte Sicherheitsentscheidung (nicht in diesem WP getroffen).

### HTTP und HTTPS
Geplanter Umfang: Erreichbarkeit · Statuscode · Antwortzeit · optionaler erwarteter Inhalt · TLS-Zertifikatsmetadaten.
**Keine Speicherung von Credentials oder sensitiven Response-Bodies.**

## 3. Initiale Read-only Domain Targets

### Generic Linux
Basisinventar · OS und Version · Hostname · Ressourcenübersicht · begrenzte Dienstgesundheit. Updatescan später separat.
**Transport sowie Agent-/Agentless-Entscheidung bleiben offen (keine Auswahl in diesem WP).**

### Proxmox
Cluster-, Node-, VM- und LXC-Inventar · Status · Versionen · ausgewählte Health-Metriken.
**Keine Start-/Stop-/Migrations-/Snapshot-Aktion.**

### Docker
Hosts · Container · Images · Compose-Projekte (sofern sicher ermittelbar) · Status · Versionen.
**Keine Shell · kein Exec · kein Restart · kein Deployment.** Der sichere Zugriffspfad auf Docker bleibt technisch offen.

### Generic Print
Hersteller · Modell · Seriennummer (sofern verfügbar) · Status · Supplies · Counters · Warnungen.
**Keine Auftragsnamen als Standard-Telemetrie · keine Queue-Manipulation · keine Treiberbereitstellung.**

### Generic Network Status
Erreichbarkeit · Hersteller und Modell (sofern standardisiert ermittelbar) · Firmwareversion · Uptime · CPU/RAM (sofern standardisiert) · Interface-Status.
**Keine Konfigurationsänderung · keine Firmware-Aktion · keine Topologiebehauptung ohne Evidenz.**

### Core Products (First-Party-Adapter-Zielklasse)
Produktidentität · Version · Health · Instanzstatus · Capabilities.
**Keine privilegierte Produktaktion.** Welche Core-Produkte zuerst integriert werden, bleibt einem späteren Work Package vorbehalten.

## 4. Explizit nicht im Observe-Zielumfang

Ausgeschlossen: schreibende Netzwerkaktionen · Drucker-Queue-Management · Druckertreiberbereitstellung · Deployments · Updates und Patch-Ausführung · Skriptausführung · Workflow Runtime · agent-seitige privilegierte Aktionen · Firmware · Konfigurations-Restore · automatische Driftkorrektur · Enforce Mode · Hyper-V-Management · VMware-Management · Kubernetes-Management · Helm · Active Directory · Multi-Tenant · Multi-Site Relay-Ausführung · HA · Compliance Automation · automatische Incident Remediation.

## 5. Abgrenzung zu Map und späteren Meilensteinen

**Topologie, LLDP/CDP, MAC/ARP, VLAN, Path Explorer und Basis-IPAM gehören zum Meilenstein `Map`, nicht zu Observe.** `Plan` enthält Desired State und Preview; `Deploy` enthält die ersten kontrollierten Write-Aktionen. Dieses Work Package zieht keine spätere Funktion in Observe vor. Ordnet das Concept eine Funktion einem späteren Meilenstein zu, bleibt diese Zuordnung bestehen.

## 6. Hersteller- und Supportgrenze (Auflösung CCR-12)

> **Proposed Foundation Governance.**

```text
Die im CoreOps Concept genannten Hersteller sind Priorisierungs- und
Integrationskandidaten.

Eine Nennung stellt weder Support, Partnerschaft, Zertifizierung,
Kompatibilitätsgarantie noch getestete Integration dar.
```

Hersteller-/Produktnamen sind ausschließlich zulässig als `candidate`, `priority candidate` oder `planned integration family`. Ohne Evidenz **nicht** zulässig: `supported`, `certified`, `verified compatible`, `official partner`, `fully compatible`.

Ein späteres Supportversprechen benötigt mindestens: (1) Hersteller und Produktfamilie, (2) getestetes Modell, (3) getestete Firmware-/Softwareversion, (4) getestete CoreOps-Version, (5) unterstützte Capabilities, (6) nicht unterstützte Capabilities, (7) benötigte Berechtigungen, (8) Authentifizierungsverfahren, (9) Test-Fixtures, (10) positive Tests, (11) Fehler- und Timeouttests, (12) Least-Privilege-Prüfung, (13) Offline-Verhalten, (14) bekannte Grenzen, (15) letztes Testdatum, (16) verantwortlicher Herausgeber, (17) dokumentierter Supportstatus.

Ohne diese Evidenz bleibt die Integration `not-supported` oder — nach vorhandener Implementierung — höchstens `experimental`.

## 7. Herausgeberklassen

```text
Built-in
First-Party Pack
Verified Partner Pack
Community Pack
Local Pack
```

Regeln:
- Herausgeberklasse und Supportstatus sind **unabhängig**.
- Ein First-Party Pack ist **nicht** automatisch `supported`.
- Ein Community Pack ist **nicht** automatisch unsicher.
- Ein Verified Partner Pack benötigt ein später definiertes Verifikationsverfahren.
- `Local Pack` bedeutet lokal verwaltet, **nicht** automatisch vertrauenswürdig.
- Kein aktuelles CoreOps-Pack besitzt bereits einen produktiven Supportstatus.

## 8. Standardisierter Support-Evidence-Satz

Ein Supportstatus über `not-supported` hinaus benötigt einen vollständigen Evidence-Satz:

```text
Integration identity
Integration version
CoreOps version
Publisher class
Target model
Target firmware/software version
Capability list
Required permissions
Authentication method
Fixture references
Test references
Failure-mode tests
Timeout tests
Rate-limit behavior
Offline behavior
Security review status
Known limitations
Last validation date
Validator
Support status
Revocation status
```

**Validator:** ausschließlich neutrale Rollen oder Projektidentitäten (z. B. `Human Maintainer`, `Nova`, `CoreOps Project`) — keine private Person.

## 9. Sicherheitsgrenzen der Observe-Grenze

- SNMPv3 ist der sichere Standard; Legacy-SNMP wird nicht stillschweigend aktiviert.
- Keine Remote-Root-Shell und kein Docker-`exec` werden als Capability zugesagt.
- Keine Write-Aktion auf Zielsystemen im Observe-Meilenstein.
- Drucker-Auftragsnamen und Benutzerdaten sind personenbezogen zu behandeln und **nicht** Teil der Standard-Observe-Telemetrie.
- Discovery nur in freigegebenen Netzbereichen mit Rate Limits.
