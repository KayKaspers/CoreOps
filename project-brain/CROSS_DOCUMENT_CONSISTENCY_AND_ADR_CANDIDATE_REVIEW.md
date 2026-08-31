# CoreOps – Cross-Document Consistency and ADR Candidate Review (CO-WP-029)

**Review Status:** performed under `CO-WP-029`
**Implementation Status:** Not applicable / no implementation performed
**Technology Selection:** None
**ADR Acceptance:** None
**Foundation Readiness:** Not assessed by this WP
**Release Readiness:** Not assessed
**Normative Framework:** NDF `v1.0.0` / `9dcadc1` (`main` informativ, **nicht** normativ)

**Work Package Type:** `review-only`
**Repository Baseline:** `b7827b89f76aba61fb255cfbf5a6682d4191cefe` (`b7827b8`)
**Nova Review:** `GO WITH NOTES`
**Nova Notes 1–3:** geschlossen (§25)
**WP Status:** `completed-go-with-notes` — Human-Maintainer-Commit ausstehend

> Dieses Dokument ist ein **Review**. Es ist keine Architekturänderung, keine Technologieauswahl, keine ADR-Annahme, kein Foundation-Readiness-Urteil und keine Release-Vorbereitung. Es erteilt **keine** Freigabe für `CO-WP-030` oder `CO-WP-031`.

---

## 1. Scope and Methodology

### 1.1 Auftrag

Foundation-weite **Cross-Document Consistency Review** und **ADR Candidate Review** über den Stand bis einschließlich `CO-WP-028`. Zu beantworten waren fünfzehn Leitfragen zu Widersprüchen, Autoritätsgrenzen, Decision-/Risk-Referenzen, Statusvokabular, Current-State-Aktualität, Zähl-Claims, Parallelmodellen, offenen ADR-Kandidaten, Konsolidierung, Deferrals und der Trennung zwischen „jetzt korrigieren" und „an `CO-WP-030` übergeben".

### 1.2 Methodik

Router → Inventar → Grep → Auswahl → Lesen → Querbezug → Klassifikation → minimale Korrektur → Bericht.

1. **Preflight** (read-only): Branch, HEAD, `origin/main`, Working Tree, Index, aktive Repository-Operationen.
2. **Router-Pass**: `WORK_PACKAGE_QUEUE.md`, `NEXT_PHASE.md`, `PROJECT_PROFILE.md`, `PROJECT_BRAIN.md`, `CONTEXT_PACK_FOUNDATION_0_1.md`, `DECISION_INDEX.md`, `RISK_REGISTER.md` sowie die drei Milestone Reviews.
3. **Mechanische Verifikation** statt Prosa-Vertrauen: alle Zähl-, Tally- und ID-Aussagen wurden aus den Tabellen selbst neu berechnet (ID-Extraktion, Lücken-/Duplikatprüfung, Spalten-Tally, Likelihood × Impact-Matrix, Link-Auflösung gegen den Dateibaum).
4. **Gezielte Tiefenlesung** nur dort, wo eine Grep-/Tally-Auffälligkeit oder eine Cross-Reference-Frage dies erforderte.
5. **Klassifikation** jedes Befundes nach `CURRENT-STATE DEFECT` / `HISTORICAL RECORD` / `AMBIGUOUS` und nach Schweregrad `BLOCKER` / `MAJOR` / `MINOR` / `NOTE`.
6. **Korrektur ausschließlich** bei eindeutiger, mechanischer Evidenz; bei Uneindeutigkeit **Befund statt Edit** (fail-closed).

### 1.3 Verwendete Skills

`ndf-work-package-runner` · `ndf-adr-governance-review` · `ndf-validation-evidence-reviewer` · `ndf-public-neutrality-guard` · `ndf-compact-context-summary-runner`. Skills sind beratend und erteilen keine Autorität.

### 1.4 Ausdrücklich nicht durchgeführt

Keine Architektur-Neugestaltung · keine Technologieauswahl · keine ADR-Annahme/-Ablehnung · keine ADR-Datei · kein Readiness-Urteil · keine Release-Vorbereitung · keine Testimplementierung/-ausführung · kein Integration Lab · keine Netzwerk-/Scan-/Zielverbindung · keine Git-Schreibaktion · kein NDF-Transfer · keine CDS-Adoption.

---

## 2. Repository Baseline

| Prüfung | Erwartet | Tatsächlich | Ergebnis |
|---|---|---|---|
| Branch | `main` | `main` | PASS |
| HEAD | `b7827b89f76aba61fb255cfbf5a6682d4191cefe` | identisch | PASS |
| `origin/main` | `b7827b89f76aba61fb255cfbf5a6682d4191cefe` | identisch | PASS |
| HEAD == `origin/main` | ja | ja | PASS |
| Working Tree | clean | clean | PASS |
| Index | clean | clean | PASS |
| merge / rebase / cherry-pick | keine | keine | PASS |
| Git-Schreibaktionen | keine | keine | PASS |

Letzter Commit: `b7827b8 docs(testing): establish foundation test strategy and integration lab governance`.

---

## 3. Sources and Inventory Reviewed

### 3.1 Dokumentbestand

| Bereich | Dokumente |
|---|---|
| `docs/architecture/` | 39 |
| `docs/security/` | 33 |
| `docs/governance/` | 14 |
| `project-system/` | 8 |
| `project-brain/` | 5 |
| `docs/testing/` | 3 |
| `docs/integrations/` | 1 |
| `docs/project-system/` | 1 |
| Repository-Wurzel (`ROADMAP.md`) | 1 |
| **Foundation gesamt (geprüfter Bestand zur Baseline `b7827b8`)** | **105** (26.245 Zeilen) |
| `.claude/skills/` (importierter NDF-Pack, nicht Foundation-Inhalt) | 39 |

### 3.2 Register und autoritative Quellen

| Quelle | Rolle | Umfang |
|---|---|---|
| `project-system/DECISION_INDEX.md` | autoritativ für Decisions | 373 `DEC-S` + 8 `DEC-P` + 8 `DEC-G` + 21 `DEC-O` + 3 `DEC-A`-Zeilen (32 Kandidaten) + 6 `DEC-D` + 12 `DEC-N` |
| `project-system/RISK_REGISTER.md` | autoritativ für Risiken | 316 |
| `docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md` | autoritativ für Capabilities | 94 über 13 Domains |
| `docs/architecture/CONCEPT_DECISION_CLASSIFICATION.md` | autoritativ für ADR-Kandidatenliste | 32 |
| `docs/security/THREAT_SCENARIO_REGISTER.md` | autoritativ für Threat Scenarios | 40 (`THR-001…040`) |
| `docs/architecture/COREOPS_MODULE_CATALOG.md` | autoritativ für logische Module | 17 (`MOD-*`) |
| `project-system/LESSONS_LEARNED_REGISTER.md` | autoritativ für Lessons | 38 (read-only geprüft) |
| `project-system/NDF_FEEDBACK_CANDIDATES.md` | autoritativ für NDF-Kandidaten | 15 (read-only geprüft) |

---

## 4. Consistency Verdict

**Die Foundation ist über alle 105 Dokumente inhaltlich konsistent.**

- **0** materielle Foundation-Contradictions.
- **0** Parallelmodelle.
- **0** konkurrierende Autoritätsmodelle.
- **0** verdeckte Autoritäts-, Approval- oder Execution-Mechanismen.
- **0** BLOCKER.

Sämtliche Befunde sind **referenzieller, zählerischer oder Current-State-Spiegel-Natur** oder betreffen die **Struktur des Decision Index**, nicht den Inhalt der Foundation-Modelle. Die inhaltliche Kette

> Concept → Scope/Capability → Sprache/Repository-Governance → Systemkontext/Planes → Threat Model → Module → Human Identity → Machine Identity → Source of Truth → State/Drift → Policy/Approval/Execution → Integration Contract → Domain Packs → Data/Migration → API → Event/Audit/Evidence → Telemetry → Topology → Deployment → Artifact Trust → Restricted/Offline/CorePack → Secrets/Key Custody → Data Classification/Retention → Self-Protection/Degraded/Recovery → UX/Dashboard → Test/Fixture/Lab

hält durchgängig: jede Stufe konsumiert die vorige als **autoritative Grenze** und **referenziert statt dupliziert**.

---

## 5. Finding Summary

Jeder Befund trägt **zwei unabhängige Dimensionen**: eine **Severity** (wie schwer) und eine **Disposition** (wie es weitergeht).

```
finding severity          !=  finding lifecycle status
MAJOR                     !=  automatischer READINESS BLOCKER
an CO-WP-030 uebergeben   !=  stillschweigend akzeptiert
```

**Severity-Tally**

| Severity | Anzahl | korrigiert | weitergereicht |
|---|---|---|---|
| **BLOCKER** | 0 | — | — |
| **MAJOR** | 7 | 2 | 5 |
| **MINOR** | 12 | 5 | 7 |
| **NOTE** | 6 | 0 | 6 |
| **Gesamt** | **25** | **7** | **18** |

**Dispositions-Tally**

| Disposition | Bedeutung | Anzahl |
|---|---|---|
| `corrected` | in diesem WP mechanisch behoben; keine Restaktion | **7** |
| `open-human-decision` | offen; nächste Entscheidung liegt beim **Human Maintainer** (`HM-1`…`HM-14`, §22.2) | **8** |
| `open-readiness-review` | offen; nächste Bewertung liegt bei **`CO-WP-030`** | **5** |
| `deferred-post-foundation` | offen; ausserhalb Foundation 0.1, spätere Design-WPs | **1** |
| `note-only` | bewertet, keine Aktion erforderlich; als Beobachtung festgehalten | **4** |
| **Gesamt** | | **25** |

> Diese Severity-Werte sind **Review-Finding-Severity**. Sie sind **keine** Risk-Register-Severity und erzeugen **kein** Risiko. **Kein** Befund dieses Reviews ist ein Readiness-Blocker (§22.1); ob ein weitergereichter Befund ein positives Readiness-Verdikt beeinflusst, entscheidet **`CO-WP-030`**, nicht dieses WP.

### 5.1 MAJOR-Befunde

