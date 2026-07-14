# CoreOps – Lessons Learned Register

> Prozessreferenz: [LESSONS_LEARNED_PROCESS.md](../docs/governance/LESSONS_LEARNED_PROCESS.md)
> Erzeugt durch `CO-WP-004B` (docs-only / governance). Retrospektive Ersterfassung von `CO-WP-001` bis `CO-WP-004A` plus Lessons aus `CO-WP-004B` selbst.

Statuswerte: `observed` · `validated` · `project-action-planned` · `project-adopted` · `ndf-candidate` · `deferred` · `rejected` · `duplicate` · `superseded` · `closed`
Owner-Rollen: `Human Maintainer` · `Nova` · `Implementation Agent` · `Security Review` · `Project Governance`

## LL-001

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-001
- **Date or Phase:** Foundation 0.1
- **Title:** Git Read versus Git Write müssen in Prompts getrennt geregelt werden
- **Context:** CO-WP-001 untersagte pauschal alle Git-Befehle; zur File-Scope-Prüfung wurde dennoch `git status --porcelain` (read-only) verwendet.
- **Observation:** Ein pauschales Git-Verbot ohne Read/Write-Trennung führt zu vermeidbaren Abweichungsmeldungen bei harmlosen read-only-Prüfungen.
- **Impact:** Nova bewertete CO-WP-001 mit `GO WITH NOTES`; die Note betraf ausschließlich diese Unschärfe.
- **Contributing Factor or Root Cause:** Fehlende explizite Unterscheidung zwischen lesenden und schreibenden Git-Operationen im ursprünglichen Prompt.
- **Recommended Project Change:** Alle Folge-Work-Packages führen seither explizite „Erlaubte read-only-Git-Befehle" / „Verboten"-Listen (bereits ab CO-WP-001A umgesetzt).
- **Classification:** LL-PROMPT
- **Secondary Classifications:** LL-PROCESS, LL-NDF-CANDIDATE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** CO-WP-001 Rückmeldung an Nova, Abschnitt „Abweichungen"; Nova-Bewertung `GO WITH NOTES`.
- **NDF Candidate ID:** NDF-FC-COREOPS-001
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits umgesetzt ab CO-WP-001A
- **Notes:** Wiederkehrendes Muster in allen Folge-Prompts bestätigt die Wiederverwendbarkeit.

## LL-002

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-001A
- **Date or Phase:** Foundation 0.1
- **Title:** Vollständige lokale Skills-Verfügbarkeit erfordert nicht vollständige Aktivierung
- **Context:** Der komplette NDF-v1.0.0-Skills-Pack (38 Skills) wurde lokal bereitgestellt, ohne dass spätere Work Packages alle Skills laden müssen.
- **Observation:** Skills-first funktioniert als zweistufiges Modell: (1) vollständige lokale Verfügbarkeit, (2) selektive Auswahl pro Work Package.
- **Impact:** Reduziert Context-Verbrauch in jedem Folge-Work-Package, ohne Werkzeuge vorzuenthalten.
- **Contributing Factor or Root Cause:** Naive Skills-first-Umsetzungen könnten „alles laden" oder „nichts vorhalten" implizieren; beides ist suboptimal.
- **Recommended Project Change:** Muster in jedem Work-Package-Prompt als „Skills-first Operating Mode" mit Positivliste + Selektionskriterien fortführen.
- **Classification:** LL-TOOLING
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-PROCESS
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** CO-WP-001A Rückmeldung, Abschnitt „Skills-first-Nutzungsmodell"; Bestätigung „nicht alle Skills geladen" in allen Folge-Work-Packages.
- **NDF Candidate ID:** NDF-FC-COREOPS-002
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits durchgängig angewendet
- **Notes:** —

## LL-003

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-001A
- **Date or Phase:** Foundation 0.1
- **Title:** Skill-Provenance und Hash-Lock verbessern Nachvollziehbarkeit lokaler Framework-Skills
- **Context:** Der Skills-Import erforderte Quellenverifikation (Tag/Commit), byte-identische Übernahme und ein maschinenlesbares Lock-File mit SHA-256 je Datei.
- **Observation:** Ohne Provenance-Dokument und Lock-Datei wäre die Herkunft und Integrität der importierten Skills später nicht mehr nachprüfbar.
- **Impact:** `NDF_SKILLS_PROVENANCE.md` und `ndf-skills-lock.json` ermöglichen jederzeit eine Neuverifikation gegen die normative NDF-Quelle.
- **Contributing Factor or Root Cause:** Externe Quellenübernahme ohne Provenance-Tracking wäre ein Supply-Chain-Risiko gewesen.
- **Recommended Project Change:** Provenance + Lock-Datei bei jeder künftigen Skills-Aktualisierung erneut erzeugen.
- **Classification:** LL-SECURITY
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-TOOLING
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **Evidence:** CO-WP-001A Rückmeldung, Abschnitt „Provenance und Lock"; `project-system/NDF_SKILLS_PROVENANCE.md`, `project-system/ndf-skills-lock.json`.
- **NDF Candidate ID:** NDF-FC-COREOPS-002
- **Owner Role:** Security Review
- **Status:** validated
- **Follow-up Work Package:** Update-Regel bereits in NDF_SKILLS_PROVENANCE.md dokumentiert
- **Notes:** Gebündelt mit LL-002 unter demselben NDF-Kandidaten (Skills-first-Modell).

## LL-004

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-002
- **Date or Phase:** Foundation 0.1
- **Title:** Extern/im Chat bereitgestellte Quelldokumente benötigen expliziten Source-Handoff
- **Context:** CO-WP-002 wurde zunächst fail-closed blockiert, weil das vollständige Concept v3.0 nicht als verlässliche lokale Datei vorlag; erst ein expliziter Source-Handoff (lokaler Dateipfad) löste den Blocker.
- **Observation:** „Im Chat bereitgestellt" ist keine ausreichende Quellenbasis für eine vollständige Registrierung; ein referenzierbarer, verifizierbarer lokaler Handoff ist nötig.
- **Impact:** Vermeidet unvollständige oder erfundene Registrierungen umfangreicher externer Dokumente.
- **Contributing Factor or Root Cause:** Prompt- und Kontextgrenzen erlauben keine verlässliche Volltextübernahme rein aus Chatverlauf.
- **Recommended Project Change:** Bei umfangreichen externen Quelldokumenten stets einen expliziten, lokal verifizierbaren Source-Handoff verlangen, bevor eine „vollständige Registrierung" behauptet wird.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-NDF-CANDIDATE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** CO-WP-002 Blocked-Report; CO-WP-002-Fortsetzungsprompt „Source Preflight".
- **NDF Candidate ID:** NDF-FC-COREOPS-003
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits als Muster etabliert
- **Notes:** —

