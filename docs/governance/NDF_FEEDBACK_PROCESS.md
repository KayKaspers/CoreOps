# CoreOps – NDF Feedback Process

> Status: **Proposed for acceptance** (bindend nach Human-Maintainer-Commit)
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004B` (docs-only / governance)

Dieser Prozess regelt, wie wiederverwendbare Erkenntnisse aus CoreOps kontrolliert als Kandidaten für das Nova Development Framework (NDF) geprüft werden — **ohne** das NDF-Repository direkt zu verändern.

## 1. Status

Proposed for acceptance.

## 2. Purpose

Kontrollierte, nachvollziehbare und öffentlich-neutrale Rückführung wiederverwendbarer Erkenntnisse aus CoreOps in das NDF, ohne unkontrollierte oder automatische Synchronisation.

## 3. CoreOps-to-NDF Boundary

CoreOps ist Anwenderprojekt des NDF, nicht dessen Maintainer-Repository. Kein CoreOps-Work-Package verändert das NDF-Repository, erzeugt ein NDF-Work-Package oder markiert einen Kandidaten als bereits übernommen. Jede Grenzüberschreitung erfolgt ausschließlich über ein **separates NDF-Work-Package** mit eigenem Nova Review und Human-Maintainer-Freigabe.

## 4. Candidate Creation

Ein NDF Feedback Candidate entsteht nur aus einer Lesson mit `Reusable Beyond CoreOps: yes`. Jeder Kandidat referenziert genau eine (oder mehrere gebündelte) Source Lesson(s) und erhält initial den Status `candidate`.

## 5. Candidate Validation

Vor jedem Statusfortschritt wird geprüft: nachvollziehbare Evidenz, generalisierbare (nicht CoreOps-spezifische) Formulierung, konkrete empfohlene NDF-Änderung. Unklare Kandidaten erhalten `validation-required`.

## 6. Public-Neutrality Review

Jeder Kandidat wird auf private Namen, Pfade, Domains, Standorte, IP-Adressen, Secrets und interne Infrastrukturdetails geprüft, bevor er über `candidate` hinaus fortschreitet. Nur `Source Project`, `Source Work Package` und eine neutralisierte Evidenzbeschreibung werden verwendet.

## 7. Security and Privacy Review

Sicherheitsrelevante Kandidaten (`Security Relevance: yes`) durchlaufen eine zusätzliche Prüfung auf unbeabsichtigte Offenlegung von Schutzmechanismen oder Schwachstelleninformationen, bevor eine Bündelung oder Eskalation erfolgt.

## 8. Duplicate and Existing-Rule Check

Vor `ready-for-bundling` wird geprüft, ob die NDF-Basis (Tag `v1.0.0`, Commit `9dcadc1`) die Regel bereits enthält oder ob ein gleichwertiger Kandidat bereits existiert. Bei Duplikat: Status `duplicate` mit Verweis.

## 9. Candidate Bundling

Kandidaten werden bevorzugt gebündelt statt einzeln transferiert (siehe Übergabeauslöser §15 unten). Ein `Suggested Bundle`-Feld gruppiert thematisch verwandte Kandidaten.

## 10. Transfer Readiness

Ein Kandidat gilt erst als transferreif, wenn **alle** folgenden Kriterien erfüllt sind:

1. die Quell-Lesson ist `validated`,
2. nachvollziehbare Evidenz liegt vor,
3. die Erkenntnis kann mehr als ein Projekt betreffen,
4. sie ist nicht bereits im NDF enthalten,
5. keine projektspezifischen Geheimnisse enthalten,
6. keine privaten Namen/Pfade/Infrastrukturdaten enthalten,
7. öffentlich neutral formuliert,
8. eine konkrete NDF-Änderung ist benennbar,
9. Breaking-Change-Potenzial ist bewertet,
10. ein NDF-Zielbereich ist benannt,
11. **Nova** hat die Transferreife bestätigt,
12. der **Human Maintainer** hat die Übergabe freigegeben.

Nur nach Erfüllung von (11) und (12) darf ein Kandidat auf `approved-for-transfer` gesetzt werden — **ausschließlich durch Nova Review und Human-Maintainer-Entscheidung**, niemals durch den Implementation Agent.

## 11. NDF Work Package Creation

Nach `approved-for-transfer` wird die tatsächliche Übernahme durch ein **eigenständiges NDF-Work-Package** (im NDF-Repository, außerhalb des CoreOps-Scopes) durchgeführt. Dieses Work Package liegt außerhalb des Scopes von `CO-WP-004B` und jedes CoreOps-Work-Packages.

## 12. Nova Review

Nova bewertet jeden Kandidaten vor Bündelung und vor Transferfreigabe (GO / GO WITH NOTES / REWORK / SPLIT / STOP), analog zum CoreOps-Work-Package-Review.

## 13. Human-Maintainer Gate

Ausschließlich der Human Maintainer entscheidet über: Transferfreigabe (`approved-for-transfer`), tatsächlichen Transfer (`transferred-to-ndf`) und Adoption-Bestätigung (`adopted-in-ndf`). Kein Implementation Agent darf diese Status setzen.

## 14. Adoption Tracking

Nach einem realen NDF-Transfer (außerhalb dieses Work Packages) wird der Kandidat mit `NDF Work Package`, `NDF Adoption Version` und `Backlink Status` aktualisiert. Bis dahin bleiben diese Felder leer/„—".

## 15. Rejection and Deferral

Ein Kandidat kann `rejected` (nicht geeignet, mit Begründung) oder `deferred` (später erneut prüfen) werden. Keine Löschung ohne Begründung.

## 16. Cross-Project Traceability

Jeder Kandidat referenziert `Source Project` (hier: CoreOps) und `Source Work Package`, sodass eine spätere NDF-seitige Nachverfolgung möglich ist, ohne private Projektinterna offenzulegen.

## 17. Emergency Security Feedback

Ein kritisches Security-Learning (z. B. eine fehlerhafte destructive Guardrail, die mehrere Projekte betrifft) darf **einzeln eskaliert** werden, statt auf eine Bündelung zu warten. Auch die Eskalation durchläuft weiterhin Nova Review und Human-Maintainer-Gate; es entfällt nur das Warten auf 5–10 gesammelte Kandidaten.

## 18. No Automatic Synchronization

Es gibt **keine** automatische Synchronisation zwischen CoreOps und NDF. Jede Übernahme ist ein bewusster, manuell freigegebener Schritt. Dieses Work Package selbst nimmt **keine** NDF-Änderung vor.

## Übergabeauslöser (Bundling Triggers)

Eine NDF-Übergabe wird geprüft, wenn mindestens einer der folgenden Punkte erfüllt ist: 5–10 geeignete Kandidaten gesammelt · ein Foundation-Meilenstein abgeschlossen · ein Kandidat betrifft mehrere aktive Projekte · ein kritisches Security-Learning liegt vor · eine fehlerhafte destructive Guardrail wurde erkannt · ein NDF-Widerspruch blockiert mehrere Projekte · Nova empfiehlt eine gebündelte Rückführung. Normale Kandidaten werden bevorzugt gebündelt; kritische Sicherheitskandidaten dürfen einzeln eskaliert werden (§17).
