# CoreOps – UX Action Safety, Accessibility and Disclosure Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation UX action-safety, accessibility and disclosure policy
> Implementation Status: Not implemented
> Technology Selection: None
> Frontend/Component/Accessibility-Tooling/i18n Technology: Not selected
> Validation Status: Not performed
> Accessibility Validation: Not performed
> WCAG Conformance: None claimed
> Screen-Reader Support: Not proven
> Keyboard Testing: Not performed
> Certification Status: None claimed
> CDS Adoption: Not started
> CDS Pilot: Inactive / not activated by this document
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-027 (docs-only / UX information architecture and dashboard system foundation)

## 1. Status

Technologieunabhängige Policy für **risikobehaftete Aktionen in der Experience-Ebene**, **Accessibility als Designanforderung** und **Disclosure-/Secret-Grenzen der Darstellung**. Companion zu [UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md](../architecture/UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md) und [DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md](../architecture/DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md). Definiert **keinen** neuen Ausführungs-Lebenszyklus, **keine** neue Autorität und **keine** neue Approval-Instanz. Keine Frontend-, Komponenten-, Accessibility-Test- oder i18n-Technologie ausgewählt. **Keine** Accessibility-Evidenz wird durch dieses Dokument erzeugt.

## 2. Purpose

Ein Bestätigungsdialog ist keine Freigabe. Die Policy legt fest, warum `preview ≠ execution`, `confirmation ≠ approval`, `approval ≠ execution authorization`, `executed ≠ successful`, `successful ≠ verified`, `partial ≠ success`, `outcome unknown ≠ failed`, `data readable ≠ data discloseable`, `redacted view ≠ source deleted`, `evidence reference ≠ raw evidence discloseable` und `secret reference ≠ secret value` — und wie eine Oberfläche gefährliche Operationen darstellen kann, ohne Policy, Approval oder Execution Authorization zu umgehen.

## 3. Scope

Experience-Ausführungsgrenze · Darstellung risikobehafteter Operationen · Aktionspfad und Pflichtkommunikation · Bestätigungsgrenze · Preview/Plan · Scope-/Zielbindung in der Darstellung · Ergebnis-/Partial-/Unknown-/Verifikationsdarstellung · read-only-/Aktionstrennung · Simple/Expert-Autoritätsgrenze · Accessibility-Designanforderungen · DE/EN-Anforderungen · Accessibility-Claim-Grenze · Klassifikation/Disclosure in der Darstellung · Redaction · Secret-/Key-Grenze · Evidence-/Audit-Darstellung · Export-/Copy-/Rendering-Flächen · Workspace-/Environment-/Target-Isolation · Failure/Unknown · CDS Reconciliation Boundary.

## 4. Non-Goals

- Keine Frontend-/Komponenten-/Dialog-/Accessibility-Test-/i18n-/Screenreader-/Kontrastwerkzeug-Auswahl.
- Kein neuer Execution-Lebenszyklus, keine neue Approval-Instanz, keine neue Autorisierungsart, kein neues Evidence-Modell.
- Keine WCAG-Konformitätsbehauptung, keine Accessibility-Validierung, kein Screen-Reader-Nachweis, kein Keyboard-Test, keine Zertifizierung, keine Rechtsaussage.
- Keine CDS-Adoption, kein CDS-Import, keine CDS-Pilot-Aktivierung.
- Keine Behauptung implementierter oder geprüfter Sicherheits-, Bedien- oder Barrierefreiheitseigenschaften.

## 5. Concepts

`risk-bearing operation` · `dangerous action` · `intent` · `scope` · `target binding` · `preview` · `plan` · `policy evaluation` · `approval` · `execution authorization` · `execution` · `result` · `verification` · `closure` · `confirmation` · `read-only presentation` · `presentation mode` · `disclosure` · `redacted view` · `derived view` · `secret reference` · `evidence reference` · `export surface` · `copy surface` · `rendering surface`.

```text
dangerous action = Operation mit Zustandsaenderung, Verfuegbarkeits-,
                   Sicherheits-, Daten- oder Wiederherstellungsrisiko
dangerous action ≠ Operation, die die Experience-Ebene ausfuehrt
```

## 6. Experience Execution Boundary