## LL-005

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-002
- **Date or Phase:** Foundation 0.1
- **Title:** Ein blockiertes Work Package ohne Änderungen benötigt keinen Zwischen-Commit, aber einen vollständigen Blocked Report
- **Context:** Nova bewertete den CO-WP-002-Blocker mit „GO – Blocker bestätigt" und bestätigte, dass kein Zwischen-Commit nötig war.
- **Observation:** Ein fail-closed-Stopp ohne Dateiänderung ist selbst ein vollständiges, bewertbares Ergebnis, wenn er strukturiert dokumentiert wird.
- **Impact:** Vermeidet unnötige Commits „für nichts" und erhält dennoch volle Nachvollziehbarkeit.
- **Contributing Factor or Root Cause:** Unklarheit, ob ein Blocker einen eigenen Commit-Zyklus benötigt.
- **Recommended Project Change:** Blocked-Report-Struktur (Grund, betroffene Prüfung, benötigte Entsperrung) als Standardmuster beibehalten.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-NDF-CANDIDATE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** CO-WP-002-Fortsetzungsprompt Abschnitt 1 „Nova Review des Blocked Reports"; „Es wurde kein Zwischen-Commit erzeugt und es ist kein Zwischen-Commit erforderlich."
- **NDF Candidate ID:** NDF-FC-COREOPS-004
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits als Muster etabliert
- **Notes:** —

## LL-006

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-003
- **Date or Phase:** Foundation 0.1
- **Title:** Produktmeilenstein, Foundation-Phase und Releaseversion müssen getrennt werden
- **Context:** Es bestand eine Kollision zwischen `Foundation 0.1` (interner Phasenname) und `Release 0.1 – Observe` (erster funktionaler Meilenstein), aufgelöst durch getrennte Tag-Kandidaten (`v0.0.1-foundation`, `v0.1.0-alpha.1`).
- **Observation:** Gleiche Nummerierung für Phase und Produktrelease erzeugt strukturelle Mehrdeutigkeit.
- **Impact:** Klare Trennung verhindert falsche Reifebehauptungen und Kommunikationsfehler.
- **Contributing Factor or Root Cause:** Ursprüngliches Concept nutzte „0.1" für zwei unterschiedliche Konzepte.
- **Recommended Project Change:** Release-Taxonomie-Dokument als verbindliche Referenz für alle künftigen Versions-/Phasenbenennungen nutzen.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** `docs/governance/RELEASE_TAXONOMY.md`; CO-WP-003 Rückmeldung Abschnitt „Konfliktauflösung" (CCR-02).
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits umgesetzt in RELEASE_TAXONOMY.md
- **Notes:** Als projektlokal eingestuft (kein eigener NDF-Kandidat vorgesehen), da Konzept eng an CoreOps-spezifische Roadmap-Struktur gebunden ist.

## LL-007

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-003
- **Date or Phase:** Foundation 0.1
- **Title:** Docker-first muss als Delivery-Anforderung von interner Anwendungsarchitektur getrennt werden
- **Context:** „Docker-first" wurde ausdrücklich als Delivery-/Betriebsanforderung eingeordnet (Compose-Standardinstallation), nicht als Festlegung auf eine bestimmte interne Architektur oder Kubernetes-Zwang.
- **Observation:** Eine Deployment-Präferenz kann fälschlich als Architekturentscheidung gelesen werden, wenn sie nicht ausdrücklich abgegrenzt wird.
- **Impact:** Verhindert vorschnelle, implizite Architekturfestlegungen durch eine reine Delivery-Aussage.
- **Contributing Factor or Root Cause:** Vermischung von „wie wird ausgeliefert" und „wie ist es intern gebaut" im ursprünglichen Concept.
- **Recommended Project Change:** Jede Delivery-/Betriebsanforderung ausdrücklich von Architekturentscheidungen abgrenzen (Muster in PROJECT_BRIEF.md fortgeführt).
- **Classification:** LL-ARCHITECTURE
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-PROCESS
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** CO-WP-003 Rückmeldung Abschnitt „Konfliktauflösung" (CCR-10); `docs/architecture/PROJECT_BRIEF.md` §8.
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits umgesetzt
- **Notes:** Projektlokal eingestuft; das allgemeine Muster „Delivery ≠ Architektur" ist in NDF-FC-COREOPS-006 (mehrdimensionale Statusmodelle) konzeptionell mit erfasst, aber nicht als eigener Kandidat dupliziert.

## LL-008

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004
- **Date or Phase:** Foundation 0.1
- **Title:** Roadmap-, Implementierungs- und Supportstatus müssen getrennte Dimensionen sein
- **Context:** Die Foundation Capability Matrix führte drei unabhängige Statusdimensionen (Roadmap Status, Implementation Status, Support Status) ein, statt eines überladenen Einzelstatus.
- **Observation:** Ein einzelner „Status" pro Capability hätte Planung, tatsächlichen Baufortschritt und Vertrauenswürdigkeit vermischt.
- **Impact:** Verhindert Fehlinterpretationen wie „geplant = unterstützt" oder „implementiert = supported".
- **Contributing Factor or Root Cause:** Naive Statusmodelle neigen zu Überladung eines einzigen Feldes.
- **Recommended Project Change:** Mehrdimensionale Statusmodelle als Standardmuster für jede künftige Capability-/Feature-Matrix verwenden.
- **Classification:** LL-ARCHITECTURE
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** `docs/architecture/FOUNDATION_CAPABILITY_MATRIX.md` Abschnitt „Drei getrennte Statusdimensionen"; CO-WP-004 Rückmeldung Abschnitt 7.
- **NDF Candidate ID:** NDF-FC-COREOPS-006
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits umgesetzt
- **Notes:** —

