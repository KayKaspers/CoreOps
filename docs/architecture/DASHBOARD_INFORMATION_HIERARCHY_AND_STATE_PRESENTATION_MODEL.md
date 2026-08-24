# CoreOps – Dashboard Information Hierarchy and State Presentation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation dashboard information hierarchy and state presentation model
> Implementation Status: Not implemented
> Technology Selection: None
> Widget/Charting/Visualization/Dashboard-Engine Technology: Not selected
> Validation Status: Not performed
> Usability Validation: Not performed
> Certification Status: None claimed
> CDS Adoption: Not started
> CDS Pilot: Inactive / not activated by this document
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-027 (docs-only / UX information architecture and dashboard system foundation)

## 1. Status

Technologieunabhängiges **Informations- und Darstellungsmodell** für CoreOps-Dashboards und -Übersichten: Autoritätsgrenze, Scope-Bindung, Informationshierarchie, Priorisierung, elf getrennte Statusdimensionen, Freshness/Confidence/Evidence, Abdeckung, Konflikte, Capability-/Permission-Sichtbarkeit, Drill-down und Provenance-Zugang. Companion zu [UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md) und [UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md). **Kein** Widget-Framework, **keine** Komponenten-, Charting- oder Dashboard-Engine ausgewählt; **kein** zweites System of Record.

## 2. Purpose

Eine Zusammenfassung wird leicht zur Wahrheit erklärt. Das Modell legt fest, warum `dashboard representation ≠ authoritative domain state`, `dashboard summary ≠ authoritative state`, `summary ≠ completeness`, `aggregate ≠ validated population`, `unknown ≠ healthy`, `stale ≠ current`, `partial ≠ complete`, `degraded ≠ failed` und `unverified ≠ verified` — und wie dichte Betriebsübersichten priorisieren dürfen, **ohne** die zugrunde liegenden Dimensionen zu löschen.

## 3. Scope

Dashboard-Autoritätsgrenze · Zweck und Scope-Bindung · Informationshierarchie Overview/Detail · Region-/Panel-/Summary-Item-Semantik ohne Komponentenwahl · Priorisierung, Sortierung, Gruppierung · elf Statusdimensionen · Freshness · Confidence/Validation · Evidence-Verfügbarkeit und Provenance-Zugang · Abdeckung/Partial/Unknown · Konfliktsichtbarkeit · Capability-/Permission-Darstellung · Operational Mode · Composite-Indicator- und CoreScore-Grenze · last-known vs current · Leer-/Fehler-/Offline-Zustände · Dashboard-Profile · read-first-Grenze · CDS Reconciliation Boundary.

## 4. Non-Goals

- Kein Widget-Framework, keine Komponentenbibliothek, keine Charting-/Visualisierungs-/Dashboard-Engine, kein Layout-System, keine Farbpalette, kein Icon-Set, keine Typografie, kein Animationssystem.
- Keine Query-/Aggregations-/Zeitreihen-/Alerting-/Storage-Technologie; keine API/Datenbank/Runtime.
- Keine CDS-Token-Übernahme, kein CDS-Paket, keine CDS-Pilot-Aktivierung.
- Keine Mockups, keine Screenshots, keine Behauptung implementierter, validierter oder performanter Dashboards.

## 5. Concepts

`dashboard` · `dashboard profile` · `presentation region` · `summary item` · `detail view` · `scope binding` · `status dimension` · `operational condition` · `severity` · `confidence` · `freshness` · `staleness` · `evidence availability` · `data coverage` · `completeness` · `capability availability` · `permission state` · `operational mode` · `conflict state` · `verification state` · `composite indicator` · `last-known value` · `provenance path` · `drill-down`.

```text
dashboard = scope-gebundene, abgeleitete Darstellung
            ueber bereits autoritativ gehaltenen Daten
dashboard ≠ System of Record
dashboard ≠ Aggregationsautoritaet
```

## 6. Dashboard Authority Boundary

Ein Dashboard ist eine **abgeleitete Darstellung** im Sinne der bestehenden Source-of-Truth-Governance ([DEC-S-42](../../project-system/DECISION_INDEX.md): `summaries and derived artifacts must not override authoritative sources`; [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md)). Es besitzt **keinen** autoritativen Zustand, erzeugt **keine** Aggregationsautorität und wird **niemals** ein zweites System of Record.

