# CoreOps – NDF Feedback Transfer Package 001

> Transfer Package Status: **Transferred to NDF Intake**
> NDF Intake Status: **Completed (NDF-INTAKE-COREOPS-001, Commit d08e35e)**
> NDF Adoption Status: **No candidate adopted**
> Framework Modification Status: **None**
> Erzeugt durch `CO-WP-004B1 – First NDF Feedback Transfer Package` (docs-only / transfer-preparation)

## 1. Status

Transferred to NDF Intake. This Transfer Package has been reviewed by the NDF team (NDF-INTAKE-COREOPS-001, Commit d08e35e). The intake review is completed; all seven candidates have been transferred to the NDF intake status.

This status confirms transfer to the NDF intake review; it does **not** claim adoption, implementation, or inclusion in any NDF version. Adoption decisions and implementation follow in later, independent NDF work packages.

Dieses Dokument ist ein **an den NDF-Intake übertragenes Transferpaket**. Die Intake-Überprüfung ist abgeschlossen; alle sieben Kandidaten wurden in den NDF-Intake-Status versetzt. Die tatsächliche Adoption erfolgt nur durch künftige eigenständige NDF-Work-Packages.

## 2. Purpose

Bündelung der sieben durch Nova als übergabeschwellen-relevant identifizierten NDF-Feedback-Kandidaten zu einem ersten, öffentlich neutralen Transferpaket, zur Vorbereitung eines späteren NDF Intake Review — ohne das NDF-Repository zu verändern.

## 3. Source Project

CoreOps (self-hosted, souveräne, BSI-orientierte Operations Control Plane; siehe [PROJECT_BRIEF.md](../architecture/PROJECT_BRIEF.md)).

## 4. Source Governance

- [LESSONS_LEARNED_PROCESS.md](LESSONS_LEARNED_PROCESS.md)
- [NDF_FEEDBACK_PROCESS.md](NDF_FEEDBACK_PROCESS.md)
- [project-system/LESSONS_LEARNED_REGISTER.md](../../project-system/LESSONS_LEARNED_REGISTER.md)
- [project-system/NDF_FEEDBACK_CANDIDATES.md](../../project-system/NDF_FEEDBACK_CANDIDATES.md)

## 5. Nova Authorization

```text
Die NDF-Übergabeschwelle ist mit sieben geeigneten Kandidaten erreicht.
```

Nova autorisiert die Vorbereitung dieses ersten Transferpakets für NDF-FC-COREOPS-001 bis -007. Diese Autorisierung bedeutet: Kandidaten dürfen für den Transfer vorbereitet und nach erfolgreicher Prüfung als `approved-for-transfer` vorgeschlagen werden. Sie sind **noch nicht** im NDF angekommen, **noch nicht** vom NDF akzeptiert, und dürfen **nicht** als `transferred-to-ndf` oder `adopted-in-ndf` gelten.

## 6. Human-Maintainer Gate

Der Human-Maintainer-Commit dieses Work Packages (`CO-WP-004B1`) bildet das **zweite Transfer-Gate** — es bestätigt die Transferfreigabe (`approved-for-transfer`), **nicht** den tatsächlichen NDF-Transfer. Der tatsächliche Transfer erfolgt ausschließlich über ein eigenständiges, außerhalb dieses Work Packages liegendes NDF-Work-Package mit eigenem Nova Review und eigener Human-Maintainer-Freigabe im NDF-Repository.

## 7. Included Candidates

Alle sieben reservierten Kandidaten sind enthalten: NDF-FC-COREOPS-001, -002, -003, -004, -005, -006, -007.

## 8. Bundle 1 – Work-Package Safety and Source Handling

| Candidate | Titel |
| --------- | ----- |
| NDF-FC-COREOPS-001 | Git Read versus Git Write in Work-Package-Prompts |
| NDF-FC-COREOPS-003 | Source-Handoff für Chat- und externe Dokumente |
| NDF-FC-COREOPS-004 | Blocked Report ohne künstlichen Zwischen-Commit |

**Generalisierte Beobachtung:** Work-Package-Prompts profitieren von expliziter Trennung read-only/schreibender Git-Operationen, einem verifizierbaren Source-Handoff-Verfahren für extern bereitgestellte Dokumente und einem strukturierten Blocked-Report-Format ohne künstlichen Zwischen-Commit. Alle drei Muster betreffen die Sicherheit und Nachvollziehbarkeit des Work-Package-Ablaufs selbst.

## 9. Bundle 2 – Skills Availability and Context Economy

| Candidate | Titel |
| --------- | ----- |
| NDF-FC-COREOPS-002 | Vollständige lokale Skills-Verfügbarkeit bei selektiver Aktivierung |

**Generalisierte Beobachtung:** Ein vollständiger, provenance-gesicherter lokaler Skills-Bestand kann bereitgestellt werden, ohne dass jedes Work Package alle Skills lädt — ein zweistufiges Modell aus vollständiger Verfügbarkeit und selektiver Aktivierung unterstützt sowohl Context Economy als auch Werkzeugvollständigkeit.

## 10. Bundle 3 – Governance and Status Modeling

| Candidate | Titel |
| --------- | ----- |
| NDF-FC-COREOPS-005 | Accepted Product Direction ohne technische Festlegung |
| NDF-FC-COREOPS-006 | Mehrdimensionale Statusmodelle statt überladener Einzelstatus |
| NDF-FC-COREOPS-007 | Framework-Tailoring vor Framework-Übernahme |

**Generalisierte Beobachtung:** Strategische, architekturelle und Rahmenwerk-bezogene Entscheidungen profitieren von differenzierten Statuswert-Familien (Richtung vs. Technik, Planung vs. Implementierung vs. Support, Kandidat vs. Übernahme), die frühe Klarheit ermöglichen, ohne verfrühte Festlegungen zu erzwingen.