## LL-009

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004
- **Date or Phase:** Foundation 0.1
- **Title:** Herstellernennung oder Protokollunterstützung ist kein nachgewiesener Produktsupport
- **Context:** `CCR-12` wurde aufgelöst: Herstellernennung ist Priorisierungskandidat, kein Support-/Partnerschafts-/Zertifizierungsnachweis; ein 17-Punkte-Evidence-Satz ist vor jedem Supportversprechen nötig.
- **Observation:** Ohne explizite Trennung entsteht das Risiko falscher Kompatibilitäts- oder Supportbehauptungen gegenüber Nutzern.
- **Impact:** Schützt vor irreführenden öffentlichen Aussagen und unrealistischen Nutzererwartungen.
- **Contributing Factor or Root Cause:** Produktkonzepte listen häufig Zielhersteller, ohne den Unterschied zu „getestet und unterstützt" auszubuchstabieren.
- **Recommended Project Change:** Evidence-Satz-Muster (Modell, Firmware, CoreOps-Version, Capabilities, Tests, Herausgeber, Supportstatus) für jede künftige Integrationsdokumentation verbindlich vorschreiben.
- **Classification:** LL-DOCUMENTATION
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-SECURITY
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** `docs/integrations/INITIAL_SUPPORT_BOUNDARY.md` §6–8; CO-WP-004 Rückmeldung Abschnitt 7.
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits umgesetzt
- **Notes:** Eng mit dem in CO-WP-004A registrierten BSI-Claims-Muster (keine unbelegte Konformität behaupten) verwandt, aber als eigenständige Lesson erfasst, da sie primär den Herstellersupport betrifft.

## LL-010

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004
- **Date or Phase:** Foundation 0.1
- **Title:** Matrixzusammenfassungen müssen gegen die tatsächlichen Kategorien und Zeilenzahlen validiert werden
- **Context:** Bei der Erstellung der Foundation Capability Matrix wurde die abschließende Zusammenfassung (Anzahl Capabilities, Domains, Statusverteilung) gegen die tatsächlich geschriebenen Tabellenzeilen geprüft, um Diskrepanzen zu vermeiden.
- **Observation:** Handgeschriebene Zusammenfassungen zu großen Tabellen driften leicht von den tatsächlichen Zeilen ab, wenn sie nicht gegengeprüft werden.
- **Impact:** Verhindert falsche Zähl-/Kategorienangaben in Governance-Dokumenten.
- **Contributing Factor or Root Cause:** Große Tabellen (hier 74 Zeilen) werden oft parallel zur Zusammenfassung verfasst, ohne abschließenden Abgleich.
- **Recommended Project Change:** Nach Erstellung jeder umfangreichen Tabelle eine Zählprüfung (grep/count) gegen die Zusammenfassung durchführen, bevor das Dokument als abgeschlossen gemeldet wird.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-TOOLING
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** CO-WP-004 interne Validierung (Capability-Matrix-Zeilenzahl vs. Zusammenfassung).
- **NDF Candidate ID:** —
- **Owner Role:** Implementation Agent
- **Status:** validated
- **Follow-up Work Package:** bereits als Arbeitsweise etabliert
- **Notes:** Projektlokal/Tooling-Praxis; kein eigenständiger NDF-Kandidat, da es sich um eine allgemeine Arbeitsdisziplin und nicht um eine NDF-Regel handelt.

## LL-011

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004A
- **Date or Phase:** Foundation 0.1
- **Title:** Accepted Product Direction kann registriert werden, ohne eine technische Umsetzung vorwegzunehmen
- **Context:** Souveränität und BSI-orientierte Entwicklung wurden als „Accepted Product Direction" registriert, ohne Technologie, Architektur oder Zertifizierung zu behaupten.
- **Observation:** Ein eigener Statuswert für „akzeptierte Richtung" (getrennt von „akzeptierte Technik") erlaubt es, strategische Klarheit früh zu schaffen, ohne verfrühte technische Festlegungen zu erzwingen.
- **Impact:** Reduziert das Risiko, dass eine Produktrichtung fälschlich als technische oder Zertifizierungsentscheidung gelesen wird.
- **Contributing Factor or Root Cause:** Ohne diese Unterscheidung müsste jede strategische Aussage entweder verzögert oder überinterpretiert werden.
- **Recommended Project Change:** Statuswert-Familie „Accepted Product Direction" (mit klar getrennten Unterwerten wie `prohibited`, `not-claimed`, `foundation-candidate`) als wiederverwendbares Muster im Decision Index fortführen.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** `project-system/DECISION_INDEX.md` Abschnitt „Product Direction and Governance Decisions (CO-WP-004A)"; `docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md`.
- **NDF Candidate ID:** NDF-FC-COREOPS-005
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits umgesetzt
- **Notes:** —

## LL-012

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004A
- **Date or Phase:** Foundation 0.1
- **Title:** Produktsouveränität muss von vollständiger Eigenentwicklung abgegrenzt werden
- **Context:** Die Sovereignty-and-Dependency-Policy stellt klar, dass Unabhängigkeit von verpflichtenden externen Managementprodukten nicht bedeutet, jede Basisfunktion selbst neu zu entwickeln.
- **Observation:** Ohne diese Klarstellung besteht das Risiko von Over-Engineering aus einem missverstandenen Souveränitätsanspruch.
- **Impact:** Ermöglicht kontrollierte Nutzung etablierter Basiskomponenten, ohne die Kernunabhängigkeit zu gefährden.
- **Contributing Factor or Root Cause:** „Souverän" und „eigenständig" werden häufig fälschlich mit „alles selbst bauen" gleichgesetzt.
- **Recommended Project Change:** Jede künftige Souveränitäts- oder Unabhängigkeitsaussage durch eine explizite Dependency-Admission-Kriterienliste ergänzen (bereits in SOVEREIGNTY_AND_DEPENDENCY_POLICY.md umgesetzt).
- **Classification:** LL-ARCHITECTURE
- **Secondary Classifications:** LL-NDF-CANDIDATE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** `docs/architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md` §2, §5–6.
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** bereits umgesetzt
- **Notes:** Projektlokal eingestuft; das generelle Prinzip überschneidet sich teilweise mit NDF-FC-COREOPS-005, ist hier aber primär eine Architektur-/Scope-Erkenntnis.