```text
dashboard representation ≠ authoritative domain state
dashboard summary        ≠ authoritative state
dashboard value          ≠ verified value
dashboard freshness      ≠ source freshness guarantee
dashboard aggregate      ≠ validated population
evidence visible         ≠ authority
```

**Verbindliche Invariante:** Eine Zusammenfassung darf **priorisieren**. Sie darf die zugrunde liegenden Dimensionen **nicht löschen**. Jede zusammengefasste Aussage muss auf ihre Dimensionen und ihre Quelle zurückführbar sein (§21).

## 7. Dashboard Purpose and Scope Binding

Jedes Dashboard trägt einen **deklarierten Zweck** und eine **deklarierte Scope-Bindung** aus den acht Presentation-Context-Scopes ([IA §7](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md)). Fehlt die Bindung, gilt fail-closed der engere Scope.

```text
dashboard scope              ≠ execution scope
dashboard filter             ≠ execution scope
dashboard scope              ≠ Berechtigungsumfang
presentation context         ≠ authorization scope
referencing a scope identity ≠ performing authorization
context narrowing            ≠ permission mutation
UI visibility                ≠ permission result
global dashboard             ≠ globale Autoritaet
scope unknown → fail-closed auf engeren Scope
```

Die Scope-Bindung eines Dashboards ist eine **Sichtbindung** auf eine bestehende Scope-Identität; die Autorisierung bleibt bei MOD-IAM-001 / MOD-POL-001 ([IA §7.1](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md)). Ein Dashboard erzeugt, erweitert oder interpretiert **keinen** Scope und **keine** Berechtigung.

Zweckklassen (konzeptionell): Orientierung · Priorisierung/Triage · Zustandsvergleich · Abdeckungsprüfung · Betriebsmodus-/Schutzlage · Evidenz-/Auditsicht. Ein Dashboard, das mehrere Zwecke vermischt, muss die Zwecke pro Region kenntlich halten.

## 8. Information Hierarchy

| Ebene | Verantwortung | Ausdrücklich nicht |
| ----- | ------------- | ------------------ |
| Overview | priorisierte Orientierung im deklarierten Scope, Abdeckungs- und Freshness-Anzeige | Vollständigkeit, autoritative Aussage, Beweis |
| List | Abgrenzung, Vergleich, Filterung, Per-Objekt-Dimensionen | Beweis der Grundgesamtheit |
| Detail | vollständige Dimensionen, Verlauf, Grenzen, Provenance | Autoritätsgewährung |
| Evidence | Evidence-Referenzen, Verfügbarkeit, Validierungsstand, Disclosure-Grenzen | automatische Rohevidenz-Freigabe |

```text
overview ist verantwortlich fuer Priorisierung, nicht fuer Vollstaendigkeit
detail   ist verantwortlich fuer Dimensionen, nicht fuer Autoritaet
```

## 9. Presentation Region Semantics

Technologieunabhängige Struktureinheiten — **keine** Komponentenauswahl, **keine** Layoutvorgabe:

- **Presentation Region:** benannter Bereich mit eigenem Zweck, eigener Scope-Bindung und eigener Freshness-Aussage.
- **Summary Item:** kleinste zusammenfassende Aussage; trägt Wert **plus** Dimensionen (§11) **plus** Abdeckung **plus** Freshness.
- **Detail Surface:** vollständige Dimensionsdarstellung mit Provenance-Zugang.

```text
region        ≠ Modul
region        ≠ Datenowner
summary item  ≠ Kennzahl ohne Kontext
region layout ≠ Bedeutungstraeger
```

Ein `summary item` **ohne** Abdeckungs- und Freshness-Aussage ist unzulässig. Die im Produktkonzept genannten Widget-Typen ([COREOPS_CONCEPT_V3.md, Abschnitt Widget-Typen](COREOPS_CONCEPT_V3.md)) bleiben **Darstellungsformen**, nie Statussemantik — insbesondere ist eine „Ampel" eine Darstellungsform und **kein** Statusmodell (§11, §19).

## 10. Prioritization, Sorting and Grouping

Priorisierung ist erlaubt und notwendig; sie ist **erklärbar** und **verlustfrei bezüglich Dimensionen**.

```text
priority     ≠ severity
priority     ≠ authority
sorted first ≠ most important truth
grouping     ≠ Aussage ueber Ursache
hidden by prioritization ≠ nonexistent
```