| ID | Befund | Disposition |
|---|---|---|
| **MAJ-1** | `DEC-S-88` wird in drei Prosa-Referenzen des Decision Index als **Break-Glass**-Entscheidung geführt. Autoritativ ist `DEC-S-88` = *Session technology* (`architecture-context` · `deferred` · `non-binding` · `CO-WP-009`). Die Fehlreferenz betrifft sicherheitsrelevante Governance-Prosa und hat fünf Work Packages überdauert. | `corrected` (§9) |
| **MAJ-2** | Die 24 Tabellen `DEC-S-38…373` (**336** Zeilen) besitzen **keine** `ADR Required`-Spalte. Der ADR-Bedarf ist für 90 % aller `DEC-S`-Einträge aus dem Schema **nicht ableitbar** (weder bejahend noch verneinend); die einzige Absicherung sind die Prosa-Aussagen „keine ADR" in den WP-Berichten. | `open-human-decision` → HM-11 |
| **MAJ-3** | Vier ADR-pflichtige Decisions tragen `ADR Required = ja`, aber **keine** ADR-Kandidaten-ID: `DEC-O-05` (Offline-Policy bei getrennter Control Plane), `DEC-O-07` (privilegierte Ausführung vs. „keine Remote-Root-Shell"), `DEC-D-03` (externe Secret-Backends), `DEC-D-06` (bidirektionaler NetBox-Sync). Sie sind im ADR-Kandidateninventar **nicht** repräsentiert. | `open-human-decision` → HM-6 |
| **MAJ-4** | Zwei Decision-Konventionen koexistieren weiterhin: `DEC-S-01…37` (37 Zeilen, **kombinierter** Pseudostatus in einer Spalte) gegenüber `DEC-S-38…373` (336 Zeilen, **getrennte** Dimensionen Decision Class · Lifecycle Status · Binding Level). Aus drei Milestone Reviews als `CO-WP-029`-Follow-up vorgemerkt. | `open-human-decision` → HM-10 (§8.4) |
| **MAJ-5** | `DEC-S-52…317` enthalten **~30** reine „Technologie deferred"-Decisions. Ab `CO-WP-024` gilt die von Nova bestätigte gegenteilige Regel: *„eine Decision allein dafür, dass **keine** Technologie ausgewählt wurde, ist unzulässig"* — `CO-WP-024…028` haben entsprechend **keine** Deferral-Decision registriert. Der Altbestand widerspricht der heute geltenden Regel. | `open-human-decision` → HM-12 (§8.5) |
| **MAJ-6** | Das Risk Register veröffentlicht **keine** Likelihood × Impact → Risk-Level-Matrix. Gegen die aus dem Register selbst abgeleitete Modal-Abbildung weichen **24** Zeilen ab. Mangels etablierter Kalibrierungsregel wurde **keine** Severity geändert. | `open-human-decision` → HM-7 (§15) |
| **MAJ-7** | Die **autoritative** Capability Matrix nannte in ihrer eigenen Zusammenfassung „**Domains: 12**", zählte in derselben Aufzählung jedoch **13** und trägt **13** eindeutige Domänen-Präfixe in den Capability-IDs. Selbstwidersprüchliche autoritative Metadaten. | `corrected` (§16) |

### 5.2 MINOR-Befunde

| ID | Befund | Klassifikation | Disposition |
|---|---|---|---|
| **MIN-1** | Capability-Count „74" statt autoritativ **94** in vier Current-State-Spiegeln und in der `DEC-O-17`-Notiz. | current-state defect | `corrected` |
| **MIN-2** | „Human-Maintainer-Commit ausstehend" für `CO-WP-028` in fünf Current-State-Spiegeln, obwohl `CO-WP-028` als `b7827b8` committet und gepusht ist. | current-state defect | `corrected` |
| **MIN-3** | `1dee29d` als **aktuelle** Repository-Baseline in Current-State-Routern; aktuell ist `b7827b8`. | current-state defect | `corrected` |
| **MIN-4** | `NEXT_PHASE.md` §Secrets sagte im selben Abschnitt zuerst „`CO-WP-024` Human-Maintainer-Commit ausstehend" und wenige Zeilen später „`CO-WP-024` nach Human-Maintainer-Commit (`916ba66`) und Push". | current-state defect (Selbstwiderspruch) | `corrected` |
| **MIN-5** | Zwei Verweise in `DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md` zeigten auf `../architecture/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md`; die Datei liegt in `docs/security/`. Vier andere Repository-Verweise auf dieselbe Datei sind korrekt. | broken internal reference | `corrected` |
| **MIN-6** | `OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md` §35 verweist auf `../architecture/SOURCE_OF_TRUTH_AND_FIELD_PROVENANCE_MODEL.md` — diese Datei existiert **nicht**. Der Dateiname vermengt `SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md` und `FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md`; der Linktext lautet „CO-WP-011", und die Aussage betrifft *Offline-Grenzen*, wofür `OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md` das sachlich passende Ziel wäre. **Zwei vertretbare Zielkandidaten → fail-closed, nicht korrigiert.** | AMBIGUOUS | `open-human-decision` → HM-14 |
| **MIN-7** | `PROJECT_PROFILE.md` §16 nennt als „Nächster Meilenstein" `CO-WP-002`. Die Sektionen 1–17 sind der originale, mit „Erzeugt durch `CO-WP-001`" gestempelte Profiltext; über 40 nachgelagerte Konsolidierungsabschnitte überschreiben ihn faktisch. | AMBIGUOUS (historisch gerahmt, aber current-state formuliert) | `open-readiness-review` → `CO-WP-030` (§20) |
| **MIN-8** | `PROJECT_PROFILE.md` §9 „Es existiert ausschließlich das Core Governance Skeleton" — nach 105 Foundation-Dokumenten unzutreffend, aber in derselben `CO-WP-001`-Sektion. | AMBIGUOUS | `open-readiness-review` → `CO-WP-030` (§20) |
| **MIN-9** | `PROJECT_BRAIN.md` §Bekannte Risiken verweist auf „`RISK_REGISTER.md`, 39 Foundation-Risiken". Gemeint sind die initialen Concept-Risiken `RISK-01…39`; ein Leser liest es als Registerumfang (tatsächlich 316). | AMBIGUOUS | `open-readiness-review` → `CO-WP-030` (§20) |
| **MIN-10** | `PROJECT_BRAIN.md` §Offene Fragen führt „Capability Matrix und Support Boundary (→ `CO-WP-004`)" als offen; `CO-WP-004` ist abgeschlossen und committet. | AMBIGUOUS (`CO-WP-002`-Sektion) | `open-readiness-review` → `CO-WP-030` (§20) |
| **MIN-11** | `RISK-66` („Capability-Zählungsinkonsistenz „74" vs. 94") trägt Target WP `later consistency WP`. Genau dieses WP ist `CO-WP-029`, und die geforderte Behandlung ist mit diesem Review **materiell ausgeführt**. Status/Target wurden **nicht** geändert (kein mechanischer Retarget, keine Schließung ohne HM). | current-state ambiguity | `open-human-decision` → HM-8 (§14.4) |
| **MIN-12** | `ROADMAP.md` behauptet „**ITIL-/PRINCE2-Tailoring** bleibt offen (Kandidaten; Entscheidung in `CO-WP-004D`)". `CO-WP-004D` ist abgeschlossen und hat entschieden (`DEC-S-23…30`: ITIL `adopted-with-tailoring`, PRINCE2 V7 `optional-profile`, Vollimplementierung beider `rejected`). | current-state defect | `open-human-decision` → HM-13 (`ROADMAP.md` außerhalb des Änderungsrahmens) |

### 5.3 NOTE-Befunde

| ID | Beobachtung | Disposition |
|---|---|---|
| **NOTE-1** | `.claude/skills/README.md` enthält fünf Verweise auf NDF-Repository-Pfade (`docs/adr/ADR-0032-…`, `docs/agent-workflows/…`, `docs/validation/foundation-0-9/…`), die in CoreOps nicht existieren. Das ist **korrekt so**: der Skills-Pack wurde byte-identisch aus NDF `v1.0.0`/`9dcadc1` übernommen; eine lokale Anpassung würde die Provenance-Bindung brechen. **Nicht** zu korrigieren. | `note-only` |
| **NOTE-2** | **Dokumentationsökonomie:** vier Current-State-Spiegel führen je eine nahezu vollständige Pro-WP-Narrative — `NEXT_PHASE.md`, `PROJECT_PROFILE.md`, `PROJECT_BRAIN.md`, `CONTEXT_PACK_FOUNDATION_0_1.md`. Ein Foundation-Sachverhalt wird typisch **vierfach** in unterschiedlicher Tiefe wiederholt. Genau diese Redundanz hat MIN-1…MIN-4 verursacht: ein Faktum ändert sich, vier Stellen müssen nachgezogen werden. | `open-readiness-review` → `CO-WP-030` |
| **NOTE-3** | Das Risk Register enthält zusätzlich zur Tabelle 33 Pro-WP-Abschnitte „`CO-WP-xxx` – geänderte Risiken", die den Tabelleninhalt in Prosa wiederholen. Konsolidierungskandidat. | `note-only` |
| **NOTE-4** | Die ADR-Kandidatenliste `ADR-0001…0030` steht wortgleich in `COREOPS_CONCEPT_V3.md` §51 **und** in `CONCEPT_DECISION_CLASSIFICATION.md`. Da das Concept unveränderlich registriert ist, ist die Duplikation vertretbar; die Klassifikationsdatei bleibt autoritativ. | `note-only` |
| **NOTE-5** | Die CDS-Grenze ist über drei Dokumente **wortgleich und korrekt** gesetzt (`Adoption: Not started`, `Pilot: Inactive / not activated`, `CDS-WP-017` inaktiv, keine Tokens, keine Consumer-Evidenz). Kein Defekt — als bestätigte Grenze festgehalten (§21). | `note-only` |
| **NOTE-6** | Die vier incident-abgeleiteten echten Erweiterungskandidaten sind durch `CO-WP-027`/`CO-WP-028` **weder umgesetzt noch entwertet**; ihre Klassifikation aus dem `CO-WP-021…026` Milestone Review gilt unverändert (§19). | `deferred-post-foundation` |


### 5.4 Finding Lifecycle Ledger

Vollständige, deterministische Zuordnung aller 25 Befunde. Jede Zeile beantwortet: **korrigiert?** · **noch offen?** · **wer entscheidet als Nächstes?** · **kann es ein positives Readiness-Verdikt beeinflussen?**

| ID | Severity | Disposition | Nächste Entscheidung bei | Readiness-Wirkung |
|---|---|---|---|---|
| MAJ-1 | MAJOR | `corrected` | — (erledigt) | keine |
| MAJ-2 | MAJOR | `open-human-decision` | Human Maintainer (HM-11), Bewertung `CO-WP-030` | möglich — `CO-WP-030` klassifiziert |
| MAJ-3 | MAJOR | `open-human-decision` | Human Maintainer (HM-6) | möglich — `CO-WP-030` klassifiziert |
| MAJ-4 | MAJOR | `open-human-decision` | Human Maintainer (HM-10) | möglich — `CO-WP-030` klassifiziert |
| MAJ-5 | MAJOR | `open-human-decision` | Human Maintainer (HM-12) | möglich — `CO-WP-030` klassifiziert |
| MAJ-6 | MAJOR | `open-human-decision` | Human Maintainer (HM-7) | möglich — `CO-WP-030` klassifiziert |
| MAJ-7 | MAJOR | `corrected` | — (erledigt) | keine |
| MIN-1 | MINOR | `corrected` | — (erledigt) | keine |
| MIN-2 | MINOR | `corrected` | — (erledigt) | keine |
| MIN-3 | MINOR | `corrected` | — (erledigt) | keine |
| MIN-4 | MINOR | `corrected` | — (erledigt) | keine |
| MIN-5 | MINOR | `corrected` | — (erledigt) | keine |
| MIN-6 | MINOR | `open-human-decision` | Human Maintainer (HM-14) | keine — redaktioneller Verweis |
| MIN-7 | MINOR | `open-readiness-review` | `CO-WP-030` (§20) | keine — Dokumentationsökonomie |
| MIN-8 | MINOR | `open-readiness-review` | `CO-WP-030` (§20) | keine — Dokumentationsökonomie |
| MIN-9 | MINOR | `open-readiness-review` | `CO-WP-030` (§20) | keine — Dokumentationsökonomie |
| MIN-10 | MINOR | `open-readiness-review` | `CO-WP-030` (§20) | keine — Dokumentationsökonomie |
| MIN-11 | MINOR | `open-human-decision` | Human Maintainer (HM-8) | keine — Register-Housekeeping |
| MIN-12 | MINOR | `open-human-decision` | Human Maintainer (HM-13) | keine — Datei ausserhalb des Rahmens |
| NOTE-1 | NOTE | `note-only` | — (bewusst unverändert) | keine |
| NOTE-2 | NOTE | `open-readiness-review` | `CO-WP-030` (§22.3 Note 4) | keine — Wartbarkeit |
| NOTE-3 | NOTE | `note-only` | — (geprüft: `keep`) | keine |
| NOTE-4 | NOTE | `note-only` | — (geprüft: `keep`) | keine |
| NOTE-5 | NOTE | `note-only` | — (bestätigte Grenze) | keine |
| NOTE-6 | NOTE | `deferred-post-foundation` | spätere Design-WPs | keine |

**Abstimmung mit dem Severity-Tally:** korrigiert **7** (MAJ-1, MAJ-7, MIN-1…MIN-5) · weitergereicht **18** (`open-human-decision` 8 + `open-readiness-review` 5 + `deferred-post-foundation` 1 + `note-only` 4). Severity 0/7/12/6 = **25**; Disposition 7/8/5/1/4 = **25**. **Beide Summen stimmen überein; keine Severity wurde zur Anpassung der Zahlen verändert.**

**Eigentümerschaft der weitergereichten Befunde:** 8 liegen beim **Human Maintainer** (`HM-6`, `HM-7`, `HM-8`, `HM-10`, `HM-11`, `HM-12`, `HM-13`, `HM-14`), 5 bei **`CO-WP-030`**, 1 bei späteren **Design-WPs**, 4 sind abschliessend bewertete Beobachtungen ohne Restaktion. **Kein Befund ist stillschweigend akzeptiert.**

---

## 6. Corrected Mechanical Defects

**Sieben** Korrekturen über **acht** bestehende Pfade; jede unterhalb der Schwelle „inhaltliche Änderung".

| # | Datei | Stelle | Vorher | Nachher | Klasse |
|---|---|---|---|---|---|
| C-1 | `project-system/DECISION_INDEX.md` | §Secrets… (`CO-WP-024`-Deduplizierungsnotiz) | `Break-Glass (DEC-S-88)` | `Break-Glass (DEC-S-84/DEC-S-85/DEC-S-147)` | falsche Decision-ID-Referenz |
| C-2 | `project-system/DECISION_INDEX.md` | §Self-Protection… (`CO-WP-026`-Registrierungsnotiz) | `… DEC-S-329 (Recovery ≠ Reinstatement), DEC-S-88 (Break-Glass).` | `… DEC-S-329 (Recovery ≠ Reinstatement), DEC-S-84/DEC-S-85/DEC-S-147 (Break-Glass).` | falsche Decision-ID-Referenz |
| C-3 | `project-system/DECISION_INDEX.md` | §Zusammenfassung, `CO-WP-026`-Registrierungen | `… und DEC-S-88 (Break-Glass) referenziert` | `… und DEC-S-84/DEC-S-85/DEC-S-147 (Break-Glass) referenziert` | falsche Decision-ID-Referenz |
| C-4 | `project-system/DECISION_INDEX.md` | Zeile `DEC-O-17`, Spalte *Notes* | `74 Capabilities` | `94 Capabilities … (Zählkorrektur CO-WP-004E …)` | Decision-Notiz ohne Quellendeckung |
| C-5 | `docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md` | §Zusammenfassung | `**Domains:** 12` | `**Domains:** 13` + Korrekturnotiz | selbstwidersprüchliche autoritative Metadaten |
| C-6 | `docs/security/DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md` | §Manual Authority und §Compatibility (2×) | `../architecture/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md` | `TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md` | fehlgeleiteter interner Verweis |
| C-7 | `NEXT_PHASE.md`, `PROJECT_PROFILE.md`, `PROJECT_BRAIN.md`, `CONTEXT_PACK_FOUNDATION_0_1.md`, `WORK_PACKAGE_QUEUE.md` | Current-State-Abschnitte | `CO-WP-028` „Commit ausstehend", Baseline `1dee29d`, Capability-Count `74`/`12 Domains`, `CO-WP-024` „Commit ausstehend" | Commit `b7827b8` gepusht; Baseline `b7827b8`; `94`/`13 Domains`; `CO-WP-024` Commit `916ba66` | veraltete Current-State-Spiegel |

**Nicht** korrigiert (bewusst): MIN-6 (zwei vertretbare Zielkandidaten), MIN-7…MIN-10 (historisch gerahmte Sektionen), MIN-11 (`RISK-66`-Status ist HM-Entscheidung), MIN-12 (`ROADMAP.md` ist forbidden file), NOTE-1 (Provenance-gebundener Import).

---

## 7. Unresolved Contradictions

**Keine.** Es wurde **kein** inhaltlicher Widerspruch zwischen zwei Foundation-Modellen gefunden, der ungelöst bliebe.

Die verbleibenden offenen Punkte sind ausnahmslos:

- **strukturell** (MAJ-2…MAJ-5: Decision-Index-Schema und Alt-Konventionen),
- **kalibrierend** (MAJ-6: fehlende, aber nachvollziehbar herleitbare Severity-Regel),
- **redaktionell/uneindeutig** (MIN-6…MIN-12),
- **oder ausdrücklich als offen registriert** (die sechs Konflikte `CCR-01`, `CCR-05`…`CCR-09`, die seit `CO-WP-002` sichtbar und mit Ziel-WPs verknüpft geführt werden — kein Konsistenzdefekt, sondern korrekt dokumentierte Offenheit).

---

## 8. Decision Index Review

### 8.1 Strukturelle Integrität

| Prüfung | Ergebnis |
|---|---|
| `DEC-S`-IDs eindeutig | **PASS** — 373 eindeutige IDs, **0** Duplikate |
| `DEC-S`-Nummernkreis lückenlos | **PASS** — `DEC-S-1` … `DEC-S-373`, **keine** Lücke |
| `DEC-P` / `DEC-G` / `DEC-O` / `DEC-A` / `DEC-D` / `DEC-N` eindeutig | **PASS** |
| Zusammenfassungs-Counts vs. Tabellen | **PASS** — Accepted Product 8 ✓ · Binding Governance 8 ✓ · Foundation Decisions 21 ✓ · ADR Candidates 32 ✓ · Deferred 6 ✓ · Non-Goals 12 ✓ |
| Verwaiste Referenzen auf nicht existierende Decision-IDs | **keine gefunden** |
| Falsche Decision-ID-Referenzen | **3** (alle `DEC-S-88`, alle korrigiert — §9) |

### 8.2 Statusvokabular

| Dimension | Werte | Verteilung (`DEC-S-38…373`) |
|---|---|---|
| Lifecycle Status | `accepted` · `deferred` · `not-claimed` | 299 · 36 · 1 |
| Decision Class | `security-context` · `governance-direction` · `architecture-context` · `product-direction` | 198 · 81 · 56 · 1 |
| Binding Level | `binding-governance` · `non-binding` · `guidance` | 296 · 37 · 3 |

Das Vokabular ist innerhalb der neuen Konvention **konsistent**; es existieren **keine** kombinierten Pseudostatuswerte und **keine** unzulässigen Werte.

### 8.3 Kritische Prüfung: „technische Entscheidung als accepted"

**PASS.** Keine `architecture-context`-Decision trägt `accepted` in einem Sinn, der eine Technologie auswählen würde. Alle 36 `deferred`-Einträge sind Technologie-Offenhaltungen. Die Grundregel des Index — *„Keine technische Entscheidung erhält `accepted`"* — ist eingehalten.

### 8.4 Zwei-Konventionen-Koexistenz (MAJ-4)

| Bereich | Zeilen | Schema |
|---|---|---|
| `DEC-S-01…37` | 37 | `ID · Topic · Status · Class · Source · Owner · Target WP · ADR Required · Notes` — Status trägt **kombinierte** Werte (`accepted-product-direction`, `proposed-binding-governance`, `controlled-candidate-process`, `foundation-candidate`, `not-current-target`, …) |
| `DEC-S-38…373` | 336 | `ID · Topic · Decision Class · Lifecycle Status · Binding Level · Source · Owner · Notes` — **getrennte** Dimensionen, **keine** `Target WP`- und **keine** `ADR Required`-Spalte |

Aus `MILESTONE_REVIEW_CO_WP_005_TO_012` (Follow-up 3), `…_013_TO_020` (Follow-up 2) und `…_021_TO_026` (Follow-up 2) jeweils für `CO-WP-029` vorgemerkt.

**Bewertung:** Eine Migration von `DEC-S-01…37` auf das neue Schema wäre eine **Massen-Statusmigration** — durch §10/§26 dieses WP ausdrücklich untersagt und inhaltlich eine Entscheidung, keine Reconciliation. **Empfehlung:** Human-Maintainer-Entscheidung in `CO-WP-030` zwischen (a) formaler Migration in einem eigenen WP, (b) dokumentierter Legacy-Deklaration von `DEC-S-01…37` mit Mapping-Tabelle, oder (c) bewusster Beibehaltung. **Keine Änderung in diesem WP.**

### 8.5 Technology-Deferral-Regelbruch (MAJ-5)

`DEC-S-52 · 53 · 60 · 61 · 62 · 63 · 73 · 74 · 75 · 87 · 88 · 89 · 103 · 104 · 105 · 119 · 120 · 121 · 134 · 135 · 136 · 151 · 152 · 166 · 167 · 168 · 184 · 185 · 203 · 219 · 236 · 254 · 273 · 288 · 303 · 317` sind `deferred`-Einträge, überwiegend reine Technologie-Offenhaltungen (`CO-WP-006`…`CO-WP-023`).

Ab `CO-WP-024` gilt die entgegengesetzte, durch Nova-Note bestätigte Regel; `CO-WP-024…028` haben konsequent **keine** Deferral-Decision registriert und die Deferral stattdessen in den *Technology-Boundary-/Non-Goals*-Abschnitten der Dokumente verankert.

**Bewertung:** Kein inhaltlicher Widerspruch (die Aussagen sind sachlich korrekt), aber eine **Regelinkonsistenz im Register**. Konsolidierung — etwa zu einer indexseitigen Sammel-Deferral-Übersicht — ist ein Kandidat, keine Reconciliation. **Keine Änderung in diesem WP.**

### 8.6 Statushärtung — geprüft, nicht durchgeführt

Es wurde **keine** Decision-Status-Änderung vorgenommen. Die zehn `proposed` / `proposed-binding-governance`-Einträge aus `CO-WP-003`/`CO-WP-004` (`DEC-O-02`, `-03`, `-04`, `-10`, `-11`, `-12`, `-16`, `-17`, `-18`, `-19`, `-20`, `-21`) sind ausdrücklich als *„vor dem Human-Maintainer-Commit unverbindlich"* definiert. Da die betreffenden Commits inzwischen vorliegen, ist ihre Verbindlichkeit **argumentierbar** — die Umstellung wäre jedoch eine **Human-Maintainer-Entscheidung**, keine mechanische Korrektur. → `CO-WP-030` (§22, HM-4).

---

## 9. Known DEC-S-88 Correction

### 9.1 Autoritative Quelle

```
| DEC-S-88 | Session technology | architecture-context | deferred | non-binding | CO-WP-009 | Nova | Deferred |
```
(`project-system/DECISION_INDEX.md`, §Human Identity, RBAC and Break-Glass Decisions)

Die tatsächlichen Break-Glass-Entscheidungen:

```
| DEC-S-84  | Break glass              | security-context | accepted | binding-governance | CO-WP-009 | HM   | Temporär, benannt, reason-/scope-bound, auditiert |
| DEC-S-85  | Break-glass permissions  | security-context | accepted | binding-governance | CO-WP-009 | Nova | Müssen ablaufen oder widerrufen werden |
| DEC-S-147 | Break-glass relationship | security-context | accepted | binding-governance | CO-WP-013 | HM   | Außergewöhnlich, nicht permanent; referenziert bestehende Break-Glass-Policy |
```

Die korrekte Schreibweise war im Repository **bereits etabliert**: `CO-WP-027` verwendet an zwei Stellen exakt `DEC-S-84/DEC-S-85/DEC-S-147 (Break-Glass)`, ebenso `UX_ACTION_SAFETY_ACCESSIBILITY_AND_DISCLOSURE_POLICY.md`.

### 9.2 Klassifikation aller `DEC-S-88`-Vorkommen

| # | Pfad | Stelle | Kontext | Klasse | Aktion |
|---|---|---|---|---|---|
| 1 | `project-system/DECISION_INDEX.md` | Tabellenzeile, §`CO-WP-009` | `DEC-S-88 \| Session technology \| architecture-context \| deferred \| non-binding` | **A — korrekte Session-technology-Referenz** | **unverändert** |
| 2 | `project-system/DECISION_INDEX.md` | `CO-WP-024`-Deduplizierungsnotiz | „Machine-Credential-Lifecycle (DEC-S-90…105), **Break-Glass (DEC-S-88)** und Offline-Autorität (DEC-S-305/314)…" | **B — falsche Break-Glass-Referenz** | **korrigiert → `DEC-S-84/DEC-S-85/DEC-S-147`** |
| 3 | `project-system/DECISION_INDEX.md` | `CO-WP-026`-Registrierungsnotiz (§Self-Protection) | „…DEC-S-329 (Recovery ≠ Reinstatement), **DEC-S-88 (Break-Glass)**." | **B — falsche Break-Glass-Referenz** | **korrigiert → `DEC-S-84/DEC-S-85/DEC-S-147`** |
| 4 | `project-system/DECISION_INDEX.md` | §Zusammenfassung, `CO-WP-026`-Registrierungen | „…DEC-S-329 (Recovery ≠ Reinstatement) und **DEC-S-88 (Break-Glass)** referenziert, nicht dupliziert." | **B — falsche Break-Glass-Referenz** | **korrigiert → `DEC-S-84/DEC-S-85/DEC-S-147`** |

**Kategorie C (ambiguous): 0.**

Gegenprobe: `git grep -n "DEC-S-88"` liefert nach der Korrektur **genau ein** Vorkommen — die autoritative Tabellenzeile.

### 9.3 Ergebnis

| Metrik | Wert |
|---|---|
| `DEC-S-88`-Vorkommen gesamt (vorher) | 4 |
| Kategorie A (korrekt, unverändert) | **1** |
| Kategorie B (falsch, korrigiert) | **3** |
| Kategorie C (ambiguous) | **0** |
| Decisions umbenannt / umnummeriert / im Status geändert | **0** |
| Neue Decisions | **0** |

`DEC-S-88` behält seine autoritative Bedeutung *Session technology* **unverändert**. Es handelt sich um eine rein **referenzielle** Korrektur. Ein Korrekturhinweis wurde am Ende des Decision Index verankert.

---

## 10. ADR Candidate Inventory

### 10.1 Bestandsermittlung (nicht angenommen, neu berechnet)

| Quelle | Aussage |
|---|---|
| `CONCEPT_DECISION_CLASSIFICATION.md` §ADR-Required Topics | 32 nummerierte Kandidaten; *„Quelle: Concept §51 (30 Kandidaten) plus zwei Foundation-Klärungen"* |
| `CONCEPT_DECISION_CLASSIFICATION.md` Klassenstatistik | `\| ADR \| 32 \|` |
| `COREOPS_CONCEPT_V3.md` §51 | `ADR-0001` … `ADR-0030` (30 Zeilen) |
| `DECISION_INDEX.md` §ADR Candidates | `DEC-A-0001…0030` (Sammelzeile) + `DEC-A-0031` + `DEC-A-0032` |
| `PROJECT_BRAIN.md` | „32 ADR-Kandidaten erfasst" |

**Bestätigt: 32.** Die Zahl wurde **nicht** aus dem Prompt übernommen, sondern aus vier unabhängigen Repository-Quellen rekonstruiert; alle stimmen überein.

> **Strukturhinweis:** Im Decision Index existieren nur **drei** `DEC-A`-Tabellenzeilen, weil `DEC-A-0001…0030` als **Sammelzeile** geführt wird. Einzeln adressierbar sind die Kandidaten nur über `CONCEPT_DECISION_CLASSIFICATION.md`. Das ist konsistent, erschwert aber die maschinelle Rückverfolgung einzelner Kandidaten — Konsolidierungshinweis für `CO-WP-030`.

### 10.2 Nicht repräsentierte ADR-Bedarfe (MAJ-3)

Vier Decisions tragen `ADR Required = ja` **ohne** ADR-Kandidaten-ID:

| Decision | Topic | Status | Ziel |
|---|---|---|---|
| `DEC-O-05` | Offline-Policy bei getrennter Control Plane (`CCR-05`) | `open` | `CO-WP-013/023` |
| `DEC-O-07` | Privilegierte Ausführung vs. „keine Remote-Root-Shell" (`CCR-07`) | `open` | `CO-WP-013` |
| `DEC-D-03` | Externe Secret-Backends | `deferred` | post-Foundation |
| `DEC-D-06` | Bidirektionaler NetBox-Sync (über read-only hinaus) | `deferred` | post-Foundation |

Zusätzlich tragen fünf Decisions `ADR Required = teils` (`DEC-G-08`, `DEC-S-03`, `DEC-S-19`, `DEC-S-21`, `DEC-S-36`) ohne Kandidatenzuordnung.

**Kein neuer Kandidat wurde angelegt.** Die Nummernvergabe für ADR-Kandidaten ist Human-Maintainer-Sache. → `CO-WP-030` (HM-6).

### 10.3 ADR-Readiness-Basis — was den ADR-Bedarf bestimmt (und was nicht)

Die Strukturbefunde **MAJ-2** und **MAJ-3** betreffen die **Ableitbarkeit** des ADR-Bedarfs aus dem Decision-Index-Schema. Sie sind ausdrücklich **keine** Aussage darüber, ob ein ADR erforderlich ist. Binding:

```
fehlende ADR-Required-Spalte an einer DEC-S-Zeile   !=  ADR nicht erforderlich
fehlende ADR-Required-Spalte an einer DEC-S-Zeile   !=  ADR erforderlich
DEC-S Lifecycle Status / Decision Class             !=  ADR-Disposition
Ableitbarkeit aus dem Schema                        !=  ADR-Bedarf
```

Eine fehlende Spalte trägt **keinen** Informationsgehalt in eine der beiden Richtungen. Aus dem Fehlen darf **weder** geschlossen werden, dass für die betreffenden 336 `DEC-S`-Einträge ein ADR nötig ist, **noch**, dass keiner nötig ist. Ebenso ist der Lifecycle Status einer `DEC-S`-Zeile (`accepted` / `deferred` / `not-claimed`) **keine** ADR-Disposition: `deferred` bedeutet Technologie offengehalten, nicht ADR offen; `accepted` bedeutet Governance-Semantik bindend, nicht ADR entschieden.

**Die readiness-relevante ADR-Basis ist:**

| # | Bestandteil | Wo | Status |
|---|---|---|---|
| 1 | **Autoritatives `DEC-A`-Kandidateninventar** — 32 Kandidaten | `DECISION_INDEX.md` §ADR Candidates + `CONCEPT_DECISION_CLASSIFICATION.md` §ADR-Required Topics | vollständig, verifiziert (§10.1) |
| 2 | **Autoritative Source-/Decision-Zuordnungen** — die `ADR Required`-Angaben aller Tabellen, die diese Spalte führen (`DEC-P`, `DEC-G`, `DEC-O`, `DEC-A`, `DEC-D`, `DEC-S-01…37`) | `DECISION_INDEX.md` | vollständig ausgewertet (§10.2) |
| 3 | **CO-WP-029 ADR-Dispositionsmatrix** — alle 32 Kandidaten mit primärer Disposition | §11 dieses Dokuments | vollständig (32/32) |
| 4 | **Explizite Human-Maintainer-Entscheidungen** | §22.2 (`HM-1`…`HM-14`) | offen, in `CO-WP-030` zu klassifizieren |

Diese vier Bestandteile sind **gemeinsam** die Grundlage jeder Readiness-Aussage zum Exit Gate der entschiedenen relevanten ADRs. Die `DEC-S-38…373`-Tabellen sind für diese Basis **nicht** erforderlich — sie registrieren Governance-Semantik, nicht ADR-Bedarf. `MAJ-2` beschreibt daher eine **Prüfbarkeits- und Komfortlücke im Schema**, keine Lücke in der Readiness-Basis.

**Ausdrücklich nicht getan:** keine `ADR Required`-Spalte an `DEC-S`-Zeilen ergänzt · keine Schema-Migration des Decision Index · keine Decision erzeugt · kein ADR erzeugt.

---

## 11. ADR Disposition Matrix

**Review-Vokabular** (keine Decision-Index-Status): `still-open` · `foundation-semantics-established-technical-choice-pending` · `candidate-for-consolidation` · `deferred-post-foundation` · `candidate-may-no-longer-be-required` · `requires-human-decision-before-readiness` · `no-readiness-blocker-currently`.

Jeder Kandidat trägt genau **eine primäre Disposition**; Konsolidierungs-Cluster sind als Sekundärmerkmal geführt.

**Für alle 32 gilt:** *Technische Wahl bereits getroffen?* = **nein**. *ADR akzeptiert?* = **nein**.

| # | ADR / Decision | Topic | Aktuelle Foundation-Semantik (Quelle) | Semantik etabliert | Noch offen | Cluster | F-0.1-Relevanz | HM-Input vor positivem Readiness-Verdikt? | Primäre Disposition |
|---|---|---|---|---|---|---|---|---|---|
| 1 | `ADR-0001` | Universelle Operations Control Plane | `PROJECT_BRIEF`, `SYSTEM_CONTEXT`, `DEC-P-01` `accepted-product` | **ja** | nein | — | Produktvision, bereits akzeptiert | nein | `candidate-may-no-longer-be-required` |
| 2 | `ADR-0002` | Modularer Monolith | `COREOPS_LOGICAL_MODULE_ARCHITECTURE` (17 Module, `module ≠ deployment unit`); `DEC-S-74/75` `deferred` | **partial** | **ja** | **C1** | Laufzeitarchitektur, post-Foundation | nein | `still-open` |
| 3 | `ADR-0003` | Read-only First | `INITIAL_SUPPORT_BOUNDARY`, Integration-Capability-Dimensionen, `EXECUTION_AUTHORIZATION`; `DEC-P-04` | **ja** | Durchsetzungsmechanik | **C6** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 4 | `ADR-0004` | Offline First | `RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL` (6 Connectivity Classes); `DEC-P-02` | **ja** | Transport-/Format-Technologie | **C4** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 5 | `ADR-0005` | Plane-Architektur | `COREOPS_PLANE_TAXONOMY` (10 Planes); **`DEC-O-01` / `CCR-01` weiterhin `open`** | **partial** | **ja** | — | offener Konflikt | **ja** (`CCR-01`) | `still-open` |
| 6 | `ADR-0006` | CoreOps Integration Contract | `COREOPS_INTEGRATION_CONTRACT_V0_1` (v0.1); `DEC-O-14` `open`, `DEC-S-166` `deferred` | **ja** | Protokoll/Schema/Transport | **C2** | Contract v0.1 vorhanden | nein | `foundation-semantics-established-technical-choice-pending` |
| 7 | `ADR-0007` | Standards First | `INITIAL_SUPPORT_BOUNDARY`, `INTEGRATION_CAPABILITY_AND_OPERATION_MODEL` | **ja** | Protokollauswahl | **C2** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 8 | `ADR-0008` | Native Widgets vor iFrames | `UX_INFORMATION_ARCHITECTURE`, `DASHBOARD_INFORMATION_HIERARCHY` (ohne Komponentenwahl); `DEC-S-254` `deferred` | **ja** | Frontend-/Komponenten-Stack | — | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 9 | `ADR-0009` | Source of Truth und Field Provenance | `SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL`, `FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD` (22 Felder); **`DEC-O-13` `open`** | **ja** | Konfliktprioritäten je Datenklasse | — | offene Teilfrage | **ja** (`DEC-O-13`) | `still-open` |
| 10 | `ADR-0010` | Desired State getrennt von Observed State | `OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL`; `DEC-P-07` | **ja** | — | **C3** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 11 | `ADR-0011` | Drift Detection vor Reconciliation | `DRIFT_DETECTION_AND_CONVERGENCE_MODEL`, `SAFE_REMEDIATION`; `DEC-S-134/135/136` `deferred` | **ja** | Engine, Auto-Remediation | **C3** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 12 | `ADR-0012` | Zentrale fail-closed Policy Engine | `POLICY_DECISION_AND_EVALUATION_MODEL` (Default-Deny, `indeterminate`/`conflicted` fail-closed); `DEC-S-151` `deferred` | **ja** | Policy-Engine-Technologie | **C6** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 13 | `ADR-0013` | Kurzlebige Machine Identities | `MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE`, `MACHINE_ENROLLMENT_AND_TRUST_LIFECYCLE`; **`DEC-O-06` / `CCR-06` `open`**; `DEC-D-02` post-Foundation | **partial** | **ja** | — | offener Konflikt | **ja** (`CCR-06`) | `still-open` |
| 14 | `ADR-0014` | Ausgehende Agent-Verbindungen | `COREOPS_PLANE_TAXONOMY` (Agent Plane optional, agentless möglich), `MACHINE_ENROLLMENT` | **ja** | Agent-/Transporttechnologie | — | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 15 | `ADR-0015` | Build und Deployment getrennt | `DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL`, `ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL` | **ja** | Pipeline-/Orchestrator-Technologie | **C1** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 16 | `ADR-0016` | Unveränderliche Deployment-Artefakte | `ARTIFACT_IDENTITY…` (`version ≠ revision`, `mutable alias ≠ final binding`) | **ja** | Registry-/Format-Technologie | **C7** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 17 | `ADR-0017` | Artifact Trust, SBOM und Provenance | `ARTIFACT_IDENTITY…`, `ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY`; `DEC-S-303` `deferred` | **ja** | SBOM-/Signing-/Scanner-Technologie | **C7** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 18 | `ADR-0018` | Versioniertes Eventmodell | `EVENT_AND_AUDIT_CORRELATION_MODEL` (20 Event-Klassen, 4 Zeitbegriffe); `DEC-S-236` `deferred` | **ja** | Event-Bus/Storage/Schema | — | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 19 | `ADR-0019` | Persistente und wiederaufnehmbare Workflows | **kein** Foundation-WP; `DEC-D-01` `deferred` post-Foundation; `RISK-10` `open` | **nein** | **ja** | — | ausserhalb Foundation 0.1 | nein | `deferred-post-foundation` |
| 20 | `ADR-0020` | Evidenzbasierte Topologie | `TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL`; `DEC-O-15`, `DEC-S-273` `deferred` | **ja** | Graph-/Discovery-Technologie | **C5** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 21 | `ADR-0021` | Manuelle Topologieautorität | `TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY` | **ja** | — | **C5** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 22 | `ADR-0022` | Domain Packs umgehen die Control Plane nicht | `DOMAIN_PACK_GOVERNANCE_MODEL`, `DOMAIN_PACK_TRUST…`; **`DEC-G-07` bereits `binding-governance`** | **ja** | keine Optionsspanne | — | Invariante, bereits bindend | nein | `candidate-may-no-longer-be-required` |
| 23 | `ADR-0023` | Configuration Vault | `SECRETS_CONFIGURATION_VAULT_AND_CUSTODY_GOVERNANCE`, `CONFIGURATION_SOURCE_OF_TRUTH…`; `DEC-S-105` `deferred`, `DEC-D-03` post-Foundation | **ja** | Vault/KMS/HSM/TPM/PKI | — | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 24 | `ADR-0024` | Data Classification und Redaction | `DATA_CLASSIFICATION_AND_HANDLING_MODEL`, `REDACTION_MINIMIZATION…`; **`DEC-O-08` / `CCR-08` `open`** | **ja** | Audit-vs-Redaction-Konflikt | — | offener Konflikt | **ja** (`CCR-08`) | `still-open` |
| 25 | `ADR-0025` | CoreOps Self-Protection | `SELF_PROTECTION_AND_CONTROL_PLANE_SAFETY_MODEL`, `DEGRADED_MODE…`, `RECOVERY_MODE…` | **ja** | Health/HA/Failover-Technologie | — | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 26 | `ADR-0026` | Integration Quality Levels | `DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL` (`SUP-0…3`/`SUP-D`), sechs Capability-Dimensionen | **ja** | — (reines Governance-Modell) | **C2** | Semantik vollständig | nein | `foundation-semantics-established-technical-choice-pending` |
| 27 | `ADR-0027` | PostgreSQL vor Graphdatenbank im MVP | `DEC-S-203` (Storage/DB) und `DEC-S-273` (Graph-DB) beide `deferred`; `DATA_OWNERSHIP…`, `SCHEMA_VERSIONING…` | **ja** (Datenmodell-Semantik) | **ja** (Technologie) | — | erste konkrete Technologie; kein Code in Foundation | nein | `deferred-post-foundation` |
| 28 | `ADR-0028` | Grafana als Analysewerkzeug | `TELEMETRY_*`; `DEC-S-254` `deferred`; `DEC-D-04` post-Foundation | **ja** (Telemetrie-Semantik) | **ja** (Technologie) | — | kein Code in Foundation | nein | `deferred-post-foundation` |
| 29 | `ADR-0029` | KI nur beratend | `DEC-P-05` `accepted-product` **und** `DEC-N-01`…`DEC-N-09` (`DEC-N-09` autonomer KI-Administrator = `non-goal`) | **ja** | keine Optionsspanne (Verbot) | — | bereits akzeptiert + Non-Goal | nein | `candidate-may-no-longer-be-required` |
| 30 | `ADR-0030` | CorePack und Offline Trust | `COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL`, `OFFLINE_TRUST_ACTIVATION…`; **`DEC-O-09` / `CCR-09` `open`**; `DEC-S-317` `deferred` | **ja** | Offline-First-Facetten | **C4** | offener Konflikt | **ja** (`CCR-09`) | `still-open` |
| 31 | `DEC-A-0031` | Delivery Baseline (Docker-first-Klassifikation) | `DEC-O-10` `proposed`: akzeptierte Delivery-/Betriebsanforderung, **keine** Anwendungsarchitektur, kein K8s-Zwang | **ja** | Verbindlichkeit des `proposed`-Status | — | **direkt release-relevant** | **ja** | `requires-human-decision-before-readiness` |
| 32 | `DEC-A-0032` | Release-Taxonomie / SemVer | `DEC-O-02` `proposed`; `RELEASE_TAXONOMY.md` `Proposed for acceptance`; Tag-Kandidat `v0.0.1-foundation` | **ja** | Verbindlichkeit des `proposed`-Status | — | **direkt release-relevant** | **ja** | `requires-human-decision-before-readiness` |

### 11.1 Dispositions-Tally

| Disposition | Anzahl | Kandidaten |
|---|---|---|
| `foundation-semantics-established-technical-choice-pending` | **18** | `ADR-0003, 0004, 0006, 0007, 0008, 0010, 0011, 0012, 0014, 0015, 0016, 0017, 0018, 0020, 0021, 0023, 0025, 0026` |
| `still-open` | **6** | `ADR-0002, 0005, 0009, 0013, 0024, 0030` |
| `deferred-post-foundation` | **3** | `ADR-0019, 0027, 0028` |
| `candidate-may-no-longer-be-required` | **3** | `ADR-0001, 0022, 0029` |
| `requires-human-decision-before-readiness` | **2** | `DEC-A-0031, DEC-A-0032` |
| **Summe** | **32** | vollständige Abdeckung, keine Doppelzählung |

**Bindende Klarstellungen:**

```
Foundation semantics established   ≠  ADR accepted
ADR candidate reviewed             ≠  ADR decided
ADR not needed for current readiness ≠ ADR permanently unnecessary
review recommendation              ≠  Human-Maintainer decision
```

**ADR akzeptiert: 0. ADR abgelehnt: 0. ADR-Dateien erzeugt: 0. Technologie ausgewählt: 0.**

---

## 12. ADR Consolidation Candidates

Sieben Cluster über 15 Kandidaten. **Nur Empfehlung** — es wurde **nichts** zusammengeführt, umnummeriert oder entfernt.

| Cluster | Kandidaten | Begründung der Überlappung | Empfehlung |
|---|---|---|---|
| **C1 — Laufzeit-/Deployment-Architektur** | `ADR-0002` (Modularer Monolith), `ADR-0015` (Build/Deployment getrennt) | Beide entscheiden dieselbe Frage aus zwei Richtungen: wie die Anwendung geschnitten **und** ausgeliefert wird. `MODULE_BOUNDARY_AND_DEPENDENCY_STANDARD` hat `module ≠ microservice ≠ deployment unit` bereits getrennt — die verbleibende offene Frage ist **eine**. | zu **einer** Runtime-/Delivery-Architektur-ADR zusammenfassen; `DEC-A-0031` (Docker-first) ist deren Delivery-Eingangsgröße |
| **C2 — Integrationsvertrag und Qualitätsstufen** | `ADR-0006` (Integration Contract), `ADR-0007` (Standards First), `ADR-0026` (Integration Quality Levels) | `CO-WP-014` und `CO-WP-015` haben Contract, sechs Capability-Dimensionen und `SUP-0…3`/`SUP-D` bereits **als ein zusammenhängendes Modell** definiert. `ADR-0007` und `ADR-0026` haben dadurch keine eigenständige Optionsspanne mehr; offen bleibt allein die Protokoll-/Schema-Wahl aus `ADR-0006`. | `ADR-0007` und `ADR-0026` als **Governance-Fragen** deklarieren (bereits beantwortet); technische ADR auf `ADR-0006` konzentrieren |
| **C3 — Zustand und Drift** | `ADR-0010` (Desired ≠ Observed), `ADR-0011` (Drift vor Reconciliation) | Der Decision Index führt sie bereits gemeinsam: `DEC-P-07` verweist auf „ADR-0010/0011". `CO-WP-012` hat beide Semantiken in einem Dokumentensatz etabliert. | zu **einer** State-/Drift-ADR zusammenfassen (Engine-Auswahl) |
| **C4 — Offline und CorePack Trust** | `ADR-0004` (Offline First), `ADR-0030` (CorePack und Offline Trust) | `CO-WP-023` definiert die Offline-Autoritätsgrenze **einmal autoritativ**; `CO-WP-024…026` referenzieren sie. Beide Kandidaten adressieren nach `CO-WP-023` dieselbe verbleibende Wahl (Paket-/Transfer-/Signaturtechnologie, `DEC-S-317`). | zu **einer** Offline-/CorePack-ADR zusammenfassen; `CCR-09` (`DEC-O-09`) ist deren Vorentscheidung |
| **C5 — Topologie-Evidenz und manuelle Autorität** | `ADR-0020`, `ADR-0021` | Der Decision Index führt sie bereits gemeinsam: `DEC-O-15` verweist auf „ADR-0020/0021". `CO-WP-020` hat beide in einem Dokumentensatz etabliert; die manuelle Autorität ist eine Facette des Evidenzmodells, keine eigene Technologiewahl. | zu **einer** Topologie-ADR zusammenfassen |
| **C6 — Policy-Durchsetzung** | `ADR-0003` (Read-only First), `ADR-0012` (zentrale fail-closed Policy Engine) | Read-only-First wird ausschließlich über die Policy-/Approval-/Execution-Guards aus `CO-WP-013` durchgesetzt. Eine separate ADR für `ADR-0003` hätte keine eigene Optionsspanne. | `ADR-0003` als **Governance-Prinzip** deklarieren; technische ADR auf `ADR-0012` konzentrieren |
| **C7 — Artefakt-Unveränderlichkeit und Trust** | `ADR-0016` (unveränderliche Artefakte), `ADR-0017` (Artifact Trust/SBOM/Provenance) | `CO-WP-022` behandelt Identität/Version/Revision und Trust/SBOM/Provenance in einem zusammenhängenden Modell; die verbleibende Wahl (Registry, Format, Hash, Signing, Trust Anchor) ist **eine** und in `DEC-S-303` gemeinsam deferriert. | zu **einer** Supply-Chain-ADR zusammenfassen |

**Weitere Beobachtung:** `ADR-0001`, `ADR-0022` und `ADR-0029` sind **keine** Architekturoptionsvergleiche mehr, sondern akzeptierte Produkt- bzw. Governance-Entscheidungen (`DEC-P-01`, `DEC-G-07`, `DEC-P-05` + `DEC-N-09`). Empfehlung: als *reine Governance-Fragen* deklarieren statt als technische ADR-Kandidaten führen. **Das ist keine Entwertung** — `ADR not needed for current readiness ≠ ADR permanently unnecessary`.

**Nicht** getan: kein Kandidat umnummeriert, zusammengeführt, entfernt oder abgelehnt.

---

## 13. Foundation-0.1 ADR / HM-Decision Dependencies

### 13.1 Grundlage

`FOUNDATION_SCOPE_LOCK.md` §Exit Gates nennt **24** Foundation Exit Gates, darunter ausdrücklich *„relevante ADRs durch Human Maintainer entschieden"* und *„konsistenter Decision Index"*. **Keines** dieser Gates ist als erfüllt markiert.

Foundation 0.1 ist ein **Dokumentations-/Governance-Release** (`v0.0.1-foundation`) **ohne** Anwendungscode. Daraus folgt:

```
Foundation 0.1 = Dokumentationsrelease
→ keine implementierungsbindende ADR ist technisch erzwungen
→ aber: "relevante ADRs entschieden" ist ein Exit Gate
→ die Definition von "relevant" ist selbst eine Human-Maintainer-Entscheidung
```

**Diese Definition zu treffen ist Aufgabe von `CO-WP-030`, nicht dieses Reviews.**

### 13.2 Kandidaten mit direkter Foundation-0.1-Relevanz

| Kandidat | Warum release-relevant | Klassifikation |
|---|---|---|
| `DEC-A-0032` (Release-Taxonomie / SemVer) | Der Foundation-Tag-Kandidat `v0.0.1-foundation` **ist** das Ergebnis dieser Entscheidung. `RELEASE_TAXONOMY.md` steht auf `Proposed for acceptance`, `DEC-O-02` auf `proposed`. Ein Release lässt sich nicht vorbereiten, solange die Taxonomie unverbindlich ist. | **HUMAN DECISION REQUIRED** |
| `DEC-A-0031` (Delivery Baseline / Docker-first) | `DEC-O-10` `proposed`; die Einordnung „akzeptierte Delivery-/Betriebsanforderung, **keine** Anwendungsarchitektur" ist die Grundlage mehrerer Foundation-Aussagen und des `RISK-20`-Treatments. | **HUMAN DECISION REQUIRED** |

### 13.3 Kandidaten ohne aktuellen Readiness-Blocker

Die übrigen **30** Kandidaten sind `no-readiness-blocker-currently` für das Dokumentationsrelease: Foundation 0.1 wählt keine Technologie aus, behauptet keine Implementierung und erzeugt keinen Code. Die vier `still-open`-Kandidaten mit offenem `CCR` (`ADR-0005`/`CCR-01`, `ADR-0013`/`CCR-06`, `ADR-0024`/`CCR-08`, `ADR-0030`/`CCR-09`) sowie `ADR-0009`/`DEC-O-13` bleiben jedoch **Readiness Notes**: `CO-WP-030` muss entscheiden, ob ein Foundation-Release mit sechs weiterhin offenen `CCR`-Konflikten als konsistent gelten darf.

> **Dieses Review trifft diese Entscheidung nicht** und sagt weder „Foundation ready" noch „Foundation not ready".

---

## 14. Risk Register Integrity Review

### 14.1 Mechanische Integrität — vollständig neu berechnet

| Prüfung | Ergebnis |
|---|---|
| Risk-IDs eindeutig | **PASS** — 316 eindeutige IDs, **0** Duplikate |
| Nummernkreis lückenlos | **PASS** — `RISK-01` … `RISK-316`, **keine** Lücke |
| Zeilen gesamt | **316** |
| Risk-Level-Tally (aus Spalte berechnet) | `high` **170** · `medium` **122** · `low` **24** |
| Risk-Level-Arithmetik | **PASS** — 170 + 122 + 24 = **316** ✓ |
| Status-Tally (aus Spalte berechnet) | `treatment-planned` **299** · `open` **17** |
| Status-Arithmetik | **PASS** — 299 + 17 = **316** ✓ |
| Publizierte Summen vs. berechnete | **PASS** — §Verteilung nach Risk Level nennt exakt 170/122/24 und 299/17 |
| Publizierte ID-Listen vs. Summen | **PASS** — die aufgezählten IDs decken sich vollständig mit den Tabellenwerten |
| `accepted-by-human` ohne HM-Entscheidung | **PASS** — **0** Einträge |
| Ungültige Feldwerte (Likelihood/Impact/Status) | **PASS** — nur zulässige Werte |
| Risiken ohne Target WP | **PASS** — alle verknüpft |

> Der `CO-WP-028`-Tally-Hinweis (Level-Summen aus den ID-Listen neu berechnet) hat gehalten: die Verschiebung von vier Einträgen besteht **nicht** mehr. Die Register-Arithmetik ist zum Baseline-Stand `b7827b8` vollständig korrekt.

### 14.2 Duplikat- und Familienprüfung

Keine echten Duplikat-IDs. Die vom `CO-WP-005…012` Milestone Review benannten „Interpretationsfamilien" (Zustands-Fehlinterpretation, Offline-Overwrite, Stale-als-aktuell) bestehen fort, sind aber **explizit als Nicht-Duplikate begründet**: jede Pro-WP-Sektion listet unter „Bereits abgedeckt und nicht dupliziert" die referenzierten Vorgänger-IDs. Das ist eine dokumentierte Konsolidierungs-Entscheidung, kein Defekt.

### 14.3 `CO-WP-028`-Ziel-Risiken (§19 des Auftrags)

**21** Risiken tragen weiterhin Target WP `CO-WP-028`: `RISK-233, 261, 263, 265, 283, 287, 288, 291, 293, 299, 300, 301, 302, 304, 305, 306, 307, 308, 309, 311, 313`. Alle stehen auf `treatment-planned`; **keines** ist geschlossen.

| Klassifikation | Ergebnis |
|---|---|
| **A — Teststrategie-Behandlungskomponente dokumentiert** | **zutreffend für alle 21.** Die Risk-to-Test-Matrix in `FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md` bildet jedes der 21 mit Test Layer, Fixture-/Szenarioklasse, erwartetem geschütztem Verhalten, geforderter Evidenz, Grenzen und Future Execution Gate ab. |
| **B — künftige ausführbare Validierung weiterhin erforderlich** | **zutreffend für alle 21.** `test strategy documented ≠ tests implemented ≠ tests executed ≠ Foundation validated`. Ohne Testausführung existiert keine Behandlungsevidenz. |
| **C — Target WP später zu überdenken** | **Empfehlung für alle 21**, aber **nicht jetzt**. Solange kein ausführendes Test-WP existiert, wäre jede Retargetierung eine Verlagerung auf ein nicht existierendes Ziel. `RISK-314` zeigt das bereits gültige Muster: Target `CO-WP-030`, `RISK-316` Target `later test WP`. |
| **D — keine Änderung nötig** | **aktuell zutreffend.** |

**Ergebnis: keine Retargetierung, keine Statusänderung, keine Schließung.** Empfehlung an `CO-WP-030`: die 21 Target-WP-Werte erst dann anfassen, wenn ein ausführendes Test-WP in der Queue registriert ist.

### 14.4 `RISK-66` — Sonderfall

`RISK-66` („Capability-Zählungsinkonsistenz „74" vs. 94", `low`/`low`/`low`, `treatment-planned`) trägt Target WP `later consistency WP` und als geforderte Behandlung *„Documentation-Consistency-Abgleich der „74"-Referenzen in späterem WP"*.

**Dieses WP ist der benannte Abgleich**, und die Behandlung ist mit §16 materiell ausgeführt. **Status und Target wurden dennoch nicht geändert** — eine Schließung ist eine Human-Maintainer-Entscheidung, kein mechanischer Reconciliation-Schritt. → `CO-WP-030` / HM (HM-8).

### 14.5 Vier auf `CO-WP-029` gerichtete Risiken

`RISK-191` (Policy-Konflikt still aufgelöst), `RISK-243` (API-Compatibility fehlbeurteilt), `RISK-249` (Deprecated/Retired API Identity wiederverwendet), `RISK-275` (Timestamp erzwingt unsafe Last-Write-Wins).

Alle vier sind **Design-/Laufzeitrisiken**, deren Behandlung eine Implementierung voraussetzt — sie sind durch ein Konsistenzreview **nicht** behandelbar. Ihre Zuordnung zu `CO-WP-029` ist vermutlich ein Platzhalter. **Empfehlung:** Target-WP-Review durch HM; **nicht** geändert. → `CO-WP-030` (HM-9).

---

## 15. Risk Calibration Observations

### 15.1 Status der Modal-Abbildung — analytische Heuristik, keine Norm

Bevor Zahlen genannt werden, die Einordnung. Binding:

```
Modal-Abbildung dieses Reviews   =   rein analytische Heuristik
Modal-Abbildung dieses Reviews   !=  verbindliche Registernorm
Abweichung von der Modal-Abbildung !=  registrierter Risiko-Defekt
Abweichung von der Modal-Abbildung !=  nachgewiesene Fehlbewertung
24 abweichende Zeilen            =   Severity-Kalibrierungs-Reviewkandidaten
24 abweichende Zeilen            !=  24 fehlbewertete Risiken
```

Es ist **derzeit keine verbindliche Likelihood × Impact → Risk-Level-Abbildung veröffentlicht**. Ohne eine solche Norm existiert **kein Massstab**, gegen den eine Zeile als falsch bewertet gelten könnte. Die nachfolgende Modal-Abbildung wurde von diesem Review **aus dem Bestand selbst abgeleitet**, um Kalibrierungsfragen überhaupt sichtbar zu machen — sie ist **kein** Sollwert und **darf nicht als solcher gelesen oder weiterverwendet werden**.

Folglich sind die 24 identifizierten Zeilen **Kandidaten für ein späteres Severity-Kalibrierungs-Review**, ausdrücklich **nicht** bestätigte Fehlbewertungen. **Sämtliche registrierten Severities bleiben unverändert gültig.**

### 15.1.1 Befund

Das Register definiert Risk Level als *„qualitative Ableitung aus Likelihood × Impact"*, **veröffentlicht die Ableitungsregel aber nicht**. Die tatsächliche Abbildung im Bestand:

| Likelihood × Impact | → `high` | → `medium` | → `low` | Modalwert | Kandidatenzeilen |
|---|---|---|---|---|---|
| `low` × `low` | — | — | 10 | `low` | 0 |
| `low` × `medium` | — | 8 | 11 | `low` | **8** |
| `low` × `high` | 147 | 6 | — | `high` | **6** |
| `medium` × `low` | — | — | 3 | `low` | 0 |
| `medium` × `medium` | — | 96 | — | `medium` | 0 |
| `medium` × `high` | 18 | 10 | — | `high` | **10** |
| `medium` × `critical` | 4 | — | — | `high` | 0 |
| `high` × `medium` | — | 2 | — | `medium` | 0 |
| `high` × `high` | 1 | — | — | `high` | 0 |

**24 Zeilen** weichen vom jeweiligen Modalwert ab und sind damit **Kalibrierungs-Reviewkandidaten** (keine nachgewiesenen Fehlbewertungen):

`RISK-07, 10, 12, 15, 21, 23, 26, 34, 36, 39, 46, 47, 55, 62, 67, 73, 77, 78, 85, 109, 113, 115, 152, 310`

Zusätzlich ist die Abbildung **asymmetrisch**: `medium × high` → überwiegend `high`, `high × medium` → `medium`.

### 15.2 Warum nichts geändert wurde

Die Voraussetzungen für eine mechanische Severity-Korrektur sind **nicht** erfüllt:

- ✗ **Keine etablierte Kalibrierungsregel.** Die Modal-Abbildung ist eine *analytische Ableitung dieses Reviews*, keine im Repository dokumentierte Regel. Sie zur Norm zu erheben wäre eine **neue Entscheidung**. Ohne Norm ist eine Abweichung **kein Fehler**, sondern lediglich eine **Auffälligkeit**.
- ✗ **Nicht eindeutig.** In sechs von neun Zellen mit Abweichung ist plausibel, dass die abweichende Bewertung die **inhaltlich korrektere** ist (etwa `RISK-15` Technologie-Lock-in: `low × high → medium` ist bei reiner Dokumentations-Foundation gut begründbar).
- ✗ **Nicht mechanisch.** Jede Änderung wäre eine inhaltliche Neubewertung.

**Folglich: 0 Severity-Änderungen. Keine Massen-Rekalibrierung. Keine Quote. Keine Zielverteilung.**

### 15.3 Empfehlung an `CO-WP-030` / Human Maintainer

1. **Zuerst die Regel entscheiden**, dann den Bestand — nicht umgekehrt. Eine explizite Likelihood × Impact → Level-Matrix im Register-Kopf beseitigt die Ursache dauerhaft.
2. Erst danach die 24 Abweichler prüfen: bestätigen (mit Begründung als dokumentierte Ausnahme) oder angleichen.
3. Die aus dem `CO-WP-021…026` Milestone Review stammende Beobachtung („Severity verliert Trennschärfe: 29 von 30 neuen Risiken `high`") ist damit **präzisiert**: die Ursache ist nicht Inflation, sondern das Fehlen einer publizierten Regel bei gleichzeitiger Dominanz von `low × high` (147 Zeilen) im Bestand.
4. `CO-WP-027` (1× `high`, 3× `medium`) und `CO-WP-028` (2× `high`, 1× `medium`) haben bereits bewusst differenziert — die Gegenmaßnahme läuft.

**Keine** Risk-Severity wurde in diesem WP geändert. **Neue Risiken: 0.** `RISK-317` wurde **nicht** angelegt. **Keine Massen-Rekalibrierung, keine Zielverteilung, keine Quote.**

**Abschliessende Einordnung:** Das Risk Register weist zum Stand `b7827b8` **keinen** Kalibrierungsdefekt auf, denn es existiert keine verletzte Norm. Was existiert, ist eine **nicht veröffentlichte Ableitungsregel** — und daraus folgt eine Empfehlung, keine Beanstandung. Die 24 Zeilen sind der Arbeitsvorrat für ein späteres Kalibrierungs-Review, **nicht** eine Mängelliste.

---

## 16. Capability and Count Reconciliation

### 16.1 Autoritative Quelle

`docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md` — deterministisch nachgezählt:

| Metrik | Berechnet | In der Matrix ausgewiesen | Ergebnis |
|---|---|---|---|
| Eindeutige Capability-IDs | **94** | 94 | **PASS** |
| Capability-Zeilen (Haupttabellen) | 94 | 94 | **PASS** |
| Duplikate | 0 | „ohne Duplikate" | **PASS** |
| Eindeutige Domänen-Präfixe | **13** | **„12"** | **FAIL → korrigiert** |
| Implementation Status | 94× `not-implemented` | 94× (100 %) | **PASS** |
| Support Status | 94× `not-supported` | 94× (100 %) | **PASS** |

**Domänen (13):** `AUTOMATION` (12) · `DEPLOY` (22) · `DISCOVERY` (16) · `IDENTITY` (12) · `INVENTORY` (14) · `MONITORING` (16) · `NETWORK` (16) · `PLATFORM` (14) · `PRINT` (16) · `PROTECT` (16) · `TOPOLOGY` (4) · `TRUST` (18) · `VIRT` (12).

> **MAJ-7:** Die Matrix nannte „Domains: 12", **listete in derselben Klammer aber 13 Namen** (Platform · Identity · Inventory · Discovery · Monitoring · Network · Topology · Print · Virtualization · Deployments · Automation · Trust · Protection) und trägt 13 ID-Präfixe. Selbstwidersprüchliche autoritative Metadaten → nach §29 des Auftrags korrigierbar. **Korrigiert auf 13** mit Korrekturnotiz; **keine** Domäne hinzugefügt, entfernt, umbenannt oder umzugeordnet, **keine** Capability verändert.

### 16.2 Klassifikation aller Capability-Count-Referenzen

| Pfad | Aussage | Klassifikation | Aktion |
|---|---|---|---|
| `FOUNDATION_CAPABILITY_MATRIX.md` §Zusammenfassung | 94 Capabilities | **korrekt** | unverändert |
| `FOUNDATION_CAPABILITY_MATRIX.md` §Zusammenfassung | „Domains: 12" | **inkorrekt** | **korrigiert → 13** |
| `FOUNDATION_CAPABILITY_MATRIX.md` §Alignment-Bestätigung | 94 zugeordnet | **korrekt** | unverändert |
| `CAPABILITY_SECURITY_AND_GOVERNANCE_ALIGNMENT.md` | 94 Zeilen / 94 IDs | **korrekt** | unverändert |
| `COREOPS_MODULE_CATALOG.md` (2×) | „94-capability list" | **korrekt** | unverändert |
| `RISK_REGISTER.md` `RISK-111` | „die 94 Capabilities" | **korrekt** | unverändert |
| `RISK_REGISTER.md` `RISK-66` | „Mehrere Dokumente nennen „74", die Matrix enthält 94 Zeilen" | **korrekt** (beschreibt den Defekt) | unverändert (§14.4) |
| `DECISION_INDEX.md` `DEC-O-17` Notes | „74 Capabilities" | **stale current-state claim** | **korrigiert → 94** |
| `PROJECT_BRAIN.md` §Capability Matrix (`CO-WP-004`) | „74 Capabilities, 12 Domains" | **stale current-state claim** | **korrigiert → 94 / 13** |
| `PROJECT_PROFILE.md` §`CO-WP-004` Konsolidierung | „74 Capabilities über 12 Domains" | **stale current-state claim** | **korrigiert → 94 / 13** |
| `CONTEXT_PACK_FOUNDATION_0_1.md` §`CO-WP-004` | „74 Capabilities, 12 Domains" | **stale current-state claim** | **korrigiert → 94 / 13** |
| `CONTEXT_PACK_FOUNDATION_0_1.md` §Bekannte Notes | „Capability Matrix (74 Capabilities)" | **stale current-state claim** | **korrigiert → 94 / 13** |
| `CONTEXT_PACK_FOUNDATION_0_1.md` §Kompakte Historie | „`CO-WP-004`: … (74 Capabilities)" | **historisch gerahmt, aber sachlich falsch** | **korrigiert mit historischer Markierung** |
| `NEXT_PHASE.md` §Capability Matrix | „(74 Capabilities)" | **historisch gerahmt, aber sachlich falsch** | **korrigiert mit historischer Markierung** |
| `MILESTONE_REVIEW_CO_WP_005_TO_012.md` (3×) | „74" als Befund/Follow-up | **HISTORICAL RECORD** (Reviewbefund, zum Zeitpunkt korrekt) | **unverändert** |
| `MILESTONE_REVIEW_CO_WP_013_TO_020.md`, `…_021_TO_026.md` | „74→94" als Follow-up | **HISTORICAL RECORD** | **unverändert** |
| `LESSONS_LEARNED_REGISTER.md` | „74" statt 94 als Lesson-Beobachtung | **HISTORICAL RECORD** + forbidden file | **unverändert** |

**Referenzen geprüft: 18. Mechanische Korrekturen: 8** (7 Capability-Count + 1 Domain-Count).

Die Capability-Inventur wurde **nicht** neu definiert. Keine Capability hinzugefügt, entfernt, umbenannt oder hochgestuft.

### 16.3 Weitere Zähl- und Tally-Prüfungen

| Claim | Behauptet | Berechnet | Ergebnis |
|---|---|---|---|
| Decision-IDs `DEC-S` | `DEC-S-373` | 373, lückenlos, duplikatfrei | **PASS** |
| ADR-Kandidaten | 32 | 32 (30 + 2), vier Quellen übereinstimmend | **PASS** |
| Risiken gesamt | 316 | 316 | **PASS** |
| Risk-Level-Tally | 170 / 122 / 24 | 170 / 122 / 24 | **PASS** |
| Risk-Status-Tally | 299 / 17 | 299 / 17 | **PASS** |
| Threat Scenarios | 40 (`THR-001…040`) | 40 eindeutig | **PASS** |
| Threat-Model-Assets | 24 (`AST-01…24`) | 24 eindeutig | **PASS** |
| Trust Boundaries | 11 (`TB-01…11`) | 11 eindeutig | **PASS** |
| Logische Module | 17 (`MOD-*`) | 17 eindeutig, repository-weit konsistent | **PASS** |
| PSR-Domänen | 18 (`PSR-01…18`) | 18 eindeutig | **PASS** |
| Lessons | 38 (`LL-001…038`) | 38 eindeutig | **PASS** |
| NDF-Feedback-Kandidaten | 15 | 15 eindeutig | **PASS** |
| Work Packages in der Queue | `CO-WP-001…031` + `CO-WP-001A` + 4 `004B*`-Varianten | Queue-Tabelle konsistent | **PASS** |
| `CO-WP-028`-Ziel-Risiken | 21 | 21 | **PASS** |
| Capabilities | 94 | 94 | **PASS** |
| Capability-Domänen | „12" | **13** | **FAIL → korrigiert** |

**Ein einziger Zähl-Defekt in der gesamten Foundation.** Das ist ein bemerkenswert gutes Ergebnis für 105 Dokumente und 26.245 Zeilen.

---

## 17. Current-State vs. History Review

### 17.1 Angewandte Regel

```
current-state mirror              ≠  historical record
historisch korrekte Aussage       ≠  current-state defect
```

| Klasse | Behandlung | Anzahl |
|---|---|---|
| `CURRENT-STATE DEFECT` | Korrektur zulässig und durchgeführt | **5** (MIN-1…MIN-5) |
| `CURRENT-STATE DEFECT`, aber forbidden file | Befund, keine Korrektur | **1** (MIN-12, `ROADMAP.md`) |
| `HISTORICAL RECORD` | erhalten, unverändert | **alle** Milestone Reviews, Lessons, WP-Historien |
| `AMBIGUOUS` | Befund, keine stille Umschreibung | **5** (MIN-6…MIN-10) |

### 17.2 Reconciliation auf Baseline `b7827b8`

| Element | Vorher | Nachher |
|---|---|---|
| `CO-WP-027` | `completed-go-with-notes`, Commit `1dee29d`, gepusht | **unverändert** (war korrekt) |
| `CO-WP-028` | `completed-go-with-notes`, **„Human-Maintainer-Commit ausstehend"** | `completed-go-with-notes`, **Commit `b7827b8`, gepusht** |
| Repository-Baseline | `1dee29d` | **`b7827b8`** |
| „Verbindlicher nächster Schritt" | „Human-Maintainer-Commit von `CO-WP-028`" | **„Human-Maintainer-Commit von `CO-WP-029`"** (nach Nova Review `GO WITH NOTES` und geschlossener Notes-Runde) |
| `CO-WP-029` | `planned` / `retained` / „nicht begonnen" | **`completed-go-with-notes`** — Nova Review `GO WITH NOTES`, Notes 1–3 geschlossen (§25), Human-Maintainer-Commit ausstehend |
| `CO-WP-030` | `planned` / `retained` | **unverändert** — `retained`, nicht freigegeben, nicht begonnen |
| `CO-WP-031` | `planned` / `retained` | **unverändert** — `retained`, nicht freigegeben, nicht begonnen |
| `CO-WP-024` (in `NEXT_PHASE.md`) | Selbstwiderspruch „Commit ausstehend" vs. „Commit `916ba66` gepusht" | **konsistent: Commit `916ba66`, gepusht** |

### 17.3 Ausdrücklich erhaltene historische Aussagen

- „Repository-Baseline war `1dee29d` während der `CO-WP-028`-Bearbeitung" — historisch korrekt, als solche markiert erhalten.
- Alle Pro-WP-Registerdeltas (`Risk Register +5, gesamt 294` usw.) — Momentaufnahmen zum jeweiligen WP-Zeitpunkt, **nicht** angetastet.
- Alle Milestone-Review-Tallies (`Risk Register 279 (high 138/medium 117/low 24)`, `309 (high 167/medium 118/low 24)`) — zum Reviewzeitpunkt korrekt, **nicht** auf 316 „aktualisiert".
- Alle `Nova Review pending`-Vermerke in `NEXT_PHASE.md` für `CO-WP-004C`…`CO-WP-012` — Zustand zum Eintragszeitpunkt.

**Historische Aufzeichnungen wurden ausschließlich dort berührt, wo sie einen sachlich falschen Zählwert transportierten** (Capability-Count) — dort mit expliziter historischer Markierung statt stiller Umschreibung.

---

## 18. Parallel-Model Review

Geprüft auf die im Auftrag genannten dreizehn Muster.

| Gesuchtes Muster | Befund | Klassifikation |
|---|---|---|
| Zweites Autoritätsmodell | **keines.** Repository-Autorität (HM) und Runtime-Autorität (`MOD-IAM-001`/`MOD-POL-001`) sind durchgängig getrennt; `CO-WP-021`…`028` binden jede Write-/Execution-Aktion an `CO-WP-013` | — |
| Zweites Statusmodell | **keines.** Die elf UX-Dimensionen sind per Notes-Closure ausdrücklich **Darstellungsordnung**; die sechs Testergebniszustände sind ausdrücklich **nur** auf Testausführung bezogen | `presentation-only view` / `domain-specific specialization` |
| Zweites Scope-Modell | **keines.** `CO-WP-027` Note 2: „Presentation Context referenziert bestehende Scope-Identitäten; kein zweites Scope-Modell"; die acht RBAC-Scopes bleiben allein autoritativ | `compatible extension` |
| Zweite Evidence-Autorität | **keine.** `CO-WP-028`: Test-Evidenz ist eine Evidence-**Klasse innerhalb** `MOD-EVD-001` / `CO-WP-018` | `domain-specific specialization` |
| Zweites Operational-Mode-Modell | **keines.** `CO-WP-023` §14 definiert Operational States; `CO-WP-026` **erweitert** sie auf zehn Modes; die Lab-Umgebungsprofile (`CO-WP-028`) setzen auf denselben Klassen auf und deklarieren „kein zweites Betriebsmodusmodell" | `compatible extension` |
| Zweiter Deployment-Lebenszyklus | **keiner.** `CO-WP-021` ist die einzige Deployment-Lifecycle-Autorität | — |
| Zweiter Recovery-Lebenszyklus | **keiner.** `CO-WP-026` definiert 15 Recovery-Stufen; `CO-WP-023`/`024`/`025` referenzieren sie | — |
| Zweiter Test-Evidence-Lebenszyklus | **keiner.** `CO-WP-028` Note 2 schließt ihn ausdrücklich aus; Alterung bleibt beim bestehenden Evidence-Freshness-Modell | — |
| Zweites Support-Modell | **keines.** `SUP-0…3`/`SUP-D` (`CO-WP-015`) und die sechs Capability-Dimensionen (`CO-WP-014`) sind komplementär, nicht konkurrierend | `compatible extension` |
| Globaler Health-Boolean | **keiner.** `DASHBOARD_INFORMATION_HIERARCHY` §11: „CoreOps kennt **keinen** globalen Gesundheits-Boolean und **keinen** normativen Aggregat-Score" | — |
| Universeller Status-Score | **keiner.** `CoreScore` ist per §19 **objektbezogen** und begleitangabenpflichtig; explizit `CoreScore ≠ globaler Plattform-Gesundheitswert` | `domain-specific specialization` |
| Neuer impliziter Approval-Mechanismus | **keiner.** `UI confirmation ≠ approval record` (`CO-WP-027`); `lab approval ≠ production deployment authorization` (`CO-WP-028`) | — |
| Neue versteckte Execution Authority | **keine.** Experience-Ebene (`MOD-EXP-001`) löst keine privilegierte Ausführung aus; Lab-Rollen sind ausdrücklich **keine** Laufzeitmodule und erhalten **keine** neuen `MOD-*`-IDs | — |

**Ergebnis: 0 Parallelmodelle · 0 tatsächliche Contradictions.** Wiederholte Konzepte wurden **nicht** pauschal als Konflikt gewertet: die durchgängige Praxis „referenzieren statt duplizieren" ist im Repository nachweisbar und wird von den WPs selbst deklariert (29 explizite „kein zweites …"-/„kein Parallelmodell"-Aussagen).

---

## 19. Incident-Derived Extension Review

Erneute Bewertung der acht Themen aus dem `CO-WP-021…026` Milestone Review — **ausschließlich** als Konsistenz-/Zukunftsdesign-Kandidaten. **Keine Implementierung, keine WP-ID, keine Technologie, keine Produkt-Spezifika.**

| Thema | Klassifikation `CO-WP-021…026` | Status nach `CO-WP-027`/`028` | ADR-relevant? |
|---|---|---|---|
| Consumer-gebundene Dependency Contract Verification | `partial extension` | **unverändert `partial`.** `CO-WP-027` benennt `provider/process healthy ≠ consumer capability usable` als Ehrlichkeitsgrenze; `CO-WP-028` fordert Consumer/Provider-Mismatch als **Pflichtszenario** — beide **ohne** die Erweiterung zu implementieren oder zu registrieren | **nein** (Modellerweiterung, keine Technologiewahl) |
| Recovery Set / Recovery-Point-Komposition | `genuine future extension` | **unverändert.** Kein neues Dokument berührt es | **nein** |
| Recovery Scope Discovery & Completeness Assessment | `partial extension` | **unverändert** | **nein** |
| Erfasster (captured) Vorzustand als Wiederherstellungsziel | `genuine future extension` | **unverändert.** Repository-weit weiterhin kein Pre-State-Capture-Konzept | **nein** |
| Expected-vs-Actual-Simulationstreue | `genuine future extension` | **unverändert.** `CO-WP-028` etabliert Expected-Outcome-Deklaration für **Fixtures** — eine andere Vergleichsachse, **kein** Ersatz | **nein** |
| Cleanup/Restoration temporärer Zustände | `partial extension` (in `CO-WP-024` secret-lokal vollständig) | **unverändert.** `CO-WP-028` Lab-Reset/Disposability ist laborlokal, **nicht** produktionsgeneralisiert | **nein** |
| Protected Configuration Drift | `genuine future extension` | **unverändert.** Orthogonale Änderungskontext-Klassifikation, ersetzt die Detection States **nicht** | **nein** |
| Path-/scope-gebundene Health | `already covered` | **verstärkt.** `CO-WP-027` fügt Freshness/Coverage/Confidence als **Pflichtangaben** jeder Aggregation hinzu | **nein** |
| „Safe Change Transaction" als Modul/Autorität | `rejected parallel model` | **weiterhin abgelehnt.** Würde mit Remediation- und Deployment-Lebenszyklus konkurrieren | — |
| Globales Health-Boolean / globales Statusmodell | `rejected parallel model` | **weiterhin abgelehnt** und in `CO-WP-027` §11/§19 explizit ausgeschlossen | — |

**Ergebnis:** Alle Klassifikationen der Coverage-Matrix (5 `already covered` · 4 `partially covered` · 4 `genuine extension` · 2 `duplicate/reject` = 15) gelten **unverändert**. Die vier echten Erweiterungen bleiben spätere Design-Anliegen, **keine** ist ADR-relevant (sie erfordern Modellpräzisierung, keine Technologiewahl), **keine** blockiert Foundation Readiness.

**Nicht getan:** keine Implementierung, keine neue WP-ID, keine Technologie, keine Übertragung produktspezifischer Vorfallsemantik in CoreOps-Kernsemantik.

---

## 20. Documentation Economy Review

| Beobachtung | Bewertung | Empfehlung |
|---|---|---|
| Vier Current-State-Spiegel führen je eine vollständige Pro-WP-Narrative (zum Reviewzeitpunkt: `NEXT_PHASE.md` 549 Z. · `PROJECT_PROFILE.md` 506 Z. · `PROJECT_BRAIN.md` 503 Z. · `CONTEXT_PACK_FOUNDATION_0_1.md` 385 Z.) | **Ursache von MIN-1…MIN-4.** Ein Faktum ändert sich → vier Stellen müssen nachgezogen werden → mindestens eine bleibt zurück | **future consolidation candidate** — Router + Referenz statt vierfacher Volltext; **nicht** in diesem WP |
| Große Contract-/Invariantenblöcke wortgleich in mehreren Routern (z. B. die `CO-WP-028`-Claim-Boundary-Liste in vier Dokumenten) | Redundanz ohne Widerspruch | **future consolidation candidate** — gemeinsames Invarianten-/Claim-Referenzdokument (bereits als `LL-030`-Follow-up vorgemerkt) |
| `PROJECT_PROFILE.md` §§1–17 (`CO-WP-001`-Originalprofil) enthält überholte Aussagen: §9 „ausschließlich das Core Governance Skeleton", §14 „Concept v3.0 noch nicht vollständig übernommen", §15 offene Fragen, §16 „Nächster Meilenstein `CO-WP-002`" | **AMBIGUOUS.** Historisch gerahmt („Erzeugt durch `CO-WP-001`"), aber current-state formuliert und im Router-Pfad | **surgical correction candidate für `CO-WP-030`** — entweder als „Historischer Ursprungsstand (`CO-WP-001`)" kennzeichnen oder in einen Historienabschnitt verschieben. **Nicht** in diesem WP: es wäre Prosa-Umbau, kein mechanischer Fix |
| `PROJECT_BRAIN.md` §Bekannte Risiken / §Offene Fragen tragen `CO-WP-002`/`CO-WP-004`-Stand (39 Risiken, „Capability Matrix → `CO-WP-004`") | **AMBIGUOUS**, gleiche Ursache | **surgical correction candidate für `CO-WP-030`** |
| `RISK_REGISTER.md` enthält 33 Pro-WP-Prosaabschnitte zusätzlich zur Tabelle | Duplizierte Wortwahl, keine widersprüchlichen Zahlen | **keep** — sie tragen die „nicht dupliziert weil…"-Begründungen, die die Duplikatprüfung stützen |
| ADR-Kandidatenliste doppelt (Concept §51 + Klassifikationsdatei) | Concept ist unveränderlich registriert | **keep** |
| Widersprüchlich duplizierte Zählwerte | **keine gefunden** (nach den Korrekturen aus §16) | **keep** |
| Doppelte „Current Goal"-Abschnitte | **keine gefunden** | **keep** |

**Keine breite Prosabereinigung durchgeführt. Die Foundation wurde nicht stilistisch umgeschrieben.**

---

## 21. NDF / CDS / Ecosystem Boundary Review

### 21.1 NDF

| Prüfung | Ergebnis |
|---|---|
| Normative Basis | `v1.0.0` / Tag `v1.0.0` / Commit `9dcadc1` — in allen Router-Dokumenten identisch deklariert | **PASS** |
| `main` als informativ markiert | durchgängig | **PASS** |
| Stiller Import neuerer NDF-Semantik | **keiner** | **PASS** |
| NDF-Transfer in diesem WP | **keiner** | **PASS** |
| `LESSONS_LEARNED_REGISTER.md` | **unverändert** (read-only geprüft: 38 Lessons, konsistent) | **PASS** |
| `NDF_FEEDBACK_CANDIDATES.md` | **unverändert** (read-only geprüft: 15 Kandidaten, `008…015` `candidate-pending-nova-review`) | **PASS** |
| Skills-Pack-Provenance | 38 Skills / 39 Dateien, byte-identisch, `9dcadc1` — Lock und Provenance konsistent | **PASS** |
| Extern offen | NDF-Release-Zuordnung der drei Adoption-Commits (`001…007`) bleibt `not yet assigned` — korrekt so ausgewiesen | **PASS** |

**NDF boundary: preserved.** `adopted-in-ndf ≠ in veröffentlichter NDF-Version enthalten` ist durchgängig korrekt formuliert.

### 21.2 CDS

| Prüfung | Ergebnis |
|---|---|
| CDS Adoption | `none` / `Not started` — in `UX_INFORMATION_ARCHITECTURE`, `DASHBOARD_INFORMATION_HIERARCHY`, `UX_ACTION_SAFETY…` wortgleich | **PASS** |
| CDS Runtime-Abhängigkeit | **keine** | **PASS** |
| CDS Pilot | `Inactive / not activated`; `CDS-WP-017` inaktiv/nicht autorisiert | **PASS** |
| Semantic Status Candidate | ausschließlich als **read-only Vergleichseingabe** dokumentiert | **PASS** |
| CDS-Tokens/-Pakete importiert | **keine** | **PASS** |
| CoreOps-Conformance-Claims gegenüber CDS | **keine** | **PASS** |
| Lokale Dokumente, die aktive Adoption behaupten | **keine gefunden** | **PASS** |

**CDS boundary: preserved.** Ein frischer externer CDS-Review war damit **nicht** erforderlich. Der CDS-Reife-Re-Check bleibt vor jeder substanziellen Designübernahme erforderlich und wurde durch dieses WP **nicht** erfüllt.

### 21.3 Übriges Ökosystem

Keine verpflichtende Laufzeitabhängigkeit auf Core Vision, Core Brain, Core-Dev, NDF, MCP, CDS oder externe Managementprodukte. `SOVEREIGNTY_AND_DEPENDENCY_POLICY` und `DEC-S-02` (`prohibited`) sind durchgängig eingehalten. **Kein impliziter Autoritätstransfer** gefunden.

**standalone-first · self-hosted · offline-/air-gap-fähig · public-interface-first: alle vier bestätigt.**

---

## 22. CO-WP-030 Handoff

### 22.1 READINESS BLOCKER

**Keine.** Dieses Review hat **keinen** Befund identifiziert, der eine belastbare Foundation-Readiness-Prüfung verhindert.

### 22.2 HUMAN DECISION REQUIRED

**Präzise Einordnung.** Die folgenden vierzehn Punkte sind **Human-Maintainer-Entscheidungen bzw. -Inputs, die vor einem *positiven* Foundation-Readiness-Verdikt aufzulösen sind** — sie sind **keine** Vorbedingungen für den *Beginn* von `CO-WP-030`. Binding:

```
offener HM-Input     !=  CO-WP-030 darf nicht starten
offener HM-Input     !=  READINESS BLOCKER
HM-Input aufgeloest  !=  positives Readiness-Verdikt
```

**`CO-WP-030` darf mit sämtlichen dieser Punkte unaufgelöst starten** und muss sie dann selbst als jeweils zutreffend klassifizieren: `HUMAN DECISION REQUIRED` · `READINESS BLOCKER` · `READINESS NOTE` · `POST-FOUNDATION / NON-BLOCKING`. **`CO-WP-029` entscheidet das Readiness-Ergebnis nicht** und stuft keinen dieser Punkte als Blocker ein.

**Vor Beginn von `CO-WP-030` zwingend erforderlich: keiner der vierzehn Punkte.** Jeder ist während oder nach dem Readiness Review auflösbar. `HM-1`, `HM-2` und `HM-3` sind lediglich diejenigen mit der stärksten Wirkung auf ein *positives* Verdikt, weil an ihnen die Release-Taxonomie, die Delivery Baseline und die Auslegung des ADR-Exit-Gates hängen.

| # | Entscheidung | Bezug | Warum HM |
|---|---|---|---|
| **HM-1** | `DEC-A-0032` / `DEC-O-02` — Release-Taxonomie verbindlich machen oder erneut öffnen | §13.2 | Der Foundation-Tag-Kandidat `v0.0.1-foundation` hängt daran |
| **HM-2** | `DEC-A-0031` / `DEC-O-10` — Delivery Baseline (Docker-first) verbindlich machen | §13.2 | Grundlage mehrerer Foundation-Aussagen und des `RISK-20`-Treatments |
| **HM-3** | Definition von „**relevante ADRs**" für das Foundation Exit Gate | §13.1 | Das Gate ist ohne diese Definition nicht prüfbar |
| **HM-4** | Verbindlichkeit der zwölf `proposed` / `proposed-binding-governance`-Einträge (`DEC-O-02, -03, -04, -10, -11, -12, -16, -17, -18, -19, -20, -21`) nach erfolgtem Commit | §8.6 | Statushärtung ist Entscheidung, nicht Reconciliation |
| **HM-5** | Umgang mit den sechs offenen Konflikten `CCR-01`, `CCR-05`, `CCR-06`, `CCR-07`, `CCR-08`, `CCR-09` | §13.3 | Darf ein Foundation-Release mit sechs offenen `CCR` als konsistent gelten? |
| **HM-6** | ADR-Kandidaten-IDs für `DEC-O-05`, `DEC-O-07`, `DEC-D-03`, `DEC-D-06` vergeben — oder begründet darauf verzichten | §10.2 / MAJ-3 | Nummernvergabe ist HM-Sache |
| **HM-7** | Severity-Kalibrierungsregel für das Risk Register festlegen (Likelihood × Impact → Level), **danach** die 24 Abweichler prüfen | §15 / MAJ-6 | Neue Regel = neue Entscheidung |
| **HM-8** | `RISK-66` — Schließung, nachdem die geforderte Behandlung mit diesem WP ausgeführt ist | §14.4 / MIN-11 | Kein Risiko schließt ohne HM |
| **HM-9** | Target-WP-Review für `RISK-191`, `RISK-243`, `RISK-249`, `RISK-275` (auf `CO-WP-029` gerichtet, aber nur implementierungsseitig behandelbar) | §14.5 | Retargetierung ist Planungsentscheidung |
| **HM-10** | Decision-Konventionsharmonisierung `DEC-S-01…37` — Migration, Legacy-Deklaration oder bewusste Beibehaltung | §8.4 / MAJ-4 | Massenmigration untersagt ohne Entscheidung |
| **HM-11** | `ADR Required`-Dimension für `DEC-S-38…373` nachrüsten oder das Fehlen bewusst deklarieren | §10 / MAJ-2 | Schemaänderung am autoritativen Register |
| **HM-12** | Umgang mit den ~30 „Technologie deferred"-Decisions (`DEC-S-52…317`) unter der ab `CO-WP-024` geltenden Regel | §8.5 / MAJ-5 | Massendemotion untersagt ohne Entscheidung |
| **HM-13** | `ROADMAP.md` — stale ITIL/PRINCE2-Aussage korrigieren (forbidden file in diesem WP) | MIN-12 | Datei außerhalb des `CO-WP-029`-Änderungsrahmens |
| **HM-14** | Zielauflösung des uneindeutigen Verweises in `OFFLINE_TRUST_ACTIVATION…` §35 | MIN-6 | Zwei fachlich vertretbare Kandidaten |

### 22.3 READINESS NOTE

1. **Sieben Konsolidierungs-Cluster** (§12) reduzieren 15 ADR-Kandidaten auf sieben Entscheidungen. Für `CO-WP-030` ist das der wirksamste Hebel, das Exit Gate „relevante ADRs entschieden" beherrschbar zu machen.
2. **Drei Kandidaten** (`ADR-0001`, `ADR-0022`, `ADR-0029`) sind faktisch reine Governance-Fragen mit bereits akzeptierter Entscheidung — Umklassifizierungskandidaten, **keine** Entwertung.
3. **Die 21 auf `CO-WP-028` gerichteten Risiken** sind teststrategieseitig vollständig abgebildet, aber **nicht** behandelt. `test strategy documented ≠ tests executed`. Sie bleiben `treatment-planned`.
4. **Dokumentationsökonomie** (§20): vier parallele Current-State-Narrativen sind die nachgewiesene Ursache der Spiegel-Drift. Konsolidierung ist der dauerhafte Fix.
5. **24 Foundation Exit Gates** aus dem Scope Lock sind unverändert **offen** und **keines** als erfüllt markiert.
6. `PROJECT_PROFILE.md` §§1–17 und `PROJECT_BRAIN.md` §Bekannte Risiken/§Offene Fragen tragen `CO-WP-001`/`CO-WP-002`-Stand (§20) — chirurgischer Korrekturkandidat.

### 22.4 POST-FOUNDATION / NON-BLOCKING

- `ADR-0019` (Workflows), `ADR-0027` (PostgreSQL/Graph-DB), `ADR-0028` (Grafana) — `deferred-post-foundation`; **keine** Technologiewahl in Foundation 0.1.
- Die vier incident-abgeleiteten echten Erweiterungen (§19) — spätere Design-WPs.
- Reporting-/Vulnerability-Roadmap — `roadmap-candidate`, ohne WP-Nummern.
- `DEC-D-01`…`DEC-D-06` — sämtlich `post-Foundation`.
- Externe NDF-Release-Zuordnung der Adoption-Commits.

---

## 23. Claim Boundaries and Limitations

**Was dieses Review belegt:**

- Die mechanische Integrität von Decision Index und Risk Register zum Baseline-Stand `b7827b8` (ID-Eindeutigkeit, Lückenlosigkeit, Tally-Arithmetik) — deterministisch nachgerechnet.
- Die Auflösbarkeit aller relativen Verweise in Foundation-Dokumenten gegen den Dateibaum.
- Die Übereinstimmung der geprüften Zähl-Claims mit den autoritativen Quellen.
- Die Abwesenheit von Parallelmodellen und konkurrierenden Autoritätsmodellen auf der geprüften Ebene.

**Was dieses Review ausdrücklich NICHT belegt:**

```
Konsistenzreview durchgeführt   ≠  Foundation vollständig
Konsistenzreview durchgeführt   ≠  Foundation korrekt
Konsistenzreview durchgeführt   ≠  Foundation validiert
Konsistenzreview durchgeführt   ≠  Foundation releasebereit
Dokumentierte Semantik          ≠  implementierte Kontrolle
Dokumentierte Semantik          ≠  ausgewählte Architektur
ADR-Kandidat dispositioniert    ≠  ADR entschieden
Kein Widerspruch gefunden       ≠  kein Widerspruch vorhanden
Mechanische Prüfung bestanden   ≠  fachliche Richtigkeit bewiesen
```

**Weitere Grenzen:**

- **Evidenzstärke: `moderate`.** Die Zähl-, ID- und Referenzprüfungen sind deterministisch und reproduzierbar (`strong`). Die Aussagen zur inhaltlichen Kohärenz und zur Abwesenheit von Parallelmodellen beruhen auf **dokumentarischer Prüfung** durch eine einzelne Reviewinstanz, nicht auf unabhängiger externer Validierung (`limited`).
- Die **Modal-Severity-Abbildung** in §15 ist eine **Ableitung dieses Reviews**, keine Repository-Norm. Sie darf **nicht** als etablierte Kalibrierungsregel behandelt werden.
- Die Cluster-Empfehlungen in §12 sind **Vorschläge**. Sie führen nichts zusammen und nummerieren nichts um.
- Nicht jede Zeile aller 105 Dokumente wurde vollständig gelesen. Die Prüftiefe war **mechanisch vollständig** (alle IDs, alle Tabellen, alle Verweise, alle Zähl-Claims) und **selektiv-vertieft** bei Auffälligkeiten. Ein unauffälliger inhaltlicher Widerspruch tief im Fließtext eines nicht vertieft gelesenen Dokuments ist damit nicht ausgeschlossen.
- **Keine** Testausführung, **keine** Laufzeitprüfung, **keine** Sicherheitsverifikation, **keine** Accessibility-Prüfung, **keine** externe Validierung.

---

## 24. Final Review Verdict

**`review-complete-with-findings-awaiting-nova`**

**Begründung.** Die Foundation ist über 105 Dokumente und 26.245 Zeilen **inhaltlich konsistent**: 0 materielle Widersprüche, 0 Parallelmodelle, 0 konkurrierende Autoritätsmodelle, 0 BLOCKER. Die Autoritäts-, Zustands-, Integrations-, Deployment-, Recovery-, Evidence-, UX- und Test-Grenzen halten durchgängig und werden referenziert statt dupliziert. Register-Integrität: **einwandfrei** — Decision Index 373 lückenlos und duplikatfrei, Risk Register 316 lückenlos und duplikatfrei mit exakt stimmender Tally-Arithmetik in beiden Dimensionen.

Die sieben MAJOR-Befunde betreffen **ausnahmslos die Struktur des Decision Index, die Severity-Kalibrierungsregel und eine selbstwidersprüchliche Zählangabe** — nicht den Inhalt der Foundation-Modelle. Zwei davon waren mechanisch eindeutig und wurden korrigiert; fünf erfordern Human-Maintainer-Entscheidungen und gehen an `CO-WP-030`.

Der wichtigste Einzelbefund ist **MAJ-2**: für 336 der 373 `DEC-S`-Einträge ist der ADR-Bedarf **aus dem Schema nicht ableitbar** — weder bejahend noch verneinend. Das Exit Gate „relevante ADRs durch Human Maintainer entschieden" lässt sich daher derzeit **nicht** mechanisch gegen die `DEC-S`-Tabellen prüfen. Es bleibt jedoch **vollständig prüfbar** gegen die eigentliche Readiness-Basis (§10.3): das autoritative `DEC-A`-Inventar, die Decision-Zuordnungen der Tabellen mit `ADR Required`-Spalte, die Dispositionsmatrix aus §11 und die expliziten Human-Maintainer-Entscheidungen. `MAJ-2` ist damit eine **Prüfbarkeitslücke im Schema**, kein Defekt der ADR-Grundlage — und **kein** Readiness-Blocker.

Alle 32 ADR-Kandidaten sind inventarisiert und dispositioniert; **18** haben etablierte Foundation-Semantik bei offener Technologiewahl, **6** bleiben genuin offen (überwiegend an offene `CCR` gebunden), **3** sind post-Foundation, **3** sind faktisch bereits entschiedene Governance-Fragen und **2** erfordern eine Human-Maintainer-Entscheidung vor der Readiness-Prüfung. **Kein ADR wurde akzeptiert oder abgelehnt, keine ADR-Datei erzeugt, keine Technologie ausgewählt.**

**Nova entscheidet über GO / GO WITH NOTES / REWORK / STOP. Dieses Review erteilt keine Freigabe und trifft kein Readiness-Urteil.**

---

---

## 25. Nova Notes Closure (1–3)

Nova Review: **`GO WITH NOTES`**. Die drei verbindlichen Notes sind geschlossen. Der Foundation-Review selbst wurde **nicht** erneut durchgeführt; es wurden **keine** Decisions, Risks oder ADRs ergänzt und **keine** bereits akzeptierte Korrektur zurückgenommen oder ausgeweitet.

### Note 1 — ADR-Readiness-Basis präzisiert

Neuer **§10.3** trennt **Ableitbarkeit** von **Bedarf** und benennt die readiness-relevante ADR-Basis abschliessend:

```
fehlende ADR-Required-Spalte an einer DEC-S-Zeile   !=  ADR nicht erforderlich
fehlende ADR-Required-Spalte an einer DEC-S-Zeile   !=  ADR erforderlich
DEC-S Lifecycle Status / Decision Class             !=  ADR-Disposition
Ableitbarkeit aus dem Schema                        !=  ADR-Bedarf
```

Die Basis besteht aus vier Bestandteilen: autoritatives `DEC-A`-Kandidateninventar (32) · autoritative Source-/Decision-Zuordnungen der Tabellen **mit** `ADR Required`-Spalte · die CO-WP-029-Dispositionsmatrix (§11, 32/32) · explizite Human-Maintainer-Entscheidungen. Sie ist **vollständig**; `MAJ-2` ist dadurch als **Prüfbarkeitslücke im Schema** eingeordnet, **nicht** als Lücke der ADR-Grundlage — und ausdrücklich **kein** Readiness-Blocker (§5.1, §24 entsprechend nachgezogen).

Die Charakterisierung der vierzehn HM-Punkte ist präzisiert: **Entscheidungen bzw. Inputs, die vor einem *positiven* Foundation-Readiness-Verdikt aufzulösen sind** — **keine** Vorbedingung für den *Beginn* von `CO-WP-030` (§22.2). Ausdrücklich festgehalten: `CO-WP-030` darf mit sämtlichen Punkten unaufgelöst starten und klassifiziert sie dann selbst als `HUMAN DECISION REQUIRED` / `READINESS BLOCKER` / `READINESS NOTE` / `POST-FOUNDATION / NON-BLOCKING`; **vor Beginn zwingend erforderlich ist keiner der vierzehn**. `CO-WP-029` entscheidet das Readiness-Ergebnis nicht. Die Spaltenüberschrift der Dispositionsmatrix lautet entsprechend nicht mehr *HM vor `CO-WP-030`?*, sondern **HM-Input vor positivem Readiness-Verdikt?**

**Nicht getan:** keine `ADR Required`-Spalte an `DEC-S`-Zeilen ergänzt · keine Schema-Migration · keine Decision · kein ADR.

### Note 2 — Befund-Lifecycle vollständig

Alle **25** Befunde tragen jetzt eine explizite Disposition aus kontrolliertem Vokabular (`corrected` · `open-human-decision` · `open-readiness-review` · `deferred-post-foundation` · `note-only`). Die Spaltenüberschriften in §5.1/§5.2/§5.3 heissen **Disposition**; die NOTE-Tabelle hat die Spalte neu erhalten. Neuer **§5.4 Finding Lifecycle Ledger** listet alle 25 Befunde mit Severity, Disposition, nächster entscheidender Instanz und Readiness-Wirkung — damit sind die vier Fragen *korrigiert? · offen? · wer entscheidet? · readiness-wirksam?* je Befund deterministisch beantwortbar.

Severity und Lifecycle sind explizit als getrennte Dimensionen ausgewiesen:

```
finding severity          !=  finding lifecycle status
MAJOR                     !=  automatischer READINESS BLOCKER
an CO-WP-030 uebergeben   !=  stillschweigend akzeptiert
```

**Abstimmung — beide Summen gehen auf, ohne dass eine Severity verändert wurde:**

| | | |
|---|---|---|
| Severity | 0 BLOCKER · 7 MAJOR · 12 MINOR · 6 NOTE | **25** |
| Disposition | 7 `corrected` · 8 `open-human-decision` · 5 `open-readiness-review` · 1 `deferred-post-foundation` · 4 `note-only` | **25** |
| Korrigiert vs. weitergereicht | 7 / 18 | **25** |

Die Ist-Werte decken sich exakt mit den erwarteten (25 / 0 / 7 / 12 / 6 und 7 / 18) — **keine Abweichung zu erklären**. Eigentümerschaft der 18 weitergereichten Befunde: 8 Human Maintainer · 5 `CO-WP-030` · 1 spätere Design-WPs · 4 abschliessend bewertete Beobachtungen ohne Restaktion. **Kein Befund ist stillschweigend akzeptiert.**

### Note 3 — Severity-Heuristik als Heuristik gekennzeichnet

Neuer **§15.1** stellt der Kalibrierungsanalyse die Einordnung voran:

```
Modal-Abbildung dieses Reviews     =   rein analytische Heuristik
Modal-Abbildung dieses Reviews     !=  verbindliche Registernorm
Abweichung von der Modal-Abbildung !=  registrierter Risiko-Defekt
Abweichung von der Modal-Abbildung !=  nachgewiesene Fehlbewertung
24 abweichende Zeilen              =   Severity-Kalibrierungs-Reviewkandidaten
24 abweichende Zeilen              !=  24 fehlbewertete Risiken
```

Da **keine verbindliche** Likelihood × Impact → Level-Abbildung veröffentlicht ist, existiert **kein Massstab**, gegen den eine Zeile als falsch bewertet gelten könnte — und damit **keine verletzte Norm**. Die Tabellenspalte heisst nicht mehr *Abweichler*, sondern **Kandidatenzeilen**; §15.3 schliesst mit der ausdrücklichen Feststellung, dass das Register **keinen** Kalibrierungsdefekt aufweist und die 24 Zeilen ein **Arbeitsvorrat**, keine **Mängelliste** sind. Die Prosa enthält an keiner Stelle mehr die Behauptung bestätigter Fehlbewertungen.

**Register unverändert:** 0 Severity-Änderungen · 0 neue Risiken · keine Massen-Rekalibrierung · keine Zielverteilung · keine Quote. Am Risk Register wurde **keine Datei-Änderung** vorgenommen.

### Erhaltene Korrekturen (unverändert, nicht ausgeweitet)

`DEC-S-88` (3 falsche Break-Glass-Referenzen korrigiert, 1 korrekte Session-technology-Referenz erhalten) · `DEC-O-17` 74 → 94 Capabilities · Capability Matrix Domains 12 → 13 · Current-State-Baseline `b7827b8` · die beiden Verweiskorrekturen in `DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md` · die `CO-WP-024`-Current-State-Bereinigung.

### Weiterhin bewusst offen (nicht geraten, nicht angefasst)

`MIN-6` — der uneindeutige Verweis in `OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md` §35 mit zwei fachlich vertretbaren Zielen (`open-human-decision` → HM-14). `MIN-12` — die veraltete ITIL/PRINCE2-Aussage in `ROADMAP.md`, ausserhalb des Änderungsrahmens dieses WP (`open-human-decision` → HM-13). **Beide Dateien wurden in dieser Notes-Closure nicht verändert.**

### Ergebnis

**Nova Notes 1–3: geschlossen.** Status `CO-WP-029`: **`completed-go-with-notes`**; Human-Maintainer-Commit ausstehend. `CO-WP-030` und `CO-WP-031` bleiben `planned` / `retained` / nicht begonnen. **Keine automatische Folgeautorisierung.**

---

**Ende `CO-WP-029`. `CO-WP-030` ist nicht begonnen und nicht autorisiert.**