## LL-013

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004A
- **Date or Phase:** Foundation 0.1
- **Title:** Mehrere Frameworks müssen zuerst als Tailoring-Kandidaten geführt werden, um Framework-Overload zu verhindern
- **Context:** BSI-Orientierung, ITIL und PRINCE2 wurden bewusst als Kandidaten (nicht als übernommene Rahmenwerke) registriert, mit Tailoring-Entscheidung erst in einem späteren Work Package.
- **Observation:** Die gleichzeitige volle Übernahme mehrerer umfangreicher Rahmenwerke ohne Tailoring würde die Governance schnell überladen.
- **Impact:** Hält die Foundation-Phase schlank, während die strategische Richtung bereits dokumentiert ist.
- **Contributing Factor or Root Cause:** Mehrere parallel diskutierte Rahmenwerke (BSI/IT-Grundschutz, ITIL, PRINCE2) ohne Priorisierung erzeugen Prozesslast.
- **Recommended Project Change:** Für jedes zusätzliche Rahmenwerk zunächst „Kandidat" statt „übernommen" registrieren; Tailoring-Entscheidung als eigenes Work Package (hier `CO-WP-004D`).
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-NDF-CANDIDATE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** `docs/architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md` §11–12; `project-system/RISK_REGISTER.md` RISK-32.
- **NDF Candidate ID:** NDF-FC-COREOPS-007
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** CO-WP-004D
- **Notes:** —

## LL-014

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004B
- **Date or Phase:** Foundation 0.1
- **Title:** Ein kontrollierter NDF-Feedback-Prozess benötigt getrennte Kandidaten- und Lesson-Statusmodelle
- **Context:** Bei der Erstellung dieses Work Packages zeigte sich, dass Lesson-Status (`observed`…`closed`) und NDF-Candidate-Status (`candidate`…`superseded`) unterschiedliche Lebenszyklen mit unterschiedlichen Freigabeinstanzen abbilden müssen.
- **Observation:** Eine Vermischung beider Statusmodelle hätte unklar gemacht, wer (Implementation Agent vs. Nova/Human Maintainer) welchen Status setzen darf.
- **Impact:** Klare Trennung verhindert, dass ein Implementation Agent versehentlich eine NDF-Transferfreigabe simuliert.
- **Contributing Factor or Root Cause:** Lessons und NDF-Kandidaten sind verwandte, aber nicht identische Konzepte (1:n-Beziehung, unterschiedliche Freigabegrenzen).
- **Recommended Project Change:** Getrennte Register (`LESSONS_LEARNED_REGISTER.md`, `NDF_FEEDBACK_CANDIDATES.md`) mit Cross-Referenzen beibehalten.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-NDF-CANDIDATE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **Evidence:** Struktur dieses Work Packages selbst (`LESSONS_LEARNED_REGISTER.md` vs. `NDF_FEEDBACK_CANDIDATES.md`).
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** ggf. spätere NDF-Kandidatenprüfung
- **Notes:** Neu erfasst in diesem Work Package selbst; noch nicht validiert durch Nova Review.

## LL-015

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004B
- **Date or Phase:** Foundation 0.1
- **Title:** Ein Human-Maintainer-Gate für Statuswerte muss technisch im Dokument erzwingbar dokumentiert werden
- **Context:** Der Prompt beschränkt explizit, welche NDF-Candidate-Statuswerte der Implementation Agent setzen darf (`candidate`…`duplicate`), und verbietet `approved-for-transfer` und höher.
- **Observation:** Eine reine Beschreibung „Nova/Human Maintainer entscheiden" ohne explizite Positivliste erlaubter Agent-Statuswerte wäre leichter zu übersehen gewesen.
- **Impact:** Die explizite Positivliste im Prozessdokument reduziert das Risiko einer versehentlichen Statusüberschreitung in künftigen Work Packages.
- **Contributing Factor or Root Cause:** Freigabegrenzen sind nur wirksam, wenn sie als konkrete erlaubte/verbotene Werte (nicht nur als Prinzip) dokumentiert sind.
- **Recommended Project Change:** Für jedes künftige mehrstufige Statusmodell mit Human-Gate eine explizite „Agent darf höchstens setzen"-Liste dokumentieren.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-NDF-CANDIDATE, LL-SECURITY
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **Evidence:** `docs/governance/NDF_FEEDBACK_PROCESS.md` §10, §13; dieser Prompt Abschnitt 13.
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** ggf. spätere NDF-Kandidatenprüfung
- **Notes:** Neu erfasst in diesem Work Package selbst; noch nicht validiert durch Nova Review.

## LL-016

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-004B2
- **Date or Phase:** Foundation 0.1
- **Title:** Commit-gated status transitions must be represented in repository state
- **Context:** A downstream NDF intake review was blocked because the Human Maintainer had committed an approval package (`CO-WP-004B1`, Commit 4ad3111) while the stored gate fields still said `pending until commit`. The commit provided evidence of approval intent, but the documentation state diverged from that intent.
- **Observation:** A commit provides evidence for a decision but does not automatically change status values stored in repository documents. The commitment of approval intent and the documentation of approval state are separate acts — both are required for completeness.
- **Impact:** The downstream intake correctly stopped fail-closed, preventing unauthorized continuation of a process based on contradictory documented state.
- **Contributing Factor or Root Cause:** The prior workflow treated the commit operation as if it could implicitly mutate documentation state within the same repository. This is not the case in a multi-step governance model with explicit state machines.
- **Recommended Project Change:** Approval fields must be set to the intended approved state BEFORE staging and committing. The Human-Maintainer commit then serves as external evidence and audit trail of that documented approval, but does not retroactively change status values.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-DOCUMENTATION, LL-SECURITY
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes, because ambiguous or contradictory gate state can cause unauthorized continuation (state validation bypass) or unnecessary blocking (false negatives in approval logic).
- **Evidence:** First NDF Intake attempt (`NDF-INTAKE-COREOPS-001`) blocked on Preflight with `Transfer Package Status: Prepared for Human-Maintainer Approval` and all seven `Human-Maintainer Gate: pending until commit` despite committed approval evidence. Status fields corrected in `CO-WP-004B2`.
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** validated
- **Follow-up Work Package:** NDF-seitige Governance-Überprüfung bei zukünftigen Intake Reviews
- **Notes:** Neue Lesson aus Blockierungserlebnis in `CO-WP-004B2` selbst; keine automatische NDF-Kandidatenförderung in diesem Work Package (nur projektlokale Dokumentation der Erkenntnis).

