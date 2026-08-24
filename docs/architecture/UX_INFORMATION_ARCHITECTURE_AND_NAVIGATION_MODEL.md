# CoreOps – UX Information Architecture and Navigation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation UX information architecture and navigation model
> Implementation Status: Not implemented
> Technology Selection: None
> Frontend/Framework/Component/Routing/State/Design-Tool Technology: Not selected
> Validation Status: Not performed
> Usability Validation: Not performed
> Certification Status: None claimed
> CDS Adoption: Not started
> CDS Pilot: Inactive / not activated by this document
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-027 (docs-only / UX information architecture and dashboard system foundation)

## 1. Status

Technologieunabhängige **Information Architecture (IA)** und **Navigationssemantik** für die CoreOps Experience-Ebene: Application Shell, Presentation Context, Informationsbereiche, Drill-down-/Rückkehrsemantik, Verfügbarkeits-/Berechtigungssichtbarkeit, Presentation Modes und degradierte Sichtbarkeit. Companions: [DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md](DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md) und [UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md). Erweitert **kein** bestehendes Autoritätsmodell; erzeugt **keine** neue Experience-Autorität und **kein** paralleles Datenmodell. Keine Frontend-, Framework-, Komponenten-, Routing-, State-Management- oder Design-Technologie ausgewählt.

## 2. Purpose

Eine Oberfläche ist kein System of Record. Das Modell legt fest, warum `presentation ≠ source of truth`, `navigation ≠ domain ownership`, `navigation ≠ target binding`, `selection ≠ authorization`, `visible ≠ authorized`, `hidden ≠ nonexistent`, `unknown ≠ healthy`, `stale ≠ current`, `partial ≠ complete` und `degraded ≠ failed` — und wie ein Nutzer sich orientiert, ohne dass Orientierung als Autorität missverstanden wird.

## 3. Scope

Experience-Autoritätsgrenze · Presentation Context und Scope-Bindung · Top-Level-Informationsbereiche mit Traceability · Application Shell/Orientierung · Navigationssemantik · Overview→List→Detail→Evidence-Fluss · State-Navigation · verfügbarkeits-/berechtigungssensitive Darstellung · Presentation Modes (Simple/Expert) · degradierte/eingeschränkte/offline Sichtbarkeit · Leer-/Unbekannt-/Verweigert-Zustände · Such-/Filter-/Hilfe-Einstiegspunkte · CDS Reconciliation Boundary. Dashboard-Detailsemantik im [Companion Dashboard Model](DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md); Aktions-, Accessibility- und Disclosure-Anforderungen in der [Companion UX Policy](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md).

## 4. Non-Goals

- Keine Frontend-Implementierung; kein HTML/CSS/JS; keine React-/Vue-/Angular-/Svelte-Auswahl; kein Komponenten-Framework; keine Routing-Bibliothek; kein State-Management; kein Design-Tool; kein Logo; keine Farbpalette; keine Typografie; keine Icon-Bibliothek; kein Animations-/Responsive-/CSS-Architektur-System.
- Keine CDS-Token-Übernahme, kein CDS-Paket, kein CDS Product Profile, keine CDS-Pilot-Aktivierung.
- Keine Mockups, keine Screenshots, keine API-/Datenbank-/Runtime-/Deployment-/MCP-/Event-Bus-Auswahl.
- Keine Behauptung implementierter, getesteter, validierter oder usability-geprüfter Oberflächen.

## 5. Concepts

`experience layer` · `application shell` · `orientation` · `information area` · `presentation context` · `context scope` · `context binding` · `view` · `overview` · `list` · `detail` · `evidence view` · `domain lens` · `drill-down` · `return semantics` · `presentation mode` · `availability state` · `permission state` · `capability state` · `empty state` · `unknown state` · `degraded visibility` · `entry point` · `explanation entry point`.

```text
information area = benannter Orientierungsbereich der Experience-Ebene,
                   der auf eine bereits bestehende CoreOps-Capability,
                   ein Modul oder eine autoritative Domaene zurueckfuehrt
information area ≠ Datenowner
information area ≠ Autoritaetsgrenze
```

## 6. Experience Authority Boundary