Die Experience-Ebene (MOD-EXP-001) darf eine risikobehaftete Operation **darstellen**, **erklären** und als **Intent** weiterreichen. Sie führt sie **nicht** aus ([DEC-S-68](../../project-system/DECISION_INDEX.md); [MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md](../architecture/MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD.md); RISK-106).

```text
UI intent            ≠ approval
UI confirmation      ≠ approval record
approval             ≠ execution authorization
action visible       ≠ action authorized
action authorized    ≠ action executed
executed             ≠ successful
successful           ≠ verified
Experience layer     ≠ Execution authority
```

## 7. Risk-Bearing Operation Presentation

Eine Operation gilt in der Darstellung als risikobehaftet, wenn sie mindestens eine der folgenden Eigenschaften hat: Zielsystem-Write · Konfigurations-/Policy-/Identitätsänderung · Deployment/Update · destruktive Wirkung · Secret-/Key-Bezug · Export/Disclosure · Recovery/Rollback/Restore · Mode-Wechsel · Delegations-/Offline-Autorisierung. Die Klassifikation stammt aus den bestehenden Modellen ([SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md), [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md](DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md)); sie wird hier **nicht** neu definiert.

## 8. Dangerous-Action Path

Der darzustellende Pfad bildet die **bestehenden** Autoritätsgrenzen ab und fügt keine hinzu:

```text
intent
→ scope/context
→ preview / plan (soweit vorhanden)
→ policy evaluation
→ approval (soweit erforderlich)
→ execution authorization
→ execution
→ result
→ verification
→ evidence / closure
```

Verbindlich:

```text
preview      ≠ execution
plan         ≠ approval
confirmation ≠ approval
policy permit ≠ approval ≠ execution authorization
approval     ≠ execution authorization
execution completed ≠ successful
successful   ≠ verified
partial      ≠ success
outcome unknown ≠ failed
closed       ≠ success
```

Keine Stufe darf in der Darstellung übersprungen, zusammengezogen oder als bereits erfolgt suggeriert werden. Fehlt eine Stufe (z. B. kein Preview verfügbar), wird das **ausdrücklich** angezeigt (§9).

## 9. Required Communication for Dangerous Actions

Vor dem Absenden eines Intents muss die Darstellung mindestens kommunizieren:

| Pflichtangabe | Bedeutung |
| ------------- | --------- |
| Was sich ändert | konkrete Wirkung, nicht nur Aktionsname |
| Scope | betroffener Presentation-/Autorisierungs-Scope |
| Ziel | betroffene Ziele bzw. dass das Ziel noch zu materialisieren ist |
| Risiko/Konsequenz | erwartete Auswirkung, inkl. Verfügbarkeits-/Datenwirkung |
| Erforderliche Autorität | welche Policy-/Approval-/Authorization-Stufen nötig sind |
| Preview vorhanden? | ob ein Preview existiert — und dass Preview keine Ausführung ist |
| Wiederherstellbarkeit | ob Recovery/Rollback vorgesehen ist und mit welcher Grenze |
| Bekannte Grenzen | ungeprüfte Annahmen, unbekannte Teilzustände, Evidenzlücken |

```text
recovery vorgesehen ≠ recovery garantiert
rollback moeglich   ≠ Zustand vor Aenderung wiederhergestellt
preview vollstaendig ≠ Ausfuehrung identisch
Warnung angezeigt   ≠ Risiko akzeptiert durch Autoritaet
```

Wo eine Pflichtangabe `unknown` ist, wird `unknown` angezeigt — nicht ausgelassen und nicht wohlwollend geschätzt.

## 10. Confirmation Boundary

Eine Bestätigung in der Oberfläche ist eine **Absichtsbekundung**, kein Genehmigungsakt.

```text
confirmation dialog ≠ approval
typed confirmation  ≠ approval
double confirmation ≠ stronger authority
checkbox acknowledgement ≠ policy exception
UI confirmation     ≠ auditierbarer Approval-Record
```

Eine Bestätigung darf **niemals** als Approval-Record gespeichert, als Approval dargestellt oder anstelle einer Approval-Instanz gewertet werden ([APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md); RISK-180/RISK-190). Approval bleibt explizit, zurechenbar, scope-gebunden und widerrufbar in seinem eigenen Modell. Eine Machine Identity imitiert keine menschliche Bestätigung (THR-038; RISK-137).