## LL-017

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-005…CO-WP-012 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Getrennte Entscheidungsdimensionen verhindern überladene Pseudo-Statuswerte
- **Context:** Ab CO-WP-005 wurden Entscheidungen durchgängig über drei getrennte Dimensionen (Decision Class · Lifecycle Status · Binding Level) statt über kombinierte Werte wie `binding-governance-direction` geführt.
- **Observation:** Die getrennten Dimensionen (DEC-S-38…136) blieben über acht Work Packages konsistent und eindeutig auswertbar, während die früheren kombinierten Werte (DEC-S-01…37) heute eine zweite, abweichende Konvention bilden.
- **Evidence:** Decision Index DEC-S-38…136 (getrennt) vs. DEC-S-01…37 (kombiniert); Milestone Review §12.
- **Impact:** Eindeutige Filterbarkeit und geringere Fehldeutung von Verbindlichkeit; zugleich Bedarf einer späteren Harmonisierung der Alt-Einträge.
- **Contributing Factor or Root Cause:** Kombinierte Statuswerte vermischen orthogonale Eigenschaften (Art, Reifegrad, Verbindlichkeit).
- **Recommended Project Change:** Getrennte Dimensionen als Standard beibehalten; Alt-Einträge in einer späteren Konsistenz-WP harmonisieren.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** partly (Verbindlichkeit von Sicherheitsentscheidungen muss eindeutig lesbar sein)
- **NDF Candidate ID:** NDF-FC-COREOPS-008
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Nova Review dieses Milestone Reviews; Konsistenz-WP (~CO-WP-029)
- **Notes:** Konsolidiert im Milestone Review CO-WP-005…012.

## LL-018

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-007…CO-WP-012 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Ein wiederverwendetes Vokabular an Sicherheitsinvarianten hält einen großen Dokumentensatz kohärent
- **Context:** Die 17 Threat-Model-Invarianten aus CO-WP-007 (§14) wurden als Designanforderungen wortgleich in Modul-, Identity-, Daten- und State-Dokumenten wiederverwendet.
- **Observation:** Die durchgängige „≠"-Kette (read≠write, unknown≠healthy, stale≠current, queued≠executed, executed≠successful, successful≠compliant, evidence-capability≠available≠satisfied) verhinderte konkurrierende Formulierungen über acht Work Packages.
- **Evidence:** CO-WP-007 §14; Wiederverwendung in CO-WP-008/009/010/011/012; Milestone Review §6, §10.
- **Impact:** Hohe konzeptionelle Konsistenz ohne Paralleldefinitionen; einfache Querverweisbarkeit.
- **Contributing Factor or Root Cause:** Ein einmal zentral definiertes, benanntes Invariantenset reduziert Neuformulierungsdrift.
- **Recommended Project Change:** Zentrales Invariantenvokabular pflegen und in Folgedokumenten referenzieren statt neu formulieren.
- **Classification:** LL-ARCHITECTURE
- **Secondary Classifications:** LL-SECURITY, LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **NDF Candidate ID:** NDF-FC-COREOPS-009
- **Owner Role:** Project Security / Governance
- **Status:** observed
- **Follow-up Work Package:** Nova Review dieses Milestone Reviews
- **Notes:** Konsolidiert im Milestone Review CO-WP-005…012.

## LL-019

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-007, CO-WP-006…CO-WP-012 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Trennung von Threat-Scenario-Register und Projekt-Risk-Register vermeidet Per-Threat-Duplikate, erfordert aber periodische Konsolidierung
- **Context:** Individuelle Bedrohungen werden im Threat Scenario Register (THR-001…040) geführt, während das Risk Register Governance-Risiken führt; einzelne Threats werden nicht als eigene Governance-Risiken dupliziert.
- **Observation:** Die Trennung verhinderte Duplikate, doch das Risk Register wuchs über die acht WPs auf 189 Einträge; mehrere späte Einträge formulieren nur eine bestehende Invariante um.
- **Evidence:** RISK-94…103 (Threat-Model-Governance-Hinweis); Risk-Register-Wachstum 66→189; Milestone Review §11.
- **Impact:** Klare Zuständigkeit je Registertyp, aber ansteigende Wartungslast des Risk Registers.
- **Contributing Factor or Root Cause:** Additive Registerpflege ohne Konsolidierungstakt führt zu Größenwachstum.
- **Recommended Project Change:** Registertrennung beibehalten; einen Konsolidierungs-/Indizierungslauf vor dem Foundation Readiness Review einplanen.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-SECURITY, LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** partly
- **NDF Candidate ID:** NDF-FC-COREOPS-010
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Risk-Register-Konsolidierung (~CO-WP-029/030)
- **Notes:** Konsolidiert im Milestone Review CO-WP-005…012.

## LL-020

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-006…CO-WP-012 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Das Muster „Modell + Standard/Register + Policy" mit additiven Registerupdates skaliert für Foundation-Governance
- **Context:** Jedes der WPs 006…012 lieferte zwei bis drei Dokumente (ein Modell/Kontext, ein Standard/Register, eine Policy) plus additive Updates an Decision Index, Risk Register, Queue und Context Pack.
- **Observation:** Dieses gleichförmige Ausgabemuster hielt den Scope je WP eng und die Querverweise stabil; die Allowed/Forbidden-File-Grenze blieb über acht WPs sauber.
- **Evidence:** Dokumentanzahl je WP; stabile IDs (CAP/MOD/THR/DEC-S/RISK); Milestone Review §5, §14.
- **Impact:** Vorhersehbare, review-freundliche Struktur; geringe Scope-Drift.
- **Contributing Factor or Root Cause:** Ein wiederholbares Dokument-Triplett-Muster reduziert Strukturentscheidungen je WP.
- **Recommended Project Change:** Muster für weitere Foundation-Governance-WPs beibehalten.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** —
- **Notes:** Konsolidiert im Milestone Review CO-WP-005…012; nicht als eigener NDF-Kandidat gefördert (Workflow-Muster, teils bereits durch bestehende NDF-Prozessführung abgedeckt).