Regeln: die angewandte Priorisierungs-/Sortier-/Gruppierungsregel ist sichtbar und erklärbar · durch Priorisierung nicht gezeigte Objekte werden mit Anzahl und Zugang ausgewiesen (`hidden ≠ nonexistent`) · `unknown` und `stale` dürfen **nicht** ans Ende sortiert und damit unsichtbar gemacht werden · Sortierung nach einem einzelnen abgeleiteten Wert ersetzt keine Dimensionsanzeige.

## 11. Status Presentation Dimensions

CoreOps kennt **keinen** globalen Gesundheits-Boolean und **keinen** normativen Aggregat-Score. Elf **Darstellungsdimensionen** bleiben **getrennt** darstellbar und **getrennt** filterbar. Sie sind eine **Darstellungsordnung über bereits bestehenden autoritativen Modellen** — **kein** neues Statusschema (§11.1):

| # | Dimension | Herkunft (bestehendes Modell) |
| - | --------- | ----------------------------- |
| 1 | `operational condition` | [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) |
| 2 | `severity / impact` | [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md), [DEC-S-57](../../project-system/DECISION_INDEX.md) (qualitativ, evidence-bounded) |
| 3 | `confidence` | [TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md §9](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md) |
| 4 | `freshness / staleness` | [Evidence Model §12](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [Topology §15](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md) |
| 5 | `evidence availability` | [Evidence Model §11](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) |
| 6 | `data coverage / completeness` | [COREOPS_CONCEPT_V3.md, Abschnitt CoreScore](COREOPS_CONCEPT_V3.md), [Topology §16](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md) |
| 7 | `capability availability` | [FOUNDATION_CAPABILITY_MATRIX.md](FOUNDATION_CAPABILITY_MATRIX.md), [INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md](INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md) |
| 8 | `permission state` | [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md) |
| 9 | `operational mode` | [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md §6](DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md) |
| 10 | `conflict state` | [Topology §17–§18](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md), [State Model](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) |
| 11 | `verification state` | [Evidence Model §16](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md](../security/SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md) |

```text
unknown             ≠ healthy
stale               ≠ current
unverified          ≠ verified
degraded            ≠ failed
permission denied   ≠ capability unavailable
capability unavailable ≠ provider failure
provider/process healthy ≠ consumer capability usable
low confidence      ≠ wrong
high confidence     ≠ authoritative truth
conflicted          ≠ resolved
```

**Verbot:** Diese Dimensionen dürfen **nicht** zu `green`/`yellow`/`red` oder zu einem einzelnen Score kollabiert werden. Eine visuelle Zusammenfassung darf existieren, **muss** aber (a) als abgeleitet gekennzeichnet sein, (b) die beitragenden Dimensionen zugänglich halten und (c) ihre eigenen Grenzen benennen.

Der reale Betriebsbefund aus dem vorangegangenen Milestone bleibt Architektureingabe: **pfadbezogene, geschichtete Health** ist bereits abgedeckt; eine **consumer-gebundene Dependency-Contract-Health** ist nur *teilweise* abgedeckt und bleibt eine spätere, eng begrenzte Erweiterung. Dieses Dokument **implementiert diese Erweiterung nicht**; es stellt lediglich sicher, dass Dimension 7 und die Trennung `provider/process healthy ≠ consumer capability usable` eine solche Unterscheidung später **darstellbar** machen.

### 11.1 Presentation-only Boundary

Die elf Dimensionen sind **ausschließlich eine Darstellungsordnung**. Sie sind **kein** neues autoritatives CoreOps-Statusschema, **kein** Statusobjekt, **kein** Lebenszyklus und **kein** Zustandsmodell.

```text
UX presentation dimensions ≠ new authoritative status schema
presentation dimension     ≠ status object
presentation dimension     ≠ lifecycle or state model
presentation              ≠ source-of-truth ownership
presentation order        ≠ Datenmodell
```

Verbindlich:

- **Jede Dimension bleibt durch ihr bestehendes autoritatives CoreOps-Modell governt** (Spalte „Herkunft"). Dieses Dokument definiert für keine Dimension Werte, Enums, Übergänge, Präzedenz oder Ownership neu und erzeugt **kein** Parallelmodell.
- **Kein Subjekt muss alle elf Dimensionen materialisieren.** Es gibt **kein** verpflichtendes Elf-Felder-Schema und **keine** universelle Statusstruktur, die jedes Objekt zu tragen hätte. Welche Dimensionen für ein Subjekt überhaupt gelten, ergibt sich aus dem jeweils zuständigen autoritativen Modell — nicht aus dieser Darstellungsordnung.
- **Anwendbarkeit und Wert sind getrennt.** Eine für ein Subjekt nicht anwendbare Dimension ist **kein** unbekannter Wert:

```text
not applicable ≠ unknown
dimension not applicable to a subject
   ≠ applicable dimension whose value is unknown
not applicable ≠ healthy
not applicable ≠ absent data
unknown        ≠ healthy
```

- Wird eine Dimension als **nicht anwendbar** dargestellt, ist dies als Anwendbarkeitsaussage kenntlich zu machen und darf **nicht** als `unknown`, `0`, „gesund" oder als Datenlücke gerendert werden. Umgekehrt darf ein anwendbarer, aber unbekannter Wert **nicht** als „nicht anwendbar" verborgen werden.
- Die Anwendbarkeit einer Dimension ist eine Aussage der jeweils zuständigen autoritativen Quelle, **nicht** eine Entscheidung der Darstellungsebene.

## 12. Freshness and Temporal Presentation

Vier bereits getrennte Zeitdimensionen ([EVENT_AND_AUDIT_CORRELATION_MODEL.md](EVENT_AND_AUDIT_CORRELATION_MODEL.md), [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md)) bleiben unterscheidbar: Beobachtungszeit · Aufzeichnungszeit · Ingestionszeit · Darstellungszeit.

```text
rendered now      ≠ observed now
last refresh      ≠ last observation
no update         ≠ no change
timestamp present ≠ trusted time
```

Jede Region trägt eine Freshness-Aussage; wo Freshness `unknown` ist, wird das ausgewiesen und **nicht** als aktuell dargestellt. Zeitangaben ohne belastbare Zeitquelle bleiben als solche kenntlich (THR-027).

## 13. Confidence and Validation Presentation

Confidence-Werte (`not-assessed` · `low` · `medium` · `high` · `source-reported` · `derived` · `conflicted` · `unknown`) und Validation-Werte (`not-assessed` · `unvalidated` · `validated-with-notes` · `validated` · `invalid` · `stale` · `superseded` · `conflicted` · `unknown`) werden **unverändert** aus [TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md §9/§10](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md) übernommen. Es entsteht **keine** neue Skala und **keine** erzwungene Numerik.

```text
confidence ≠ validation
validated  ≠ currently fresh
derived    ≠ measured
source-reported ≠ independently confirmed
```

## 14. Evidence Availability and Provenance Access

Die sechs Evidence-Dimensionen aus [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) (Availability · Freshness · Integrity · Provenance · Validation · Sufficiency) bleiben getrennt.

```text
evidence capability ≠ available ≠ current
integrity stated    ≠ verified
validated           ≠ sufficient
sufficient          ≠ compliance certified
evidence reference  ≠ raw evidence discloseable
```

Jede dargestellte Aussage bietet einen **Provenance-Pfad** ([FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md)): Quelle · Erhebungsart · Transformationen · Zeitpunkte · Grenzen. Der Zugang zur Evidenz selbst bleibt disclosure-gebunden ([UX Policy §22](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md)).

## 15. Coverage, Partial and Unknown Data

Abdeckung ist eine **Pflichtangabe**, keine Option (`Fehlende Daten dürfen nicht automatisch als gesund gewertet werden` — [COREOPS_CONCEPT_V3.md, CoreScore](COREOPS_CONCEPT_V3.md); [DEC-P-06](../../project-system/DECISION_INDEX.md)).

```text
partial          ≠ complete
covered scope    ≠ full population
collector healthy ≠ every source reported
no data          ≠ zero ≠ healthy
sampling         ≠ full observation
aggregation      ≠ validated population
```

Jede Aggregation weist aus: einbezogene Objekte · nicht einbezogene Objekte · nicht erhobene Objekte · Objekte mit `unknown` · Objekte außerhalb der Berechtigung (als Anzahl, soweit disclosure-zulässig). Fehlt eine dieser Angaben, ist die Aggregation als **unvollständig** zu kennzeichnen.

## 16. Conflict Presentation

Konflikte bleiben **sichtbar** und werden nicht still aufgelöst ([TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md §18](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md)).

```text
conflicted ≠ resolved
newest     ≠ correct
kein last-write-wins in der Darstellung
keine Timestamp-Praezedenz allein
reconciled ≠ conflict-free automatically
```

Ein Dashboard darf einen Konflikt **priorisieren**, aber nicht **verbergen**; die konkurrierenden Aussagen und ihre Quellen bleiben erreichbar.

## 17. Capability Availability and Permission Presentation

Die vier Ursachenklassen aus [IA §13](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md) (`capability unavailable` · `permission denied` · `state unknown` · `restricted by operational mode`) gelten unverändert und werden **nicht** zu einem leeren Feld verschmolzen.

```text
leere Kachel ≠ Aussage
kein Wert    ≠ Wert null
nicht sichtbar ≠ nicht vorhanden
gesperrt     ≠ ausgefallen
```

Zusätzlich bleiben die sechs Capability-Dimensionen aus dem Integrationsvertrag (`advertised` · `detected` · `permitted` · `implemented` · `supported` · `validated`, [INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md](INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md)) getrennt darstellbar; `advertised ≠ validated`.

## 18. Operational Mode and Degraded Presentation

Der aktive Operational Mode ist eine **eigene Dimension** (§11 Nr. 9), kein Fehlerzustand. Durch einen Mode suspendierte Fähigkeiten erscheinen als **suspendiert**, nicht als fehlend oder ausgefallen.

```text
degraded mode ≠ failure
mode entered  ≠ root cause resolved
read-only mode ≠ platform read-only
containment   ≠ recovery
unknown operational state → fail-closed, keine Gesundheitsaussage
```

## 19. Composite Indicators and CoreScore Boundary

Das Produktkonzept sieht einen zusammengesetzten Gesundheitswert (**CoreScore**) für Assets, Anwendungen und Instanzen vor ([COREOPS_CONCEPT_V3.md, Abschnitt CoreScore](COREOPS_CONCEPT_V3.md)). Dieses Dokument bestätigt ihn als **objektbezogenen abgeleiteten Indikator** und begrenzt ihn:

```text
CoreScore ist objektbezogen (Asset/Anwendung/Instanz)
CoreScore ≠ globaler Plattform-Gesundheitswert
CoreScore ≠ normativer Aggregat-Score ueber Domaenen
CoreScore ≠ Ersatz fuer die elf Statusdimensionen
score present ≠ evidence sufficient
```

Ein CoreScore muss die im Konzept geforderten Begleitangaben zeigen — Datenabdeckung · letzte Aktualisierung · nicht unterstützte Prüfungen · ignorierte Findings · manuelle Overrides · Datenquellen — und darf **niemals** fehlende Daten als gesund werten. Kein weiterer, breiterer Aggregatwert wird eingeführt.

## 20. Last-Known versus Current

```text
last-known value ≠ current value
last-known state ≠ current state
cached view      ≠ live view
offline view     ≠ failure view
absence of update ≠ stability
```

Last-known-Darstellungen tragen Erhebungszeitpunkt, Quelle und Grund der Nichtaktualität. Sie werden **nicht** stillschweigend als aktuell gerendert.

## 21. Drill-down and Provenance Path

Jede zusammengefasste Aussage bietet einen deterministischen Weg: **Summary Item → beitragende Objekte → Objektdetail mit Dimensionen → Provenance → Evidence-Referenz (disclosure-gebunden)**. Der Presentation Context wird dabei mitgeführt ([IA §10](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md)).

```text
drill-down    ≠ Autoritaetsgewinn
drill-down    ≠ Disclosure-Erweiterung
provenance    ≠ evidence disclosure
Weg vorhanden ≠ Zugriff gewaehrt
```

Ist der Drill-down aus Berechtigungs- oder Disclosure-Gründen nicht möglich, wird die Ursache klassifiziert (§17), nicht verschwiegen.

## 22. Empty, Error and Offline States

Die sechs Nullzustände aus [IA §16](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md) gelten unverändert. Ergänzend für Dashboards:

| Zustand | Erforderliche Aussage |
| ------- | --------------------- |
| `empty` | Scope, Abfragezeitpunkt, Hinweis dass Scope leer ist |
| `filtered-empty` | angewandter Filter, Weg zum Zurücksetzen |
| `not-collected` | dass nie erhoben wurde, nicht dass nichts vorliegt |
| `unavailable` | betroffene Quelle/Fähigkeit, letzter bekannter Stand |
| `denied` | Berechtigungsursache, soweit disclosure-zulässig |
| `unknown` | ausdrücklich `unknown`, nie `0` und nie „gesund" |
| `error` | Fehlerklasse, Scope, ob Teilergebnisse enthalten sind |
| `offline` | Konnektivitätsklasse, `last-known`-Zeitpunkt |

```text
error in one region ≠ error of the whole dashboard
partial result      ≠ complete result
partial result      ≠ failure
```

## 23. Dashboard Profiles and Personalization Boundary

Die im Konzept genannten Dashboard-Profile ([COREOPS_CONCEPT_V3.md, Abschnitt Dashboard-Profile](COREOPS_CONCEPT_V3.md)) bleiben **Darstellungsprofile**. Sie führen keine Autorität, keine Datenownership und keine Berechtigung mit sich.

```text
profile     ≠ role
profile     ≠ permission
profile     ≠ scope grant
personalization ≠ authorization
saved dashboard ≠ standing authorization
```

Ein Profil kann Regionen priorisieren; es darf keine Dimension dauerhaft entfernen, ohne dies auszuweisen (`hidden ≠ nonexistent`).

## 24. Read-first and Action Entry Boundary

Dashboards und Übersichten sind **read-first**. Der zulässige konzeptionelle Pfad lautet:

```text
observation → inspect → explain → propose intent
```

Erst danach greift der bestehende privilegierte Pfad ([UX Policy §7–§13](../security/UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md)).

```text
viewing data      ≠ executing an action
selection         ≠ authorization
navigation        ≠ target binding
dashboard filter  ≠ execution scope
bulk selection    ≠ bulk authorization
```

Ein Ausführungsziel muss weiterhin durch die bestehenden operativen Modelle materialisiert und autorisiert werden; ein Dashboard erzeugt es nicht.

## 25. CDS Reconciliation Boundary

Externer CDS-Stand, Adoptionsverbot und Achsenvergleich: siehe [IA §20](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md) — dieselbe Grenze gilt hier unverändert und wird nicht dupliziert. Dashboard-spezifische Einordnung:

| Befund | Klassifikation |
| ------ | -------------- |
| Keine normative Aggregat-Health; Achsen bleiben unabhängig; `unknown` explizit; Evidenzgrenzen sichtbar; Farbe/Icon/Position allein tragen keine Bedeutung | **compatible overlap** |
| Elf CoreOps-Statusdimensionen inkl. `data coverage`, `capability availability`, `permission state`, `operational mode`, `verification state` | **CoreOps-local** |
| CoreScore als objektbezogener Indikator mit Pflichtbegleitangaben | **CoreOps-local** |
| Spätere Abbildung `condition`/`severity`/`confidence`/`freshness`/`evidence` auf die Dimensionen 1/2/3/4/5 | **future mapping candidate** |
| Reduktion der elf Dimensionen auf fünf Achsen | **would lose meaning — not adopted** |
| Übernahme von CDS-Statuskomponenten, -Tokens oder -Darstellungsregeln | **requires future CDS pilot evidence** |

```text
CDS Candidate ≠ CoreOps adoption
CDS semantics ≠ CoreOps domain authority
```

Kein CDS-Import, keine Tokens, kein Pilot, keine Adoption-Evidenz.

## 26. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Eine Dashboard-Darstellung ist abgeleitet und niemals autoritativer Domänenzustand oder zweites System of Record.
2. Eine Zusammenfassung darf priorisieren, aber keine zugrunde liegende Dimension löschen.
3. Jede Dashboard-Region trägt eine deklarierte Scope-Bindung; unbekannter Scope ist fail-closed.
4. Die elf Statusdimensionen bleiben getrennt darstellbar; kein globaler Gesundheits-Boolean und kein normativer Aggregat-Score.
4a. Die elf Dimensionen sind eine Darstellungsordnung, **kein** neues autoritatives Statusschema, Statusobjekt, Zustandsmodell oder Lebenszyklus; `UX presentation dimensions ≠ new authoritative status schema`; `presentation ≠ source-of-truth ownership`. Jede Dimension bleibt durch ihr bestehendes autoritatives Modell governt.
4b. Kein Subjekt muss alle elf Dimensionen materialisieren; es existiert kein verpflichtendes Elf-Felder-Schema.
4c. `not applicable ≠ unknown`; eine für ein Subjekt nicht anwendbare Dimension ist kein unbekannter Wert, keine Datenlücke und keine Gesundheitsaussage; Anwendbarkeit entscheidet die zuständige autoritative Quelle, nicht die Darstellungsebene.
5. `unknown ≠ healthy`, `stale ≠ current`, `unverified ≠ verified`, `partial ≠ complete`, `degraded ≠ failed`, `conflicted ≠ resolved`.
6. `permission denied ≠ capability unavailable ≠ provider failure ≠ restricted by mode`; ein leeres Feld ist keine Aussage.
7. Abdeckung, Freshness und Nicht-Einbezogenes sind Pflichtangaben jeder Aggregation.
8. Konflikte bleiben sichtbar; kein Last-Write-Wins und keine Timestamp-Präzedenz allein in der Darstellung.
9. Durch Priorisierung oder Profil nicht gezeigte Objekte sind ausgewiesen und erreichbar; `hidden ≠ nonexistent`.
10. Ein CoreScore bleibt objektbezogen, abgeleitet und begleitangabenpflichtig; er ersetzt keine Dimension und wertet fehlende Daten nicht als gesund.
11. `last-known` wird nie still als `current` dargestellt.
12. Dashboards sind read-first; Auswahl, Filter und Navigation erzeugen weder Autorisierung noch Zielbindung.
13. Ein Dashboard-Profil trägt keine Rolle, keine Berechtigung und keinen Scope-Grant.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 27. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): manipulierte Telemetrie THR-012; stale Telemetrie als aktuell behandelt THR-013; Audit-Löschung/-Manipulation THR-016/THR-017; Evidence-Export an falschen Empfänger THR-018; Secret-Leak in Logs THR-019; Secret-Leak in Exports THR-020; falsche Zeit THR-027; managed resource kompromittiert CoreOps THR-034; privilegierter Insider THR-037. Keine erfundenen IDs; kein Parallelregister.

## 28. Technology Boundary

Nicht ausgewählt: Dashboard-Engine · Widget-/Komponenten-Framework · Charting-/Visualisierungsbibliothek · Layout-/Grid-System · Query-/Aggregations-Engine · Zeitreihen-Storage · Alerting · Caching · Streaming/Push · Export-/Rendering-Pipeline · Farb-/Icon-/Typografie-System · CDS-Paket/Tokens. Alle `deferred`. `DEC-S-254` (Telemetry-/Dashboard-Technologie `deferred`) bleibt gültig und wird nicht dupliziert.

## 29. Compatibility

Additiv zu [UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md](UX_INFORMATION_ARCHITECTURE_AND_NAVIGATION_MODEL.md), [SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) (DEC-S-42), [OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md), [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md), [DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md](DEGRADED_MODE_AND_CAPABILITY_RESTRICTION_MODEL.md) und [FOUNDATION_CAPABILITY_MATRIX.md](FOUNDATION_CAPABILITY_MATRIX.md). Konkretisiert [DEC-P-06](../../project-system/DECISION_INDEX.md) darstellungsseitig. Keine bestehende Invariante geschwächt; kein Parallelmodell; keine neue Autorität. Adressiert RISK-12 und RISK-73 darstellungsseitig, ohne sie zu duplizieren.

## 30. Open Questions

- Konkrete Schwellen für „stale" je Datenklasse (`deferred`; keine numerische Vorgabe hier).
- Consumer-gebundene Dependency-Contract-Health als spätere eng begrenzte Erweiterung (§11) — nicht in diesem WP.
- Umfang der Abdeckungsanzeige bei sehr großen Grundgesamtheiten (Kosten-/Lesbarkeitsabwägung, `deferred`).
- Zulässige Aggregationsstufen über Workspace-Grenzen hinweg unter Isolationsauflagen (`deferred`).

## 31. Next Decision

Verbindlicher nächster Schritt gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): Nova Review von `CO-WP-027`, danach Human-Maintainer-Commit. Der CDS-Reife-Re-Check bleibt vor jeder substanziellen Designübernahme erforderlich und wird durch dieses Dokument **nicht** erfüllt und **nicht** aktiviert.