## 11. Preview and Plan Presentation

```text
preview        ≠ execution
preview        ≠ Zusicherung des Ergebnisses
plan generated ≠ plan approved
plan approved  ≠ target set unchanged
kein Preview   ≠ geringes Risiko
```

Ein Preview kennzeichnet: Erhebungszeitpunkt · Datenbasis und deren Freshness · nicht ausgewertete Ziele · unbekannte Teilzustände · dass eine materielle Plan-/Zieländerung eine erneute Bewertung auslöst ([Approval Lifecycle](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md); RISK-196). Ein Preview darf **keine** Nebenwirkungen als „keine Nebenwirkungen" behaupten, wenn dies nicht belegt ist (`read-only ≠ side-effect-free`, [SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md](SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL.md); DEC-S-349).

## 12. Scope and Target Binding in the Experience Layer

```text
navigation       ≠ target binding
selection        ≠ target binding
filter           ≠ execution scope
bulk selection   ≠ bulk authorization
saved view       ≠ standing authorization
dashboard scope  ≠ execution scope
```

Ein Ausführungsziel wird **außerhalb** der Experience-Ebene materialisiert und autorisiert ([Execution Authorization](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md); [Deployment Control Plane](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), Target-Set-Snapshot und Re-Evaluation). Bei Mehrfachauswahl bleibt **Per-Ziel-Autorität** erhalten; ein Sammelvorgang erzeugt keine Sammelautorisierung.

## 13. Result, Partial, Unknown and Verification Presentation

```text
submitted   ≠ accepted
accepted    ≠ executed
executed    ≠ successful
successful  ≠ verified
partial     ≠ success ≠ failure
outcome unknown ≠ failed
no error shown  ≠ no side effects
cancelled   ≠ ohne Nebenwirkungen
closed      ≠ success
```

Ergebnisdarstellung führt **pro Ziel** Ergebnis, Verifikationsstand und Grenzen; Teil- und Unbekannt-Ergebnisse bleiben sichtbar und **blockieren** eine unsichere Wiederholung (konsistent mit [DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md](DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md) und [API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md](API_ERROR_IDEMPOTENCY_AND_REPLAY_POLICY.md); THR-026/THR-031). Ein Erfolgshinweis ohne Verifikationsaussage ist unzulässig — im Produktkonzept bereits als „keine irreführenden Erfolgsanzeigen" verankert ([COREOPS_CONCEPT_V3.md, Abschnitt UX](../architecture/COREOPS_CONCEPT_V3.md)).

## 14. Read-only and Action Separation

Übersichten und Dashboards sind read-first ([Dashboard Model §24](../architecture/DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md)). Zulässiger Pfad: `observation → inspect → explain → propose intent`.

```text
viewing data ≠ executing an action
selection    ≠ authorization
navigation   ≠ target binding
read-only presentation ≠ read-only platform
```

Aktionseinstiegspunkte in Übersichten sind zulässig, führen aber ausschließlich in den Pfad aus §8 — nie an ihm vorbei.

## 15. Simple / Expert Authority Boundary

Simple/Expert ist im autoritativen Produktmodell bereits benannt ([COREOPS_CONCEPT_V3.md, Abschnitt UX](../architecture/COREOPS_CONCEPT_V3.md)); die UX-Semantik steht in [IA §14](../architecture/UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md). Sicherheitsseitig gilt:

```text
Simple mode       ≠ different authorization model
Simple mode       ≠ fewer permissions
Expert mode       ≠ authority expansion
Expert mode       ≠ more permissions
Expert mode       ≠ Break Glass
hidden in Simple  ≠ nonexistent
expert visibility ≠ permission
mode switch       ≠ privilege change
mode preference   ≠ authorization
mode preference unavailable
  → safe presentation fallback: Simple / reduced-complexity presentation
safe presentation fallback ≠ different authorization state
safe presentation fallback ≠ reduced permissions
```