## LL-021

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-005…CO-WP-012 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Milestone-Lessons-Konsolidierung nach fünf bis acht Work Packages ist die richtige Taktung
- **Context:** Jedes WP 005…012 hat Routine-Lessons bewusst zurückgestellt; die Konsolidierung erfolgte gebündelt in diesem Milestone Review.
- **Observation:** Die Zurückstellung hielt einzelne WPs schlank und vermied verfrühte NDF-Churn; nach acht WPs lagen genügend Muster für eine sinnvolle Konsolidierung vor.
- **Evidence:** „Lessons-Learned Register: unchanged" in den Rückmeldungen von CO-WP-005…012; dieser gebündelte Review.
- **Impact:** Geringerer Overhead je WP; höhere Signalqualität der konsolidierten Lessons.
- **Contributing Factor or Root Cause:** Lessons gewinnen an Aussagekraft, wenn sie über mehrere WPs gemustert werden.
- **Recommended Project Change:** Milestone-Lessons-Review-Takt (5–8 WPs) als Standard beibehalten.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** —
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **NDF Candidate ID:** — (teilweise durch bestehende NDF-Feedback-Bündelungsregel abgedeckt)
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** —
- **Notes:** Konsolidiert im Milestone Review CO-WP-005…012; nicht als eigener NDF-Kandidat gefördert wegen Überschneidung mit der bereits adoptierten Bündelungsregel.

