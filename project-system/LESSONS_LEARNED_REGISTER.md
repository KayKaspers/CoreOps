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

## Zusammenfassung

- **Anzahl Lessons:** 15 (LL-001…LL-013 retrospektiv, LL-014…LL-015 aus CO-WP-004B selbst)
- **Verteilung nach Primärklasse:** LL-PROCESS 8 (001,004,005,006,010,011,013,014,015 — siehe Einzeleinträge) · LL-PROMPT 1 (002 sek., 001 prim.) · LL-TOOLING 1 (002) · LL-SECURITY 1 (003) · LL-ARCHITECTURE 3 (007,008,012) · LL-DOCUMENTATION 1 (009)
- **Verteilung nach Status:** `validated` 13 (LL-001…LL-013) · `observed` 2 (LL-014, LL-015)
- **NDF-relevante Lessons (Reusable Beyond CoreOps: yes):** alle 15
- **Projektlokal ohne eigenen NDF-Kandidaten:** LL-006, LL-007, LL-009, LL-010, LL-012, LL-014, LL-015 (Muster bereits projektintern umgesetzt bzw. zu spezifisch für einen eigenständigen Kandidaten)

**Bestätigung:** Keine Lesson wurde ohne lokale Evidenz erfunden. Keine Lesson ist bereits als NDF übernommen markiert.