Der Modus darf **nie** das autoritative Berechtigungsergebnis, eine Policy-Bewertung, eine Approval-Anforderung oder eine Execution Authorization verändern. Der aktive Modus ist sichtbar und programmatisch bestimmbar. Ist die Modus-Präferenz nicht verfügbar oder nicht ermittelbar, greift der **safe presentation fallback** ([IA §14](../architecture/UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md)) — eine reine Darstellungsentscheidung, **kein** Autorisierungszustand und **keine** Sicherheitsmaßnahme. Expert Mode ist **keine** Break-Glass-Variante ([BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md); DEC-S-84/DEC-S-85, DEC-S-147).

## 16. Accessibility Design Requirements

Technologieunabhängige **Designanforderungen** (keine Konformitätsaussage, §18):

- **Tastaturbedienbarkeit:** jede dargestellte Funktion ist ohne Zeigegerät erreichbar und auslösbar; keine reinen Hover-/Drag-Abhängigkeiten für kritische Information.
- **Logische Fokus- und Lesereihenfolge:** Fokusreihenfolge folgt der inhaltlichen Struktur; Fokus bleibt sichtbar; Fokus wird bei Kontextwechsel deterministisch gesetzt und nicht verloren.
- **Statusbedeutung nicht nur über Farbe:** jede Statusdimension (§[Dashboard §11](../architecture/DASHBOARD_INFORMATION_HIERARCHY_AND_STATE_PRESENTATION_MODEL.md)) trägt zusätzlich Text und/oder ein nicht-farbliches Merkmal; Farbe, Icon oder Position **allein** tragen keine Bedeutung.
- **Textäquivalente:** grafische, kartografische und diagrammatische Darstellungen haben eine textliche bzw. tabellarische Entsprechung mit denselben Dimensionen.
- **Überschriften und Orientierung:** eindeutige Bereichs-, Kontext- und Ebenenkennzeichnung; der aktive Presentation Context ist programmatisch bestimmbar.
- **Klare Sprache:** Statuswerte, Einschränkungsgründe und Aktionsfolgen werden in verständlicher Sprache erklärt; Fachbegriffe werden erläutert, nicht ersetzt.
- **Navigation dichter Daten:** große Listen/Matrizen sind ohne Übersichtsverlust navigierbar; Sortierung, Filter und Position sind wahrnehmbar und angekündigt.
- **Fehleridentifikation:** Fehler benennen Ort, Ursache-Klasse und nächsten Schritt; keine reine Farb-/Symbolmarkierung.
- **Gefahren-/Aktionskommunikation:** die Pflichtangaben aus §9 sind nicht-visuell zugänglich; eine Warnung darf nicht ausschließlich visuell existieren.
- **Textexpansion DE/EN:** Layout und Struktur bleiben bei abweichender Textlänge bedeutungserhaltend; Abkürzung darf keine Dimension entfernen.
- **Unbekannte/degradierte/partielle Zustände:** `unknown`, `stale`, `partial`, `degraded`, `denied` und `unavailable` sind auch nicht-visuell unterscheidbar.
- **Reduzierte Bewegung (konzeptionell):** Bedeutung darf nicht an Bewegung oder Zeitablauf gebunden sein; eine bewegungsarme Darstellung bleibt gleichwertig.

## 17. Bilingual (DE/EN) Presentation Requirements

DE/EN ist Produktanforderung ([DEC-P-03](../../project-system/DECISION_INDEX.md)); Englisch bleibt kanonisch für maschinennahe Identifier ([DEC-S-38](../../project-system/DECISION_INDEX.md)); Übersetzungsparität ist explizit und evidenzbasiert, nicht automatisch behauptet ([DEC-S-39](../../project-system/DECISION_INDEX.md)).

```text
display label   ≠ identity
translated      ≠ parity verified
missing translation ≠ missing capability
sprachliche Vereinfachung ≠ Bedeutungsaenderung
```

Statuswerte, Modusnamen und Einschränkungsgründe behalten in beiden Sprachen dieselbe Semantik; eine Übersetzung darf eine Unterscheidung (z. B. `unknown` vs. `unavailable`) nicht einebnen.

## 18. Accessibility Claim Boundary

```text
Accessibility ist hier eine Designanforderung, keine Zertifizierung.
```