## LL-022

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-005…CO-WP-012 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Governance-Register nähern sich einer Wartbarkeitsschwelle und brauchen einen Konsolidierungsschritt vor dem Readiness Review
- **Context:** Risk Register (189 Einträge) und Decision Index (DEC-S bis 136, plus DEC-P/-Alt-Einträge) sind über acht WPs stark gewachsen; zwei Decision-Status-Konventionen koexistieren.
- **Observation:** Ohne einen Konsolidierungs-/Indizierungsschritt steigt das Risiko von Duplikaten, uneinheitlichem Status (`treatment-planned` pauschal) und veralteten Count-Angaben (z. B. „74" statt 94 Capabilities in Alt-Abschnitten).
- **Evidence:** Milestone Review §11, §12, §13; RISK-66 (Capability-Count-Korrektur); PROJECT_BRAIN/Context-Alt-Abschnitte.
- **Impact:** Wartbarkeits- und Konsistenzrisiko vor dem Foundation Readiness Review.
- **Contributing Factor or Root Cause:** Rein additive Registerpflege ohne periodische Konsolidierung.
- **Recommended Project Change:** Dedizierten Konsistenz-/Konsolidierungsschritt (~CO-WP-029/030) vor Readiness Review vorsehen; als Follow-up dokumentiert, nicht still umgesetzt.
- **Classification:** LL-DOCUMENTATION
- **Secondary Classifications:** LL-PROCESS
- **Reusable Beyond CoreOps:** partly
- **Security Relevance:** no
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Konsistenz-/Konsolidierungs-WP (~CO-WP-029/030)
- **Notes:** Konsolidiert im Milestone Review CO-WP-005…012; projektlokal, kein NDF-Kandidat.

## LL-023

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-013…CO-WP-020 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Mehrdimensionale Trennung von Identity, Authority, Ownership, Lifecycle und Status skaliert domänenübergreifend
- **Context:** Acht Foundation-WPs trennten wiederholt orthogonale Dimensionen (Policy/Approval/Execution/Authorization in 013; Data Owner/Steward/Storage/Write/Migration/Retention/Recovery in 016; Source Trust vs. Authority in 018/020; Manual vs. Source/Execution Authority in 020).
- **Observation:** Die konsequente Trennung verhinderte Autoritäts-Overreach (z. B. Storage wird nicht Owner, Machine kann sich nicht selbst genehmigen, Manual imitiert keine beobachtete Wahrheit) und blieb über acht WPs konsistent.
- **Evidence:** Milestone Review §9; DEC-S-137/142/186…188/267; Cross-Document-Invarianten §7.
- **Impact:** Klare, auditierbare Autoritätsgrenzen als tragfähige Basis für die spätere Execution-/Deployment-Schicht.
- **Contributing Factor or Root Cause:** Kombinierte Autoritäts-/Status-Begriffe verbergen orthogonale Eigenschaften und ermöglichen Overreach.
- **Recommended Project Change:** Mehrdimensionale Trennung als Standard-Governance-Muster für alle Foundation-Domänen beibehalten.
- **Classification:** LL-ARCHITECTURE
- **Secondary Classifications:** LL-SECURITY, LL-PROCESS
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **NDF Candidate ID:** — (Muster über LL-017/NDF-FC-COREOPS-008 teilabgedeckt)
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Nova Review dieses Milestone Reviews
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020.

## LL-024

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-013, 014, 016, 017, 018 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** „Unknown outcome" benötigt eine durchgängige Governance über alle Ausführungs- und Datenpfade
- **Context:** Execution (013), Integration (014), Migration (016), API (017) und Audit (018) mussten jeweils den Fall behandeln, dass das Ergebnis einer Aktion unbekannt bleibt.
- **Observation:** Dasselbe Muster wiederholte sich: `unknown outcome ≠ failed ≠ not executed ≠ successful`; kein automatischer Retry; Reconciliation erforderlich; Side-Effect-Unsicherheit sichtbar. Ein einheitliches Muster verhinderte gefährliche Auto-Wiederholung.
- **Evidence:** Milestone Review §7, §10; Execution §17, Integration Trust §13/§14, Migration §21, API Error §14, Audit §22.
- **Impact:** Konsistente Fail-Safe-Behandlung unsicherer Ergebnisse; Grundlage für sichere Retry-/Reconciliation-Governance.
- **Contributing Factor or Root Cause:** Verteilte Systeme liefern häufig unbekannte statt eindeutige Ergebnisse; naive Gleichsetzung mit Erfolg/Fehler ist gefährlich.
- **Recommended Project Change:** Unknown-Outcome-Governance als übergreifendes Foundation-Muster kodifizieren.
- **Classification:** LL-SECURITY
- **Secondary Classifications:** LL-ARCHITECTURE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **NDF Candidate ID:** NDF-FC-COREOPS-011
- **Owner Role:** Project Security / Governance
- **Status:** observed
- **Follow-up Work Package:** Nova Review dieses Milestone Reviews
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020.

## LL-025

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-016, 018, 019, 020 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Source Trust ist zeitgebunden und von Source Authority getrennt; Observation Time ≠ Ingestion Time
- **Context:** Migration (016), Event/Audit (018), Telemetry (019) und Topology (020) mussten Quellen bewerten, die zeitweise vertrauenswürdig, später kompromittiert oder verzögert sein können.
- **Observation:** Source Trust wurde konsistent als zeitgebunden modelliert (Kompromittierung löscht historische Daten nicht, löst aber Neubewertung aus), getrennt von Source Authority (scopegebunden); Observation/Recording/Ingestion Time blieben getrennt (`recently ingested ≠ recently observed`).
- **Evidence:** Milestone Review §12; Topology Evidence §7-8, Event Time §11-12, Telemetry Freshness §23, Migration §17.
- **Impact:** Verhindert falsche Aktualitäts-/Autoritätsannahmen und stille Rückwirkung späterer Kompromittierung.
- **Contributing Factor or Root Cause:** Trust und Zeit werden leicht mit Autorität und Aktualität verwechselt.
- **Recommended Project Change:** Zeitgebundene Source-Trust-Neubewertung und Trennung der Zeitbegriffe als Standard beibehalten.
- **Classification:** LL-SECURITY
- **Secondary Classifications:** LL-ARCHITECTURE
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **NDF Candidate ID:** — (kandidatfähig; überschneidet sich teils mit NDF-FC-COREOPS-009, nicht in dieser Runde promoted)
- **Owner Role:** Project Security
- **Status:** observed
- **Follow-up Work Package:** ggf. spätere NDF-Kandidatenprüfung
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020.

## LL-026

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-018, 019, 020 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Die Evidence-Dimensionskette plus Source Independence bildet ein wiederverwendbares Sufficiency-Vokabular
- **Context:** Event/Evidence (018), Telemetry (019) und Topology (020) benötigten ein gemeinsames Vokabular, um Evidenz zu bewerten.
- **Observation:** Dieselbe Kette (capability ≠ available ≠ fresh ≠ integrity-verified ≠ validated ≠ sufficient) und `many items ≠ independent sources` wurde konsistent wiederverwendet; Sufficiency blieb decision-/scope-/zeitgebunden.
- **Evidence:** Milestone Review §12; Evidence Model §16-17, Topology Evidence §12-13, Telemetry Quality §20.
- **Impact:** Konsistente, nicht überzogene Evidenzaussagen über Domänen hinweg; Basis für Readiness-/Validierungsbewertung.
- **Contributing Factor or Root Cause:** Evidenz wird leicht überinterpretiert (verfügbar ⇒ gültig ⇒ ausreichend ⇒ compliant).
- **Recommended Project Change:** Evidence-Dimensionskette als gemeinsames Referenzvokabular pflegen.
- **Classification:** LL-ARCHITECTURE
- **Secondary Classifications:** LL-SECURITY, LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **NDF Candidate ID:** NDF-FC-COREOPS-012
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Nova Review dieses Milestone Reviews
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020; baut auf Evidence-Capability-Trennung (CO-WP-004E) auf.

## LL-027

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-020 (mit CO-WP-009/012) (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Nicht-destruktive Manual Authority mit Override und Suppression schützt konkurrierende Evidenz
- **Context:** Topology (020) führte Manual Authority/Override/Suppression ein, aufbauend auf Break Glass (009) und Divergence Exceptions (012).
- **Observation:** Manuelle Eingriffe blieben human-attributable, scope-bound, reviewbar und **nicht-destruktiv** — Overrides löschen keine konkurrierenden Observations/Evidence, Suppression bedeutet keine Abwesenheit, Machine Principals imitieren keine Manual Authority.
- **Evidence:** Milestone Review §14; Manual Authority Policy §6-11/§16, Break-Glass-Policy, Drift Exceptions §12.
- **Impact:** Menschliche Korrektur ohne Verlust von Beweislage/Audit; verhindert stille Realitätsverfälschung.
- **Contributing Factor or Root Cause:** Manuelle Autorität wird leicht als „Wahrheit" behandelt, die konkurrierende Daten überschreibt.
- **Recommended Project Change:** Nicht-destruktive Manual-Authority-Grenzen als Standardmuster für manuelle Korrekturen beibehalten.
- **Classification:** LL-SECURITY
- **Secondary Classifications:** LL-PROCESS
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** yes
- **NDF Candidate ID:** NDF-FC-COREOPS-013
- **Owner Role:** Project Security / Governance
- **Status:** observed
- **Follow-up Work Package:** Nova Review dieses Milestone Reviews
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020.

## LL-028

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-015, 016, 017, 019 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Support-, Compatibility- und Version-Claims müssen version-, target-, profil- und evidence-gebunden sein
- **Context:** Domain Pack (015), Schema/Migration (016), API (017) und Telemetry (019) definierten jeweils Compatibility-/Support-Aussagen.
- **Observation:** Konsistent galt: `same version ≠ same capability set`; `newer ≠ automatically compatible`; formal additiv kann semantisch breaking sein; Claims ohne Scope/Evidenz sind unzulässig; `supported ≠ validated`.
- **Evidence:** Milestone Review §11; Domain Pack Support §11-13, Schema Compatibility §8, API Versioning §7-12, Telemetry Compatibility §20.
- **Impact:** Verhindert Kompatibilitäts-Overclaim und blindes Vertrauen in Versionsgleichheit.
- **Contributing Factor or Root Cause:** Versionsnummern und Support-Labels werden leicht als Garantien missverstanden.
- **Recommended Project Change:** Claim-Bindung (version/target/profil/evidence) als Standard für alle Compatibility-/Support-Aussagen beibehalten.
- **Classification:** LL-PROCESS
- **Secondary Classifications:** LL-DOCUMENTATION
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** partly
- **NDF Candidate ID:** — (Muster über LL-017/018 und bestehende Kandidaten teilabgedeckt)
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Nova Review dieses Milestone Reviews
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020.

## LL-029

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-013…CO-WP-020 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Governance-Register erreichen die Wartbarkeitsschwelle und brauchen einen bestätigten Konsolidierungsschritt
- **Context:** Risk Register wuchs von 189 auf 279 (+90) und Decision Index auf DEC-S-273; die Zwei-Konventionen-Koexistenz (DEC-S-01…37 vs. 38…273) besteht fort.
- **Observation:** Wiederkehrende „X≠Y-Interpretation"-Risk-Familien und pauschaler Status `treatment-planned` erhöhen ohne Konsolidierung die Wartungslast; der Zuwachs wurde ab CO-WP-015 diszipliniert auf ≤10/WP gedeckelt, das Gesamtvolumen bleibt aber hoch.
- **Evidence:** Milestone Review §16, §17; Risk-Register-Verteilung (high 138/medium 117/low 24); Decision Index max DEC-S-273.
- **Impact:** Wartbarkeits-/Konsistenzrisiko vor dem Foundation Readiness Review; bestätigt und verstärkt LL-019/LL-022.
- **Contributing Factor or Root Cause:** Rein additive Registerpflege über 16 WPs ohne Konsolidierungslauf.
- **Recommended Project Change:** Konsolidierungs-/Indizierungslauf (Merge/Reclassify, Review-Daten, Konventionsharmonisierung) bei CO-WP-029/030 fest einplanen.
- **Classification:** LL-DOCUMENTATION
- **Secondary Classifications:** LL-PROCESS
- **Reusable Beyond CoreOps:** partly
- **Security Relevance:** no
- **NDF Candidate ID:** —
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Konsistenz-/Konsolidierungs-WP (~CO-WP-029/030)
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020; projektlokal, erweitert LL-022.

## LL-030

- **Source Project:** CoreOps
- **Source Work Package:** CO-WP-013…CO-WP-020 (Milestone Review)
- **Date or Phase:** Foundation 0.1
- **Title:** Ein gemeinsames Invarianten-/Template-Referenzmodell verbessert die Dokumentationsökonomie
- **Context:** 24 Foundation-Dokumente in gleichförmigem 3-Doc-Triplett wiederholen dieselben Invariantenvokabeln, Statusheader und Technology-Boundary-/Threat-References-Abschnitte.
- **Observation:** Die Gleichförmigkeit ist review-freundlich, erzeugt aber hohe Wiederholung und Token-Kosten; Begriffe (Provenance, Evidence, Confidence) werden mehrfach eingeführt.
- **Evidence:** Milestone Review §18; wiederholte „≠"-Ketten und Status-/Technology-Boundary-Abschnitte über CO-WP-013…020.
- **Impact:** Steigende Token-/Wartungskosten; ein gemeinsames Referenzdokument könnte Wiederholung senken, ohne Konsistenz zu verlieren.
- **Contributing Factor or Root Cause:** Jedes WP definiert Invarianten/Status/Boundaries eigenständig statt zu referenzieren.
- **Recommended Project Change:** Gemeinsames Invarianten-/Terminologie-Referenzdokument + Status-/Technology-Boundary-Template prüfen (spätestens Readiness CO-WP-030); bestehende Dokumente nicht rückwirkend kürzen.
- **Classification:** LL-DOCUMENTATION
- **Secondary Classifications:** LL-PROCESS
- **Reusable Beyond CoreOps:** yes
- **Security Relevance:** no
- **NDF Candidate ID:** — (projektlokaler Guardrail; spätere Kandidatenprüfung möglich)
- **Owner Role:** Project Governance
- **Status:** observed
- **Follow-up Work Package:** Dokumentationsökonomie-Prüfung (~CO-WP-030)
- **Notes:** Konsolidiert im Milestone Review CO-WP-013…020; erweitert LL-020/LL-021.

## Zusammenfassung

- **Anzahl Lessons:** 30 (LL-001…LL-013 retrospektiv, LL-014…LL-015 aus CO-WP-004B selbst, LL-016 aus CO-WP-004B2, LL-017…LL-022 aus dem Milestone Review CO-WP-005…012, LL-023…LL-030 aus dem Milestone Review CO-WP-013…020)
- **Verteilung nach Primärklasse:** LL-PROCESS 14 (004,005,006,010,011,013,014,015,016,017,019,020,021,028) · LL-PROMPT 1 (001) · LL-TOOLING 1 (002) · LL-SECURITY 4 (003,024,025,027) · LL-ARCHITECTURE 6 (007,008,012,018,023,026) · LL-DOCUMENTATION 4 (009,022,029,030)
  > Korrektur (CO-WP-004B1): Der vorherige Eintrag listete fälschlich 9 IDs für LL-PROCESS (inkl. „001"), obwohl „8" als Zahl genannt war. Ursache: LL-001 trägt `LL-PROCESS` nur als **Sekundärklasse**; ihre Primärklasse ist `LL-PROMPT`. Keine Lesson wurde umklassifiziert. Fortschreibung (Milestone Review 005…012): LL-016…LL-022. Fortschreibung (Milestone Review 013…020): LL-023…LL-030. Summe 14+1+1+4+6+4 = 30.
- **Verteilung nach Status:** `validated` 14 (LL-001…LL-013, LL-016) · `observed` 16 (LL-014, LL-015, LL-017…LL-030)
- **NDF-relevante Lessons (Reusable Beyond CoreOps: yes):** 28 (`yes`); LL-022, LL-029 = `partly`
- **NDF-Kandidaten aus den Milestone Reviews:** LL-017 → NDF-FC-COREOPS-008 · LL-018 → NDF-FC-COREOPS-009 · LL-019 → NDF-FC-COREOPS-010 · LL-024 → NDF-FC-COREOPS-011 · LL-026 → NDF-FC-COREOPS-012 · LL-027 → NDF-FC-COREOPS-013 (alle `candidate-pending-nova-review`)
- **Projektlokal ohne eigenen NDF-Kandidaten:** LL-006, LL-007, LL-009, LL-010, LL-012, LL-014, LL-015, LL-020, LL-021, LL-022, LL-023, LL-025, LL-028, LL-029, LL-030 (Muster bereits projektintern umgesetzt, durch bestehende NDF-Prozessführung/Kandidaten teilabgedeckt, kandidatfähig-aber-nicht-promoted bzw. zu spezifisch für einen eigenständigen Kandidaten)

**Bestätigung:** Keine Lesson wurde ohne lokale Evidenz erfunden. Keine Lesson ist bereits als NDF übernommen markiert.