## 11. Excluded Lessons

Die folgenden Lessons aus dem Register sind **nicht** Teil dieses Transferpakets, da sie als projektlokal (zu CoreOps-spezifisch für eine generalisierte NDF-Regel) eingestuft sind: LL-006, LL-007, LL-009, LL-010, LL-012, LL-014, LL-015. Diese bleiben im Lessons-Learned-Register mit ihrem jeweiligen Status und können bei künftiger Bündelung erneut geprüft werden.

## 12. Public-Neutrality Review

Alle sieben Kandidaten wurden auf private Namen, Pfade, Domains, Standorte, IP-Adressen, Secrets und interne Infrastrukturdetails geprüft: **PASS** für alle. Jeder Kandidat verwendet ausschließlich `Source Project`, `Source Work Package` und eine neutralisierte Evidenzbeschreibung (siehe [NDF_FEEDBACK_CANDIDATES.md](../../project-system/NDF_FEEDBACK_CANDIDATES.md)).

## 13. Security and Privacy Review

- NDF-FC-COREOPS-002 trägt `Security Relevance: yes` (Provenance/Hash-Lock als Supply-Chain-Kontrolle) — Inhalt beschreibt ein Verfahren, keine konkrete Schwachstelle; **PASS**.
- Keine der sieben Kandidaten enthält Schwachstelleninformationen, Secrets oder interne Zugangsdaten.
- Keine Chain-of-Thought, keine vollständigen Logs.

## 14. Existing-Rule Check Status

Dieses CoreOps-Work-Package führt **keine** verbindliche Prüfung gegen die aktuelle NDF-Arbeitsversion durch (kein Netzwerkzugriff, kein NDF-Repository-Zugriff). Es dokumentiert lediglich eine erste projektlokale Einschätzung. Die verbindliche Duplikat-/Existing-Rule-Prüfung erfolgt erst im späteren **NDF Intake Review**.

Besonders zu prüfen im NDF Intake Review:

```text
NDF-FC-COREOPS-002:
mögliche Überschneidung mit Skills-first und Context Economy
(ggf. bereits teilweise durch bestehende NDF-Skills-first-Prinzipien abgedeckt).

NDF-FC-COREOPS-006:
möglicherweise Guidance oder optionales Muster statt NDF-Core-Pflicht
(mehrdimensionale Statusmodelle könnten als empfohlenes Muster statt
verpflichtende Regel eingeordnet werden).
```

Für die übrigen fünf Kandidaten wurde projektlokal keine offensichtliche Überschneidung mit bekannten NDF-v1.0.0-Prinzipien identifiziert; dies ersetzt **nicht** die verbindliche NDF-seitige Prüfung.

## 15. NDF Intake Requirements

Für den späteren NDF Intake Review werden mindestens benötigt: Zugriff auf die aktuelle NDF-Arbeitsversion (ggf. `main`, informativ), ein NDF-seitiger Reviewer/Maintainer, eine Entscheidung je Kandidat (Annahme, Ablehnung, Änderung, Zusammenlegung), und — bei Annahme — ein eigenes NDF-Work-Package pro Bundle oder Kandidat gemäß NDF-eigenem Prozess.

## 16. Breaking-Change Assessment

Alle sieben Kandidaten sind bereits einzeln in [NDF_FEEDBACK_CANDIDATES.md](../../project-system/NDF_FEEDBACK_CANDIDATES.md) mit „Breaking-Change Potential: gering" (bzw. „gering bis mittel" bei NDF-FC-COREOPS-005) bewertet. Zusammenfassend für das Paket: Alle Vorschläge sind additiv (neue optionale Muster, Klarstellungen, Prozessergänzungen) und ändern keine bestehende NDF-Skill-Definition oder bestehendes Pflichtverhalten. Das Paket als Ganzes hat daher ein **geringes** Breaking-Change-Potenzial.

## 17. Proposed NDF Target Areas

- Bundle 1: NDF-Work-Package-Prompt-Vorlagen / `ndf-work-package-runner`-Dokumentation
- Bundle 2: NDF-Skills-Onboarding-Dokumentation / `ndf-skill-quality-reviewer`-Umfeld
- Bundle 3: NDF-Decision-Index-Vorlage, Capability-/Feature-Matrix-Vorlagen, Governance-/ADR-Vorbereitungsdokumentation

## 18. Adoption Is Not Yet Claimed

**Keiner** der sieben Kandidaten ist `transferred-to-ndf` oder `adopted-in-ndf`. Dieses Transferpaket ist ein **Vorschlag**; die tatsächliche Adoption erfordert ein eigenständiges NDF-Work-Package, NDF-seitigen Review und NDF-seitige Human-Maintainer-Freigabe außerhalb des CoreOps-Scopes.

## 19. Cross-Project Traceability

Jeder Kandidat referenziert `Source Project: CoreOps` und die jeweilige `Source Work Package`-ID (CO-WP-001, CO-WP-001A, CO-WP-002, CO-WP-004, CO-WP-004A), sodass eine spätere NDF-seitige Nachverfolgung ohne Offenlegung privater Projektinterna möglich ist.

## 20. Next Steps

1. Nova Review dieses Work Packages (`CO-WP-004B1`).
2. Human-Maintainer-Commit als zweites Transfer-Gate.
3. **NDF Intake Review** für Transfer Package 001 — außerhalb des CoreOps-Repositorys, mit eigenem NDF-seitigem Prozess.
4. Fortsetzung der CoreOps-Foundation-Queue mit `CO-WP-004C`.