Ausdrücklich **nicht** behauptet: WCAG-Konformität · Accessibility validiert · Screen-Reader-Unterstützung nachgewiesen · Keyboard-Tests durchgeführt · Kontrastprüfung durchgeführt · Barrierefreiheits-Zertifizierung · rechtliche Konformität. Dieses Work Package erzeugt **keine** Accessibility-Evidenz. Eine spätere Prüfung erfordert reale Artefakte und eigene Evidenz ([EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](../architecture/EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md)).

## 19. Data Classification and Disclosure in Presentation

Darstellung unterliegt unverändert [DATA_CLASSIFICATION_AND_HANDLING_MODEL.md](../governance/DATA_CLASSIFICATION_AND_HANDLING_MODEL.md), [DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md](../governance/DATA_RETENTION_DELETION_AND_PRESERVATION_POLICY.md), [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md), [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md) und [AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md). Es entsteht **keine** neue Disclosure-Autorität.

```text
data readable       ≠ data discloseable
data present        ≠ user may view it
read                ≠ export
rendered            ≠ authorized to leave the boundary
unknown classification → fail-closed zur strengeren Grenze
```

## 20. Redaction and Derived View Presentation

Eine redigierte Ansicht ist eine **gebundene Derived View** ([DEC-S-340](../../project-system/DECISION_INDEX.md)); sie verändert die Quelle nicht.

```text
redacted view ≠ source deleted
redacted view ≠ disclosure-safe
masked        ≠ anonymous
pseudonymized ≠ anonymous
aggregated    ≠ anonymous
```

Die Darstellung kennzeichnet, **dass** redigiert wurde und **welcher Klasse** die Redaction folgt — ohne den redigierten Inhalt zu rekonstruieren. Restrisiko (Residual Disclosure) bleibt sichtbar benannt.

## 21. Secret and Key Boundary in Presentation

Governte Objekte tragen **Secret References**, nie Werte ([SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md); [KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md](KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md)).

```text
secret reference ≠ secret value
masked secret    ≠ non-secret
key identity     ≠ key material
credential visible ≠ credential authorized
```

**Verbot:** wiederverwendbares Secret-Material erscheint **niemals** in Dashboard, Kachel, Tooltip, Fehlermeldung, Log-Ansicht, Audit-Ansicht, Evidence-Vorschau, Kopier-Repräsentation oder Export-Repräsentation (THR-019/THR-020). Ein maskiertes Secret bleibt secret-bearing und wird durch Maskierung **nicht** herabklassifiziert.

## 22. Evidence and Audit Presentation Boundary

```text
evidence reference ≠ raw evidence discloseable
evidence visible   ≠ authority
audit view         ≠ audit alteration
view-audit         ≠ export-evidence
audit administrator ≠ unrestricted disclosure authority
recorded           ≠ validated ≠ sufficient
```

Eine Evidence-Ansicht zeigt Referenz, Verfügbarkeit, Freshness, Integritätsaussage, Provenance und Validierungsstand; die Freigabe der Rohevidenz ist ein **eigener**, autorisierter Disclosure-/Export-Akt (THR-018). Fehlende Audit-Einträge sind **keine** Aussage über Nichtausführung.

## 23. Export, Copy and Rendering Surfaces

Präsentationsnebenflächen sind Disclosure-Flächen und unterliegen denselben Grenzen wie Exporte ([REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md §11](REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md)):

| Fläche | Behandlung |
| ------ | ---------- |
| Kopieren/Zwischenablage | wie Export; Redaction vor Erzeugung |
| Datei-/Berichtsexport | eigener Disclosure-/Export-Autorisierungsakt |
| Druck-/Rendering-Ausgabe | Redaction vor Rendering |
| Tooltip/Detailfläche | keine Erweiterung der Disclosure-Grenze |
| Fehlermeldung/Diagnose | keine Secret-/Klassifikationsleckage |
| Deep Link/geteilte Ansicht | erzeugt keinen Zugriff; keine sensiblen Daten im Link |
| Offline-/Snapshot-Ansicht | Provenance und Klassifikation bleiben gebunden (THR-025) |

```text
sichtbar ≠ exportierbar
copy     ≠ read
render   ≠ authorized disclosure
link     ≠ access grant
```

## 24. Workspace, Environment and Target Isolation in Presentation