Die Experience-Ebene bleibt **MOD-EXP-001 (`foundation-core`)** gemäß [COREOPS_LOGICAL_MODULE_ARCHITECTURE.md §9](COREOPS_LOGICAL_MODULE_ARCHITECTURE.md) und [DEC-S-68](../../project-system/DECISION_INDEX.md). Sie **stellt dar**, **validiert Eingaben**, **nimmt Nutzerintents entgegen** und **routet Anfragen** an policy-/control-nahe Module. Sie besitzt **keinen** autoritativen Betriebszustand, **keine** Approval-, **keine** Execution-Autorität, führt **keine** privilegierten Aktionen direkt aus und schreibt **nicht** direkt über Adapter.

```text
UI intent            ≠ approval
approval             ≠ execution authorization
action visible       ≠ action authorized
action authorized    ≠ action executed
executed             ≠ successful
successful           ≠ verified
evidence visible     ≠ authority
observation          ≠ desired state
dashboard summary    ≠ authoritative state
presentation         ≠ source of truth
```

Dieses Dokument schafft **keine** neue Experience-Autorität. Jede Autoritätsfrage bleibt bei den bestehenden Modellen ([Policy](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Approval](../security/APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md), [Execution Authorization](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [Source of Truth](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md)).

## 7. Presentation Context and Scope Binding

Jede Ansicht trägt einen **Presentation Context**. Dessen Scope-Typen sind **nicht neu**: die Experience-Ebene **referenziert die bereits registrierten Scope-Identitäten** aus [WORKSPACE_RBAC_AND_SCOPE_MODEL.md §10](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md). Es entsteht **kein** zweites Scope-Modell; die genannte RBAC-Datei bleibt unverändert und autoritativ.

| Context Scope | Bedeutung in der Darstellung |
| ------------- | ---------------------------- |
| `global platform scope` | plattformweite Orientierung; keine implizite Zielbindung |
| `workspace scope` | mandanten-/arbeitsbereichsgebundene Sicht |
| `environment scope` | umgebungsgebundene Sicht innerhalb eines Workspace |
| `resource-group scope` | gruppierte Ressourcen; Gruppe ≠ Ausführungsziel |
| `individual resource scope` | einzelne registrierte Ressource |
| `workflow scope` | Workflow-/Orchestrierungssicht |
| `job scope` | einzelne Ausführungsinstanz |
| `evidence scope` | Evidence-/Auditsicht mit eigener Disclosure-Grenze |

Verbindliche Regeln:

```text
presentation context ist explizit, sichtbar und programmatisch bestimmbar
global context ≠ workspace context ≠ resource context
context switch ist ein expliziter Akt, nie ein Nebeneffekt
kein stilles Ausweiten des Scopes
kein stilles Uebertragen eines Scopes zwischen Workspaces
unknown context → fail-closed auf den engeren Scope
```

`A broader role name does not automatically imply a broader scope` ([RBAC §10](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md)) gilt in der Darstellung unverändert: eine breitere Ansicht impliziert keine breitere Autorität.

### 7.1 Presentation Context versus Authorization Scope

Ein Presentation Context ist eine **Sichtbindung**, keine Autorisierungsaussage. Die Experience-Ebene darf eine Ansicht an eine **bestehende Scope-Identität binden**; die Autorisierungsentscheidung selbst bleibt vollständig bei **MOD-IAM-001 / MOD-POL-001** ([WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md), [POLICY_DECISION_AND_EVALUATION_MODEL.md](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md)).

```text
presentation context           ≠ authorization scope
referencing a scope identity   ≠ performing authorization
context narrowing              ≠ permission mutation
context widening               ≠ permission grant
navigation                     ≠ authorization
UI visibility                  ≠ permission result
view bound to a scope          ≠ access granted within that scope
```

Verbindlich:

- Die Experience-Ebene **referenziert** Scope-Identitäten; sie **definiert**, **erzeugt**, **erweitert** und **interpretiert** sie nicht.
- Ein Wechsel des Presentation Context verengt oder erweitert **ausschließlich die Sicht**. Er verändert **keine** Berechtigung, **keine** Rollenzuweisung, **keinen** Scope-Grant und **keine** Policy-Bewertung.
- Eine Verengung des Contexts ist **keine** zusätzliche Sicherheitskontrolle: sie ersetzt keine Autorisierungsprüfung und darf nicht als solche dargestellt oder verlassen werden.
- Was in einer Ansicht sichtbar ist, ist **Ergebnis** der autoritativen Autorisierungsentscheidung, **nicht deren Ursache**. Eine Oberfläche leitet aus Sichtbarkeit niemals eine Berechtigung ab.
- Das bestehende RBAC-/Policy-Modell bleibt die **einzige** Autorität für Rollen, Permissions, Scope-Zuweisung und Autorisierungsergebnis. Dieses Dokument ändert es nicht und dupliziert es nicht.

## 8. Top-Level Information Areas

Jeder Bereich führt auf eine **bereits bestehende** Capability-Domäne ([FOUNDATION_CAPABILITY_MATRIX.md](FOUNDATION_CAPABILITY_MATRIX.md)) und ein bestehendes Modul zurück. Kein Bereich wurde erfunden, weil er nützlich klingt. **Ein Navigationseintrag begründet keine Datenownership.**

| Information Area | Zweck | Traceability (Capability) | Traceability (Modul) |
| ---------------- | ----- | ------------------------- | -------------------- |
| Orientation and Overview | scope-gebundener Einstieg, Übersichten | `CAP-PLATFORM-003`, `CAP-PLATFORM-007` | MOD-EXP-001 |
| Inventory and Resources | registrierte Ressourcen, Konfiguration/State | `CAP-INVENTORY-*` | MOD-INV-001, MOD-STA-001 |
| Discovery and Observation | beobachteter Zustand, Monitoring, Telemetrie | `CAP-DISCOVERY-*`, `CAP-MONITORING-*` | MOD-OBS-001 |
| Topology | Beziehungs-/Pfadsichten (`future-extension`) | `CAP-TOPOLOGY-*` | MOD-TOP-001 |
| State and Drift | desired/observed/effective/drift | `CAP-INVENTORY-*`, `CAP-MONITORING-*` | MOD-STA-001 |
| Operations and Automation | Workflows, Jobs, Remediation-Intents | `CAP-AUTOMATION-*` | MOD-WFL-001, MOD-EXE-001 |
| Deployment and Change | Deployments, Updates, Blueprints | `CAP-DEPLOY-*` | MOD-DEP-001 |
| Policy, Approval and Access | Policy, Approval Inbox, Rollen/Scopes | `CAP-IDENTITY-*` | MOD-POL-001, MOD-IAM-001 |
| Trust, Artifacts and Offline | Artifact Trust, CorePack, Offline-Intake | `CAP-TRUST-*` | MOD-OFF-001 |
| Protection and Recovery | Operational Mode, Degraded, Recovery | `CAP-PROTECT-*` | MOD-EVD-001, MOD-STA-001 |
| Evidence and Audit | Audit Explorer, Evidence-Referenzen | `CAP-PLATFORM-006` | MOD-EVD-001 |
| Integrations and Domain Packs | Adapter, Domain Packs, Erweiterungen | `CAP-*` Integrationsanteile | MOD-ADP-001, MOD-EXT-001 |
| Platform Administration | Setup, Sprache, Dokumentation, Plattformkonfiguration | `CAP-PLATFORM-001`, `CAP-PLATFORM-002`, `CAP-PLATFORM-004` | MOD-EXP-001 |

**Notification/Communication** (MOD-NOT-001, `optional-platform`) erscheint als Konfigurations- und Verlaufsbereich innerhalb *Platform Administration* bzw. *Operations*; ein Benachrichtigungskanal ist **kein** Steuerkanal.

### 8.1 Domain Lenses (keine eigenen Autoritätsbereiche)

Fachdomänen wie Netzwerk (`CAP-NETWORK-*`), Druck (`CAP-PRINT-*`) und Virtualisierung/Container (`CAP-VIRT-*`) werden als **Domain Lenses** über den obigen Bereichen dargestellt — als gebundene Filter-/Sichtprofile, **nicht** als parallele Top-Level-Autorität. Ein Domain Pack darf Lenses beisteuern; es erhält dadurch **keine** Runtime- oder Datenautorität ([DOMAIN_PACK_GOVERNANCE_MODEL.md](DOMAIN_PACK_GOVERNANCE_MODEL.md)).

```text
domain lens        ≠ eigenes System of Record
domain lens        ≠ zusaetzliche Berechtigung
pack-provided view ≠ pack-owned data
```

## 9. Application Shell and Orientation

Die Shell trägt technologieunabhängig: aktuellen **Presentation Context** · aktiven **Information Area** · aktiven **Presentation Mode** · aktuellen **Operational Mode** der Plattform (falls nicht `normal`) · **Freshness-Indikator** der dargestellten Daten · **Identitäts-/Rollenkontext** · **Sprachkontext (DE/EN)** · Zugang zu Suche, Hilfe und Evidenz. Jedes dieser Elemente ist **beschreibend**; keines ist ein Steuerelement mit eigener Autorität.

```text
shell = Orientierung + Kontextanzeige
shell ≠ Autoritaetstraeger
platform mode indicator ≠ platform health verdict
```

## 10. Navigation Semantics

```text
navigation          ≠ target binding
selection           ≠ authorization
filter              ≠ execution scope
breadcrumb          ≠ authority chain
deep link           ≠ granted access
route available     ≠ capability available
route hidden        ≠ capability nonexistent
```

Regeln:

- **Kontexterhalt beim Drill-down:** Der Presentation Context wird beim Wechsel Overview→List→Detail→Evidence **mitgeführt und angezeigt**; er verengt sich höchstens, er weitet sich nie still.
- **Rückkehrsemantik:** Rückkehr zu einer Übersicht stellt den zuvor gültigen Kontext wieder her, ohne ihn still zu erweitern; ein zwischenzeitlich veralteter Stand wird als `stale` markiert, nicht still ersetzt.
- **Zielbindung:** Ein Ausführungsziel entsteht **nicht** durch Navigation, Auswahl oder Filter. Es muss durch die bestehenden operativen Modelle materialisiert und autorisiert werden ([Execution Authorization](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [Deployment Control Plane](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)).
- **Kein stiller Autoritätswechsel:** Ein Wechsel von Workspace, Environment oder Identität ist explizit, sichtbar und auditierbar.

## 11. Overview, List, Detail and Evidence Flow

Vier Verantwortungsstufen mit getrennter Aufgabe:

| Stufe | Verantwortung | Ausdrücklich nicht |
| ----- | ------------- | ------------------ |
| Overview | Priorisierung, Orientierung, Scope-Anzeige | Vollständigkeit, autoritative Aussage |
| List | Abgrenzung, Vergleich, Filterung, Abdeckungsanzeige | Beweis der Vollständigkeit der Grundgesamtheit |
| Detail | Dimensionen, Provenance, Verlauf, Grenzen | Autoritätsgewährung |
| Evidence | Referenzen, Verfügbarkeit, Validierungsstand, Grenzen | Rohevidenz-Freigabe per se |

```text
overview → list → detail → evidence
summary            ≠ completeness
list count         ≠ population count
detail             ≠ authority
presentation       ≠ source-of-truth ownership
evidence reference ≠ raw evidence discloseable
```

Jede Stufe muss **erklärbar** auf ihre Quelle zurückführen ([FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md), [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md)).

Die auf der Detailstufe gezeigten **Dimensionen** sind die Darstellungsdimensionen aus [Dashboard Model §11/§11.1](DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md). Sie sind eine **Darstellungsordnung über bestehenden autoritativen Modellen** und **kein** neues Statusschema; jede Dimension bleibt durch ihr eigenes autoritatives Modell governt. **Kein Subjekt muss alle Dimensionen materialisieren**, und eine für ein Subjekt nicht anwendbare Dimension ist kein unbekannter Wert:

```text
UX presentation dimensions ≠ new authoritative status schema
not applicable             ≠ unknown
dimension not applicable to a subject
   ≠ applicable dimension whose value is unknown
```

## 12. State Navigation

Die Navigation zwischen Zustandsarten folgt [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) und [DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md](DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md); es entsteht **kein** paralleles Zustandsmodell.

```text
desired state ≠ observed state ≠ effective state ≠ reported state ≠ last-known state
no observation  ≠ no drift
drift detected  ≠ drift remediated
effective state ≠ compliance
```

Die Darstellung hält getrennt navigierbar: gewünschter Zustand (inkl. Lifecycle `proposed`/`approved`/`active`/`superseded`/`suspended`/`withdrawn`/`expired`/`conflicted`) · beobachteter Zustand · effektiver Zustand · Drift · Operationen/Jobs · Deployment · Evidence/Audit. Ein Wechsel zwischen diesen Sichten verändert **keinen** Zustand.

## 13. Availability, Permission and Capability-sensitive Presentation

Vier **getrennte** Ursachenklassen dürfen nicht zu „nicht vorhanden" verschmelzen:

| Klasse | Bedeutung | Nutzerseitige Konsequenz |
| ------ | --------- | ------------------------ |
| `capability unavailable` | Fähigkeit im aktuellen Profil/Deployment nicht vorhanden | keine Berechtigungsfrage |
| `permission denied` | Fähigkeit vorhanden, Autorisierung verweigert | Berechtigungs-/Approval-Pfad |
| `state unknown` | Zustand nicht bestimmbar | fail-closed, keine Gesundheitsaussage |
| `restricted by operational mode` | Fähigkeit durch Mode suspendiert/verboten | Mode-/Recovery-Pfad |

```text
capability unavailable   ≠ permission denied
permission denied        ≠ capability unavailable
unknown                  ≠ healthy
restricted by mode       ≠ provider failure
provider/process healthy ≠ consumer capability usable
```

Die Unterscheidung ist **erklärungspflichtig**, aber **disclosure-gebunden**: die Begründungstiefe unterliegt [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md) und [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](../security/REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md). Wo die Existenz eines Objekts selbst schutzbedürftig ist, wird eine **einheitliche, nicht unterscheidende** Antwort gegeben und dies als bewusste Grenze dokumentiert — **nicht** als „nicht vorhanden" behauptet. Diese Spannung ist in §21/§25 offen benannt.

## 14. Presentation Modes (Simple / Expert)

Simple Mode und Expert Mode sind **bereits** im autoritativen Produktmodell benannt ([COREOPS_CONCEPT_V3.md, Abschnitt UX](COREOPS_CONCEPT_V3.md)); dort stehen sie als Produktanforderung, **ohne** definierte Semantik. Dieses Dokument definiert die **UX-Semantik** und erzeugt **keine** neue Autorität.

```text
Simple mode       ≠ different authorization model
Simple mode       ≠ fewer permissions
Expert mode       ≠ authority expansion
Expert mode       ≠ more permissions
Expert mode       ≠ Break Glass
hidden in Simple  ≠ nonexistent
expert visibility ≠ permission
mode preference   ≠ authorization
mode preference   ≠ Berechtigungsquelle
```

Verbindlich:

- Der Modus verändert **Informationsdichte und Darstellung fortgeschrittener Operationen** — **nie** das autoritative Berechtigungsergebnis.
- Der **aktive Modus ist sichtbar und programmatisch bestimmbar**.
- In Simple ausgeblendete Inhalte existieren weiterhin; ihre Existenz wird nicht geleugnet, und ein expliziter Weg zur vollständigen Sicht bleibt bestehen (soweit berechtigt).
- Ist die Modus-Präferenz **nicht verfügbar oder nicht ermittelbar**, gilt ein **safe presentation fallback**: die **Simple-/reduced-complexity-Darstellung** mit sichtbarem Hinweis; kein stilles Umschalten. Dieser Fallback ist **ausschließlich** eine Darstellungsentscheidung — er ist **kein** anderer Autorisierungszustand, **keine** Rechteeinschränkung und **keine** Sicherheitsmaßnahme:

```text
safe presentation fallback = Simple / reduced-complexity presentation
safe presentation fallback ≠ reduced permissions
safe presentation fallback ≠ different authorization state
safe presentation fallback ≠ degraded/restricted operational mode (§15)
```

- **Abgrenzung zu §15:** Presentation Modes (Simple/Expert) und Operational Modes (`normal`…`unknown`) sind getrennte Konzepte. Nur Operational Modes können Fähigkeiten tatsächlich einschränken; ein Presentation Mode kann das nie.
- Ein Moduswechsel ist **keine** privilegierte Operation, erzeugt **keine** Zielbindung und verändert **keine** Policy-Bewertung.

## 15. Degraded, Restricted and Offline Visibility

Die zehn Operational Modes aus [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md §6](DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md) (`normal` · `guarded` · `restricted` · `read-only` · `degraded` · `containment` · `recovery-only` · `recovery` · `emergency-stop` · `unknown operational state`) bleiben in der Darstellung **benannt und unterscheidbar**; die IA erfindet keine eigenen Modes.

```text
degraded     ≠ failed
degraded     ≠ normal operation with warnings
read-only UI ≠ read-only platform
restricted   ≠ read-only
offline      ≠ broken
last-known   ≠ current
```

Anforderungen: der aktive Mode ist in der Shell sichtbar · durch den Mode suspendierte Fähigkeiten werden als *suspendiert* (nicht als *fehlend*) dargestellt · Offline-/eingeschränkte Sichten kennzeichnen `last-known state` mit Erhebungszeitpunkt · Konnektivitätsklassen ([RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md](RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md)) werden nicht mit Fehlerzuständen vermischt.

## 16. Empty, Unavailable, Denied and Unknown States

Sechs **getrennt** darzustellende Nullzustände: `empty` (Scope hat keine Objekte) · `filtered-empty` (Filter ergibt keine Treffer) · `not-collected` (nie erhoben) · `unavailable` (Quelle/Fähigkeit nicht verfügbar) · `denied` (Autorisierung verweigert) · `unknown` (nicht bestimmbar).

```text
empty          ≠ not-collected ≠ unknown
filtered-empty ≠ empty scope
unavailable    ≠ denied
unknown        ≠ healthy ≠ zero
zero value     ≠ absence of data
```

Kein Nullzustand darf als `0` in einer Kennzahl erscheinen, ohne die Klasse mitzuführen ([TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md): `missing telemetry ≠ zero`).

## 17. Search, Filter and Selection Entry Points

Konzeptionell; keine Suchtechnologie ausgewählt (`CAP-PLATFORM-005 Global Search` bleibt `not-implemented`):

```text
search result ≠ authorization to view detail
search scope  ≠ execution scope
filter        ≠ policy
selection     ≠ target binding
saved view    ≠ standing authorization
```

Suche und Filter arbeiten **innerhalb** des aktiven Presentation Context und der geltenden Berechtigungen; Treffer außerhalb der Berechtigung werden nicht als Existenzbeweis geleakt (§13). Eine gespeicherte Ansicht speichert **Darstellung**, nie Autorität.

## 18. Help and Explanation Entry Points

Erklärung ist ein **erstklassiger** IA-Bestandteil, keine Fußnote (`CAP-PLATFORM-004 Integrated Documentation`, `not-implemented`). Erforderliche Einstiegspunkte: Bedeutung eines Statuswerts · Herkunft eines Feldes (Provenance) · Grund einer Einschränkung (§13) · Bedeutung des aktiven Operational Mode · Bedeutung des aktiven Presentation Mode · Voraussetzungen einer risikobehafteten Operation ([UX Policy §7–§9](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md)).

```text
explanation ≠ authorization
explanation ≠ evidence disclosure
help text   ≠ normative source
```

## 19. Bilingual and Accessibility Anchors

DE/EN ist Produktanforderung ([DEC-P-03](../../project-system/DECISION_INDEX.md), `CAP-PLATFORM-002`); Englisch bleibt kanonisch für maschinennahe Identifier ([DEC-S-38](../../project-system/DECISION_INDEX.md)). Die IA muss Textexpansion, sprachunabhängige Struktur und lokalisierungsfähige Bereichsnamen tragen; **Bereichs- und Kontextidentität ist sprachunabhängig stabil** (Anzeigename ≠ Identität, konsistent mit [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md)). Die vollständigen Accessibility-Anforderungen stehen in der [Companion UX Policy §16–§18](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md). **Keine** Konformitäts-, Validierungs- oder Zertifizierungsbehauptung.

## 20. CDS Reconciliation Boundary

Externer Stand zum Zeitpunkt der Prompt-Erstellung, **read-only** und **nicht** normativ übernommen: Core Design System HEAD `1fc53ae`; Semantic-Status-Familie `Candidate/Approved` (Source Revision `semantic-status-rev-0002-candidate`); zugelassene Evidenz `AE1-CDS-WP016-SEMSTATUS-004` (AE-1, nur Source-/Contract-Scope); `Stable: none`; Consumer-Evidenz: keine; Product-Evidenz: keine; AE-2/AE-3/AE-4: keine; Publikation: Private Development; CoreOps CDS Pilot **INACTIVE / NOT AUTHORIZED FOR EXECUTION**; `CDS-WP-017` inaktiv/nicht autorisiert/nicht definiert.

```text
CDS Candidate       ≠ Stable dependency
CDS Candidate       ≠ CoreOps adoption
CDS source evidence ≠ CoreOps consumer evidence
CDS pilot contract  ≠ active pilot
CDS semantics       ≠ CoreOps domain authority
```

Vergleich der CDS-Achsen (`condition` · `severity` · `confidence` · `freshness` · `evidence`) mit dieser IA:

| Befund | Klassifikation |
| ------ | -------------- |
| Achsenunabhängigkeit; kein normativer Aggregat-Score; `unknown` explizit; `stale ≠ current`; `unverified ≠ verified`; Evidenzgrenzen sichtbar; Farbe/Icon/Position allein tragen keine Bedeutung | **compatible overlap** |
| `presentation context`/Scope-Bindung mit acht CoreOps-Scopes und `navigation ≠ target binding` | **CoreOps-local** |
| Trennung `capability unavailable` / `permission denied` / `restricted by operational mode` / `unknown` | **CoreOps-local** |
| Operational Modes und Degraded-Sichtbarkeit als Darstellungsdimension | **CoreOps-local** |
| Mögliche spätere Abbildung der fünf CDS-Achsen auf die CoreOps-Statusdimensionen ([Dashboard Model §11](DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md)) | **future mapping candidate** |
| Direkte 1:1-Abbildung der CoreOps-Dimensionen auf nur fünf CDS-Achsen | **would lose meaning — not adopted** |
| Jede substanzielle Designübernahme | **requires future CDS pilot evidence** |

Diese Zuordnung ist **projektintern und nicht projektübergreifend normativ**. CoreOps Operational Truth bleibt CoreOps-eigen. Es wurden **keine** CDS-Dateien importiert, **keine** Tokens kopiert, **kein** Pilot aktiviert und **keine** CDS-Adoption-Evidenz erzeugt.

### 20.1 Pilot-Szenario-Themen als nicht-normativer Vollständigkeitsabgleich

Die externen Pilot-Szenarienthemen (`Application Foundation` · `Operations Overview` · `Inventory and Dense Data` · `State and Safety Patterns` · `Help / Accessibility / Localization`) wurden **ausschließlich** als Abdeckungsprüfung verwendet: Application Foundation → §9/§10; Operations Overview → §8/§11 und [Dashboard Model](DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md); Inventory and Dense Data → §8/§11/§17; State and Safety Patterns → §12/§13/§15 und [UX Policy](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md); Help/Accessibility/Localization → §18/§19. Es wurden **keine** Rollen, Layouts oder Komponenten daraus übernommen und **keine** CDS-Adoption-Evidenz erzeugt.

## 21. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Die Experience-Ebene stellt dar, validiert und routet; sie besitzt keinen autoritativen Zustand und keine Ausführungsautorität.
2. Ein Navigationsbereich begründet keine Datenownership und keine Autorität.
3. Der Presentation Context ist explizit, sichtbar und programmatisch bestimmbar; unbekannter Kontext ist fail-closed auf den engeren Scope.
3a. `presentation context ≠ authorization scope`; `referencing a scope identity ≠ performing authorization`; `context narrowing ≠ permission mutation`; `navigation ≠ authorization`; `UI visibility ≠ permission result`. Die Experience-Ebene bindet eine Sicht an eine bestehende Scope-Identität; MOD-IAM-001/MOD-POL-001 bleiben allein autoritativ für Autorisierung. Kein zweites Scope-Modell.
4. Navigation, Auswahl und Filter erzeugen keine Zielbindung und keine Autorisierung.
5. Sichtbarkeit impliziert keine Autorisierung; Ausblendung impliziert keine Nichtexistenz.
5a. Die Darstellungsdimensionen sind eine Darstellungsordnung, kein neues autoritatives Statusschema; `presentation ≠ source-of-truth ownership`. Kein Subjekt muss alle Dimensionen materialisieren, und `not applicable ≠ unknown`.
6. `capability unavailable`, `permission denied`, `restricted by operational mode` und `unknown` bleiben unterscheidbar; wo Existenz schutzbedürftig ist, wird eine einheitliche Antwort gegeben statt einer falschen Aussage.
7. Nullzustände (`empty`/`filtered-empty`/`not-collected`/`unavailable`/`denied`/`unknown`) werden nicht zu `0` oder „gesund" verschmolzen.
8. Presentation Modes verändern Dichte und Darstellung, nie das Berechtigungsergebnis; `Simple mode ≠ fewer permissions`, `Expert mode ≠ more permissions`, `Expert mode ≠ Break Glass`, `mode preference ≠ authorization`. Fehlende Modus-Präferenz führt zu einem `safe presentation fallback` (Simple/reduced-complexity presentation) — einer Darstellungsentscheidung, keinem anderen Autorisierungszustand.
9. Operational Modes bleiben benannt und unterscheidbar; `degraded ≠ failed`, `read-only UI ≠ read-only platform`, `last-known ≠ current`.
10. Zustandsarten (desired/observed/effective/reported/last-known) bleiben getrennt navigierbar; ein Sichtwechsel verändert keinen Zustand.
11. Bereichs- und Kontextidentität ist sprachunabhängig stabil; ein Anzeigename ist keine Identität.
12. CDS-Semantik ist Vergleichseingabe, keine CoreOps-Autorität; CDS `Candidate` ist keine Adoption.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 22. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): gestohlene Administratoridentität mit Experience-Einstieg THR-001; Privilege Escalation THR-002; Approval-Gate-Bypass THR-003; Job-Manipulation vor Ausführung THR-005; manipulierte Telemetrie THR-012; stale Telemetrie als aktuell behandelt THR-013; Evidence-Export an falschen Empfänger THR-018; Secret-Leak in Logs THR-019; Automation Client mit Experience-/Control-Einstieg THR-038. Keine erfundenen IDs; kein Parallelregister.

## 23. Technology Boundary

Nicht ausgewählt: Frontend-Framework · UI-Komponentenbibliothek · Routing · State Management · CSS-Architektur · Design-Tool · Icon-Set · Typografie · Farbpalette · Animations-/Motion-System · Responsive-Framework · Suchmaschine · Charting/Visualisierung · i18n-Bibliothek · Accessibility-Testwerkzeug · CDS-Paket/Tokens. Alle `deferred`. Die hier beschriebenen Muster sind konzeptionell.

## 24. Compatibility

Additiv zu [COREOPS_LOGICAL_MODULE_ARCHITECTURE.md](COREOPS_LOGICAL_MODULE_ARCHITECTURE.md) (MOD-EXP-001, DEC-S-68), [COREOPS_PLANE_TAXONOMY.md](COREOPS_PLANE_TAXONOMY.md) (Experience Plane), [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md) (Scopes/Rollen), [FOUNDATION_CAPABILITY_MATRIX.md](FOUNDATION_CAPABILITY_MATRIX.md) (Traceability), [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md) und [DOMAIN_PACK_GOVERNANCE_MODEL.md](DOMAIN_PACK_GOVERNANCE_MODEL.md). Keine bestehende Invariante geschwächt; kein Parallelmodell; keine neue Autorität. Adressiert RISK-12 (Health-Anzeige bei fehlenden Daten) und RISK-106 (UI umgeht Control/Policy) darstellungsseitig, ohne sie zu duplizieren.

## 25. Open Questions

- Balance zwischen ehrlicher Ursachenunterscheidung (§13) und Existenz-Disclosure bei schutzbedürftigen Objekten — Entscheidungsregel pro Objektklasse `deferred`.
- Konkrete Bereichszuschnitte für `future-extension`-Domänen (Topology) hängen an deren späterem Reifegrad.
- Umfang und Persistenz gespeicherter Ansichten/Personalisierung (§17) — bewusst nicht festgelegt.
- Mobile-/kleinflächige Priorisierung (`CAP-PLATFORM-007`) ist konzeptionell benannt, nicht ausgestaltet.

## 26. Next Decision

Verbindlicher nächster Schritt gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): Nova Review von `CO-WP-027`, danach Human-Maintainer-Commit. Der CDS-Reife-Re-Check bleibt vor jeder substanziellen Designübernahme erforderlich und wird durch dieses Dokument **nicht** erfüllt und **nicht** aktiviert.