```text
cross-workspace view    ≠ cross-workspace authority
aggregated across scopes ≠ zulaessige Zusammenfuehrung
environment label       ≠ isolation boundary
target visible          ≠ target bound
```

Eine übergreifende Darstellung darf Isolationsgrenzen nicht stillschweigend aufheben ([WORKSPACE_RBAC_AND_SCOPE_MODEL.md](WORKSPACE_RBAC_AND_SCOPE_MODEL.md); [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md §16](REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md)). Wo eine Zusammenführung unzulässig ist, wird dies ausgewiesen statt teilweise dargestellt.

## 25. Failure and Unknown State

```text
unknown state       → fail-closed
unknown permission  → keine Aktionsdarstellung als verfuegbar
unknown scope       → engerer Scope
unknown outcome     → sichtbar; kein automatischer Retry
unknown classification → strengere Grenze
```

Bei unklarem Zustand wird die restriktivere Darstellung gewählt und der Grund benannt; eine Oberfläche „rät" keinen sicheren Zustand.

**Abgrenzung:** Dieses Fail-closed-Verhalten betrifft **unbekannte Zustands-, Berechtigungs-, Scope-, Ergebnis- oder Klassifikationswerte**. Es ist etwas anderes als der `safe presentation fallback` bei fehlender Modus-Präferenz (§15), der ausschließlich Informationsdichte betrifft. In beiden Fällen gilt: eine restriktivere **Darstellung** ist **kein** anderer Autorisierungszustand — sie zeigt weniger, sie entzieht nichts.

```text
restriktivere Darstellung ≠ reduzierte Berechtigung
restriktivere Darstellung ≠ anderer Autorisierungszustand
unknown-state fail-closed ≠ safe presentation fallback (§15)
```

## 26. CDS Reconciliation Boundary

Externer CDS-Stand und Adoptionsverbot: siehe [IA §20](../architecture/UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md) — dieselbe Grenze gilt hier unverändert und wird nicht dupliziert. Policy-spezifische Einordnung:

| Befund | Klassifikation |
| ------ | -------------- |
| „Farbe/Icon/Position allein tragen keine Statusbedeutung"; `unknown` explizit; Evidenzgrenzen sichtbar | **compatible overlap** |
| Gefahren-Aktionspfad mit Policy/Approval/Execution-Authorization-Trennung | **CoreOps-local** |
| Secret-/Evidence-/Disclosure-Grenzen der Darstellungsflächen | **CoreOps-local** |
| Simple/Expert als reine Darstellungsdimension ohne Autoritätswirkung | **CoreOps-local** |
| Spätere Nutzung von CDS-Statusdarstellungsregeln für die Achsen `condition`/`severity`/`confidence`/`freshness`/`evidence` | **future mapping candidate** |
| Übernahme von CDS-Komponenten, -Tokens oder -Dialogmustern für gefährliche Aktionen | **not adopted — requires future CDS pilot evidence** |

```text
CDS Candidate ≠ CoreOps adoption
CDS semantics ≠ CoreOps domain authority
```

Kein CDS-Import, keine Tokens, kein Pilot, keine Adoption-Evidenz.

## 27. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Die Experience-Ebene stellt risikobehaftete Operationen dar und routet Intents; sie führt sie nicht aus.
2. Eine UI-Bestätigung ist keine Approval und wird nie als Approval-Record gewertet oder gespeichert.
3. `preview ≠ execution`; ein fehlendes Preview wird ausgewiesen und bedeutet kein geringeres Risiko.
4. `policy permit ≠ approval ≠ execution authorization`; keine Stufe wird in der Darstellung übersprungen oder als erfolgt suggeriert.
5. Auswahl, Filter und Navigation erzeugen keine Zielbindung und keine Sammelautorisierung; Per-Ziel-Autorität bleibt erhalten.
6. `executed ≠ successful ≠ verified`; Teil- und Unbekannt-Ergebnisse bleiben sichtbar und blockieren unsichere Wiederholung.
7. Presentation Modes verändern nie das autoritative Berechtigungsergebnis; `Simple mode ≠ fewer permissions`, `Expert mode ≠ more permissions`, `mode preference ≠ authorization`; Expert Mode ist keine Break-Glass-Variante. Der `safe presentation fallback` bei fehlender Modus-Präferenz ist eine Darstellungsentscheidung, kein anderer Autorisierungszustand.
8. Accessibility ist eine Designanforderung; keine Konformitäts-, Validierungs- oder Zertifizierungsaussage wird getroffen.
9. Statusbedeutung wird nie ausschließlich über Farbe, Icon, Position oder Bewegung getragen.
10. `data readable ≠ data discloseable`; `read ≠ export`; unbekannte Klassifikation ist fail-closed zur strengeren Grenze.
11. Eine redigierte Ansicht löscht die Quelle nicht und ist nicht automatisch disclosure-sicher; maskiert ist nicht anonym.
12. Secret References ersetzen Werte; wiederverwendbares Secret-Material erscheint auf keiner Darstellungs-, Kopier- oder Exportfläche.
13. Eine Evidence-Referenz ist keine Freigabe der Rohevidenz; eine Audit-Ansicht ist keine Audit-Änderung.
14. Übergreifende Darstellungen heben Workspace-/Environment-/Target-Isolation nicht stillschweigend auf.
15. Unbekannte Zustände führen zur restriktiveren Darstellung mit benanntem Grund.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 28. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): gestohlene Administratoridentität mit Experience-Einstieg THR-001; Privilege Escalation THR-002; Approval-Gate-Bypass THR-003; Job-Manipulation vor Ausführung THR-005; Evidence-Export an falschen Empfänger THR-018; Secret-Leak in Logs THR-019; Secret-Leak in Exports THR-020; Offline-Export mit sensiblen Daten THR-025; Replay eines zuvor autorisierten Kommandos THR-026; Teil-Job-Fehler ohne sicheren Zustand THR-031; privilegierter Insider THR-037; Automation Client mit Experience-/API-Einstieg THR-038. Keine erfundenen IDs; kein Parallelregister.

## 29. Technology Boundary

Nicht ausgewählt: Frontend-/Komponenten-Framework · Dialog-/Modal-Muster-Bibliothek · Formular-/Validierungsbibliothek · i18n-/Übersetzungswerkzeug · Accessibility-Testwerkzeug · Screenreader-Testumgebung · Kontrast-/Farbwerkzeug · Export-/PDF-/Rendering-Pipeline · Clipboard-/Sharing-Mechanismus · CDS-Paket/Tokens. Alle `deferred`.

## 30. Compatibility

Additiv zu [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md), [POLICY_DECISION_AND_EVALUATION_MODEL.md](POLICY_DECISION_AND_EVALUATION_MODEL.md), [SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md), [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](WORKSPACE_RBAC_AND_SCOPE_MODEL.md), [REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md](REDACTION_MINIMIZATION_AND_CONTROLLED_DISCLOSURE_POLICY.md), [SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md](SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE.md), [AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md) und [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](../architecture/DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md). Konkretisiert [DEC-S-68](../../project-system/DECISION_INDEX.md) und [DEC-G-04](../../project-system/DECISION_INDEX.md) darstellungsseitig. Keine bestehende Invariante geschwächt; kein neuer Lebenszyklus; keine neue Autorität. Adressiert RISK-106, RISK-180 und RISK-190 darstellungsseitig, ohne sie zu duplizieren.

## 31. Open Questions

- Zulässige Tiefe einer Verweigerungsbegründung ohne Existenz-Disclosure (gemeinsam mit [IA §25](../architecture/UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md), `deferred`).
- Umgang mit Kopier-/Zwischenablageflächen, die technisch außerhalb der Plattformkontrolle liegen (`deferred`; hier nur als Grenze benannt).
- Konkrete Accessibility-Prüfkriterien und Evidenzanforderungen für eine spätere Bewertung (bewusst nicht festgelegt, keine Konformitätsaussage).
- Verhältnis zwischen Simple Mode und Notfall-/Recovery-Situationen (bewusst nicht als Break-Glass-Variante ausgestaltet).

## 32. Next Decision

Verbindlicher nächster Schritt gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): Nova Review von `CO-WP-027`, danach Human-Maintainer-Commit. Der CDS-Reife-Re-Check bleibt vor jeder substanziellen Designübernahme erforderlich und wird durch dieses Dokument **nicht** erfüllt und **nicht** aktiviert.
