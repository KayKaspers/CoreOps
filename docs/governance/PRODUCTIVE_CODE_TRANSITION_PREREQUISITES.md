# CoreOps – Productive-Code Transition Prerequisites

> Document Status: Implemented; Nova Review abgeschlossen (Initial `REWORK — narrow semantic closure`, beide blockierenden Notes CLOSED, **Nova Final Review `GO`**); `completed-go-with-notes`
> Human-Maintainer-Repository-Integration: COMPLETE — integrierter Commit `9999114200bf18baaadfb508e8464720b75e352e`; Push COMPLETE; Remote-Integration COMPLETE; `origin/main` = `9999114200bf18baaadfb508e8464720b75e352e`
> Artifact Class: Governance-Gate-Dokument (docs-only)
> Phase: `Observe` — betreten und mit Grenze autorisiert (`HM-O1` `APPROVED WITH BOUNDARY`)
> Productive Code: **NOT AUTHORIZED** · Implementation: **NOT AUTHORIZED** (Gate A, §16.1)
> `P-1` Target Authorization: **NOT AUTHORIZED** — Gate B (§16.2); **not** a Gate-A prerequisite
> `P-2` No-Mutation Evidence: method/plan **DEFINED** (Gate A) · satisfaction **NOT SATISFIED** — Gate B, downstream of implementation and execution
> `P-3` Collection Mechanism: **NOT SELECTED**
> Language / Runtime: **NOT SELECTED**
> Source Tree: **NOT CREATED** · Dependencies: **NONE ADMITTED** · Lockfile: **NONE**
> `README.md` / `LICENSE`: **NOT CREATED** — Empfehlung ist keine Artefaktautorisierung
> Work Package: `CO-WP-032`, primärer Typ `docs-only` — autorisiert durch `HM-O3` `APPROVED`; Ausführungsgrenze `HM-O4` `APPROVED WITH EXACT BOUNDARY` (docs-only)
> Successor Work Package: none created, none reserved
> Accepted ADRs: **0** — dieses Dokument akzeptiert keinen
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-032` (docs-only / Observe Slice Contract and Productive-Code Transition Prerequisites)

## 1. Status und Zweck

Dieses Dokument definiert das **Governance-Gate vor späterem produktivem Anwendungscode**. Es bereitet Entscheidungen **vor**; es trifft keine. Jede hier enthaltene Empfehlung ist ausdrücklich `PROPOSED / UNACCEPTED` und wird erst durch eine eigene, ausdrückliche Human-Maintainer-Entscheidung verbindlich.

```text
recommendation           != authorization
decision packet prepared != decision made
option evaluated         != option selected
gate documented          != gate passed
```

## 2. Foundation Relationship

**`Foundation 0.1` ist `CLOSED`** — förmlich geschlossen durch die Human-Maintainer-Entscheidung `HM-F1` (`APPROVED / COMPLETE`). Der Stand ist als annotierter, unsignierter Tag `v0.0.1-foundation` veröffentlicht; ein GitHub Release existiert nicht (`HM-R11` `NOT AUTHORIZED / CLOSED`).

Der [FOUNDATION_SCOPE_LOCK.md](FOUNDATION_SCOPE_LOCK.md) **governt und beschreibt die abgeschlossene Foundation-0.1-Phase** und bleibt für diese verbindlich. Er ist **keine dauerhafte Verbotsnorm für jede künftige produktive Arbeit**. Seine Liste *Forbidden Implementation Types* beschreibt, was **innerhalb der Foundation-Phase** unzulässig war — nicht, was für alle Zeit unzulässig bleibt.

Daraus folgt jedoch **nicht** das Gegenteil:

```text
Foundation phase closed  != productive code authorized
scope lock scoped to a closed phase != post-Foundation freedom
tag published            != functional product release
Observe entered          != implementation authorized
```

Post-Foundation-Arbeit — insbesondere produktiver Anwendungscode — benötigt eine **eigene, ausdrückliche Human-Maintainer-Autorisierung**. Eine solche ist **nicht erteilt**. Der Scope Lock selbst wird durch dieses Dokument **nicht** geändert; eine Änderung benötigte ein eigenes Work Package mit Nova Review und Human-Maintainer-Freigabe (Scope Lock, *Change Control*).

## 3. Observe Authority Boundary

`HM-O1` hat den Eintritt in die Phase `Observe` **mit Grenze** autorisiert. Der Eintritt in eine Phase ist eine Governance-Aussage, keine technische Freigabe.

| Durch Observe-Eintritt erteilt | Durch Observe-Eintritt **nicht** erteilt |
| ------------------------------ | ---------------------------------------- |
| Arbeit an Observe-Verträgen, Testdesign und Übergangs-Governance (docs-only) | Zielzugriff · reale Beobachtung · Testausführung |
| Auswahl **eines Wertslice** als Betrachtungsgegenstand (`HM-O2`) | Implementierung dieses Slice |
| Autorisierung **eines** Work Packages (`HM-O3`) mit docs-only Ausführungsgrenze (`HM-O4`) | produktiver Anwendungscode · Technologieauswahl · Nachfolge-Work-Package |

```text
Observe entered              != target access
value slice selected         != productive code authorized
CO-WP-032 authorized         != implementation authorized
work package execution authorized != productive code authorized
```

## 4. Implementation Authority

### 4.1 Zwei getrennte Gates

Produktiven Quellcode **zu schreiben** und ein **reales Ziel zu berühren** sind zwei verschiedene Autoritätsfragen mit zwei verschiedenen Gates. Sie werden hier ausdrücklich getrennt:

| Gate | Frage | Gegenstand |
| ---- | ----- | ---------- |
| **Gate A** | Darf produktiver Anwendungscode entstehen? | Implementierungs-/Produktivcode-Autorisierung (§16.1) |
| **Gate B** | Darf ein reales Ziel beobachtet werden? | Zielzugriff und Ausführung (§16.2) |

```text
productive code authorization != target authorization
implementation authorization  != target authorization
source code present           != target access granted
Gate A passed                 != Gate B passed
```

**`P-1` und die Erfüllung von `P-2` gehören zu Gate B, nicht zu Gate A.** Quellcode zu schreiben berührt kein Ziel. Ein implementierter, nie ausgeführter Beobachtungspfad greift auf nichts zu; Zielzugriff entsteht erst mit Ausführung gegen ein reales Ziel, und die dafür nötige Autorität wird eigenständig und später erteilt.

Für `P-2` gilt darüber hinaus zwingend: die No-Mutation-Evidenz **kann** vor der ersten Implementierung nicht vorliegen, weil sie einen ausführbaren Beobachtungspfad, Fixtures, eine Umgebung sowie getrennte Ziel- und Ausführungsautorität voraussetzt. `P-2` als Vorbedingung für Gate A zu führen wäre eine **zirkuläre Anforderung**. Auf Gate-A-Ebene ist deshalb ausschließlich verlangt:

```text
P-2 evidence method / plan DEFINED     <- maximal zulässige Gate-A-Anforderung
P-2 SATISFIED                          <- Gate B, nach Implementierung und Ausführung
P-2 evidence plan != P-2 satisfied
test design       != test execution
```

### 4.2 Gate A — Voraussetzungen der Implementierungsautorität

Implementierungsautorität ist **eigenständig** und entsteht nicht als Nebenwirkung. Sie setzt kumulativ voraus:

1. ein **ausdrücklich autorisiertes** implementierungsorientiertes Work Package mit benanntem Scope,
2. eine getroffene `P-3`-Entscheidung,
3. eine getroffene Sprach-/Runtime-Entscheidung,
4. eine getroffene Source-Tree-Entscheidung,
5. eine in Kraft gesetzte Dependency-Admission-Policy,
6. die Disposition von `NEW-8` (`README` / `LICENSE`),
7. die Disposition der erforderlichen ADR-/Decision-Punkte,
8. eine definierte Teststrategie beziehungsweise definierte künftige Validierungsanforderungen — einschließlich der **definierten `P-2`-Evidenzmethode**, ausdrücklich **nicht** deren Erfüllung,
9. die Build-/Packaging-Disposition, soweit dieses Dokument sie als tatsächlich erforderlich ausweist (§12).

Fehlt auch nur eine, ist **keine** Implementierungsautorität gegeben. Derzeit sind die neun Voraussetzungen **nicht kumulativ erfüllt**: mehrere stehen offen, und die Voraussetzung zu Teststrategie und `P-2`-Evidenzmethode (Punkt 8) ist lediglich **teilweise** erfüllt — Testdesign und Evidenzmethode sind definiert, die erforderliche Human-Maintainer-Disposition steht aus (Gate-A-Punkt `A-12` **teilweise**, §16.1).

**Ausdrücklich nicht Bestandteil von Gate A:** `P-1` · die Erfüllung von `P-2` · Testausführungsautorität · Fixture-Freigabe · Lab-Bereitstellung. Diese gehören zu Gate B (§16.2).

Unverändert gilt der Human-Maintainer-Vorbehalt ([Repository Governance Standard](REPOSITORY_GOVERNANCE_STANDARD.md) §6, `DEC-G-01`/`DEC-G-03`): Freigabe, Staging, Commit, Push, Merge, Branch, Tag, Release und jede irreversible Aktion liegen ausschließlich beim Human Maintainer.

## 5. `P-1` — Zielautorisierung (Gate B)

**Status: `NOT AUTHORIZED`.**

**Gate-Zuordnung: `P-1` ist eine Gate-B-Voraussetzung.** Sie ist **keine** Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren. `implementation authorization != target authorization`; Code, der nie gegen ein reales Ziel ausgeführt wird, berührt kein Ziel.

`P-1` ist die separate Human-Maintainer-Autorisierung für **das konkrete Ziel** und **die konkrete read-only Beobachtung**. Weder die Foundation, noch das Readiness Review, noch der Observe-Eintritt, noch der Observation Contract, noch dieses Dokument erteilt sie.

Eine spätere `P-1`-Erteilung muss **mindestens** benennen:

| Element | Anforderung |
| ------- | ----------- |
| **Konkretes Ziel** | welcher Host, in welcher Umgebung, in welcher Rolle; **keine** Zielklasse als Blankettautorisierung |
| **Konkrete Beobachtung** | welche Felder des [Observation Contract](../architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md), aus welchen Quellklassen, mit welcher Erhebungsmechanismusklasse |
| **Ausführungsgrenze** | read-only; keine Schreib-, Management- oder Ausführungsautorität; keine Rechteeskalation; keine Credentials; keine Secrets; kein Netzwerkziel; zeitliche Begrenzung; Widerrufbarkeit; Auditpflicht |

```text
target class authorized != every target authorized
one observation authorized != a standing observation right
P-1 granted != P-2 satisfied
P-1 granted != productive code authorized
```

Dieses Dokument **autorisiert kein Ziel** und benennt keines.

## 6. `P-2` — No-Mutation- und Ausführungssicherheits-Evidenz

### 6.1 Zwei getrennte Zustände

| Zustand | Gate | Stand heute |
| ------- | ---- | ----------- |
| **`P-2`-Evidenzmethode / -plan definiert** | Gate A (§4.2 Punkt 8) | **DEFINIERT** — die sieben kumulativen Anforderungen stehen im [Test Envelope](../testing/OBSERVE_LOCAL_LINUX_HOST_TEST_ENVELOPE.md) §8.3 |
| **`P-2` erfüllt** | Gate B | **`NOT SATISFIED`** |

```text
P-2 evidence plan != P-2 satisfied
```

### 6.2 Warum die Erfüllung erst nach Gate A möglich ist

Die `P-2`-Evidenz ist **beobachtetes geschütztes Verhalten**. Sie setzt voraus: einen ausführbaren Beobachtungspfad (also eine bereits erfolgte Implementierung), freigegebene Fixtures, eine deklarierte Umgebung, `P-1` und eine ausdrückliche Testausführungsautorität. Keine dieser Voraussetzungen kann vor der ersten Implementierung erfüllt sein.

`P-2` als Vorbedingung für die **Autorisierung** produktiven Codes zu führen wäre daher eine **zirkuläre Anforderung**: die Evidenz verlangt genau das, was sie angeblich freigeben soll. Die einzige auf Gate-A-Ebene zulässige Anforderung ist deshalb, dass die **Methode** feststeht — nicht, dass die Evidenz vorliegt.

```text
P-2 satisfaction is downstream of implementation and execution
defining how it will be proven != having proven it
test design                    != test execution
```

### 6.3 Inhalt der künftigen Evidenz

**Status der Erfüllung: `NOT SATISFIED`.**

Der No-Mutation-Vertrag muss durch **beobachtetes geschütztes Verhalten** belegt werden, nicht durch das Ausbleiben eines Fehlers. Die erforderliche künftige Evidenz ist vollständig im [Test Envelope](../testing/OBSERVE_LOCAL_LINUX_HOST_TEST_ENVELOPE.md) §8 beschrieben und wird hier **nicht** dupliziert. Kern:

```text
read-only                  != side-effect-free
no error observed          != protected behavior observed
"built as read-only"       != no mutation observed
one successful run         != regression confidence
not observable             != not present
```

Mindestens die Negativ-Familien `unknown target`, `wrong target` und `missing execution authorization` müssen als **implementierte und ausgeführte** Tests vorliegen ([Teststrategie](../testing/FOUNDATION_TEST_STRATEGY_AND_VALIDATION_MODEL.md) §12). Aktuell: **null Tests implementiert, null Tests ausgeführt, keine Evidenz vorhanden.**

**Zur Beschaffung dieser Evidenz wird in diesem Work Package nichts ausgeführt** — kein Zielzugriff, kein Test, kein Kommando, keine Beobachtung.

## 7. `P-3` — Decision Packet: lokaler Erhebungsmechanismus

**Status: `NOT SELECTED`.** Das Folgende ist ein **Entscheidungspaket** — Optionsvergleich und Empfehlung, ausdrücklich `PROPOSED / UNACCEPTED`.

### 7.1 Aktuelle Decision Ceiling

Das Paket bewegt sich ausschließlich innerhalb dieser Grenze:

```text
local Linux host only · no remote host · no network transport ·
no credentials · no secrets · no privileged execution
```

Optionen, die diese Grenze überschreiten, werden **nicht** bewertet, sondern ausgeschlossen.

### 7.2 Optionen (Mechanismusklassen, keine Produkte)

| Option | Beschreibung |
| ------ | ------------ |
| **A — Direktes Lesen standardisierter OS-Identitätsdateien** | Lesen der vom Betriebssystem bereitgestellten, standardisierten textuellen Identitäts- und Systeminformationsdateien im Dateisystem des lokalen Hosts |
| **B — Systeminformationsaufrufe der Standardbibliothek/Plattform** | In-Process-Abfrage der von der gewählten Sprachplattform bereitgestellten Standard-Systeminformationsschnittstellen |
| **C — Aufruf von OS-Kommandozeilenwerkzeugen und Parsen ihrer Ausgabe** | Starten mitgelieferter Systemwerkzeuge als Subprozess und Interpretieren der Textausgabe |
| **D — Abfrage eines lokalen System- oder Paketverwaltungsdienstes** | Nutzung einer lokalen Dienst-/Datenbankschnittstelle des Hosts als Informationsquelle |
| **E — Lokaler privilegierter Agent oder Daemon** | Ein dauerhaft laufender lokaler Sammler mit erhöhten Rechten |

### 7.3 Bewertung

Skala: **hoch** / **mittel** / **gering** — jeweils im Sinne der Spaltenüberschrift günstig zu lesen (bei Risiko- und Aufwandsspalten ist *gering* günstig).

| Kriterium | A — Dateien | B — Plattform-Aufrufe | C — Werkzeugaufruf | D — lokaler Dienst | E — Agent |
| --------- | ----------- | --------------------- | ------------------ | ------------------ | --------- |
| **Semantische Treue** | hoch — Quellsemantik bleibt sichtbar, Raw ist der Dateiinhalt | mittel — Plattform hat bereits interpretiert; Raw ist ein API-Rückgabewert | mittel — Ausgabe ist bereits formatiert und für Menschen gedacht | mittel bis gering — Dienstsicht ist ein weiteres Modell | gering — zusätzliche Sammler-Semantik |
| **Mutations-/Seiteneffektrisiko** | gering — lesender Dateizugriff; Rest-Seiteneffekte (Zugriffszeiten, Auditspuren) bleiben deklarationspflichtig | gering | mittel — Prozessstart, Ressourcenbindung, Logspuren | mittel — Dienstinteraktion, ggf. Zustandsänderung im Dienst | hoch — dauerhafte Präsenz, Konfigurations- und Zustandswirkung |
| **Privilegienbedarf** | gering — üblicherweise unprivilegiert lesbar | gering | mittel — Werkzeugverfügbarkeit und -rechte variieren | mittel bis hoch | **hoch — verletzt die Decision Ceiling** |
| **Portabilität (Linux)** | hoch für die standardisierten Identitätsquellen; einzelne Distributionen weichen ab | hoch | gering bis mittel — Werkzeugpräsenz, -version und -ausgabeformat variieren stark | gering — dienstabhängig, nicht überall vorhanden | mittel |
| **Provenance-Qualität** | hoch — Quelle ist eindeutig benennbar und referenzierbar | mittel — Quelle ist die Plattform, nicht die Ursprungsquelle | gering — Provenance endet beim Werkzeug, dessen eigene Quelle verborgen bleibt | mittel | gering |
| **Fehlerbeobachtbarkeit** | hoch — Abwesenheit, Verweigerung und Malformedness sind sauber trennbar | mittel — Fehlerklassen sind plattformabhängig zusammengefasst | gering — Exit-Code und Text vermischen Ursachen | mittel | gering |
| **Abhängigkeitslast** | gering — keine Drittabhängigkeit erforderlich | gering | mittel — Laufzeitabhängigkeit von externen Werkzeugen | hoch | hoch |
| **Offline-Eignung** | hoch | hoch | hoch | mittel | mittel |
| **Testbarkeit** | hoch — Quellen sind als Dateien fixture-fähig; alle zehn Envelope-Fälle sind darstellbar | mittel — Plattformaufrufe brauchen Abstraktion und Test Doubles | gering — Subprozessverhalten ist schwer deterministisch nachzustellen | gering | gering |
| **Implementierungskomplexität** | gering bis mittel — Parsing, aber klar begrenzt | gering | mittel | hoch | hoch |

### 7.4 Empfehlung

```text
RECOMMENDATION (P-3):  Option A als primärer Mechanismus,
                       Option B ergänzend, wo A keine Quelle bietet.
                       Option C nicht als Standardweg.
                       Option D und Option E für diesen Slice ausgeschlossen.
STATUS:                PROPOSED / UNACCEPTED
P-3:                   NOT SELECTED
```

**Begründung.** Option A liefert als einzige Option gleichzeitig die höchste Provenance-Qualität (die Quelle ist benennbar und der Raw-Wert ist der Quellinhalt selbst), die sauberste Fehlerbeobachtbarkeit (Abwesenheit, Verweigerung und Malformedness bleiben trennbar, was Abschnitt §10 des Observation Contract überhaupt erst prüfbar macht) und die höchste Testbarkeit (alle zehn Envelope-Fälle sind mit Datei-Fixtures darstellbar). Option B schließt Lücken dort, wo keine standardisierte Datei existiert, verliert dabei aber Provenance-Tiefe: die Plattform hat bereits interpretiert. Option C wird **nicht** als Standardweg empfohlen, weil Werkzeugausgaben für Menschen und nicht für Verträge formatiert sind, weil Exit-Code und Textausgabe die vier zu trennenden Fehlerursachen vermischen und weil ein Subprozessstart das Seiteneffektprofil unnötig vergrößert. Option D und E überschreiten Abhängigkeits- beziehungsweise Privilegiengrenze und sind für einen ersten read-only Slice nicht vertretbar; E verletzt die Decision Ceiling unmittelbar.

**Diese Empfehlung ist keine Auswahl.** Sie benennt keine konkreten Pfade, keine Bibliothek und kein Werkzeug. `P-3` bleibt **`NOT SELECTED`**, bis der Human Maintainer entscheidet.

## 8. Decision Packet: Sprache und Runtime

**Status: `NOT SELECTED`.** Es wird **keine** frühere Technologieentscheidung unterstellt: das Repository führt `Technische Entscheidungen: keine`, **0** akzeptierte ADRs und `Technische Architektur: unbestätigt`.

### 8.1 Randbedingungen aus bestehender Governance

- **Standalone-First / Souveränität** — der Kern muss ohne verpflichtende externe Plattform betreibbar sein ([Sovereignty and Dependency Policy](../architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md) §2–§4).
- **Offline-/Air-Gap-Fähigkeit** — jede zugelassene Basisabhängigkeit muss offline verfügbar sein (ebenda §7).
- **Docker-first ist Delivery Baseline** (`HM-2` `APPROVED WITH BOUNDARY`) — und ausdrücklich `Docker-first != Docker-only != zwingende interne Anwendungsarchitektur != zwingende Runtime-Abhängigkeit != Observe-Voraussetzung`. Docker-first **präjudiziert die Sprachwahl nicht**.
- **DE/EN als Produktsprachen**, Englisch kanonisch für maschinenbezogene Artefakte ([Language Standard](COREOPS_LANGUAGE_STANDARD.md) §6).

### 8.2 Kandidaten und Bewertung

Bewertet werden **Sprach-/Runtime-Klassen**, nicht Versionen oder Distributionen.

| Kriterium | Go | Rust | Python | TypeScript / Node.js | JVM (Java/Kotlin) | .NET (C#) |
| --------- | -- | ---- | ------ | -------------------- | ----------------- | --------- |
| **Standalone/self-hosted-Eignung** | hoch — einzelnes statisch gelinktes Binary | hoch — einzelnes Binary | gering — Interpreter plus Umgebung | gering bis mittel — Runtime plus Modulbaum | mittel — Runtime nötig; self-contained Images möglich | mittel — self-contained Publish möglich |
| **Offline-Betrieb** | hoch | hoch | mittel — Umgebungsauflösung offline aufwendig | mittel | mittel | mittel |
| **Minimaler Deployment-Footprint** | hoch | hoch | gering | gering | gering bis mittel | mittel |
| **Linux-Portabilität** | hoch — breite Architekturabdeckung, einfache Cross-Kompilierung | hoch | hoch (Quellebene), gering (Verteilung) | mittel | hoch | mittel bis hoch |
| **Sicherheitsoberfläche** | mittel bis hoch — speichersicher im Normalfall, kleine Laufzeitoberfläche | hoch — stärkste Speicher- und Nebenläufigkeitsgarantien | mittel — große Laufzeit- und Paketoberfläche | gering bis mittel — große Paketoberfläche | mittel | mittel |
| **Deterministische Paketierung** | hoch — Modul-Prüfsummen, Vendoring, reproduzierbare Builds gut erreichbar | hoch — Lockfile plus Vendoring | gering bis mittel | gering bis mittel | mittel | mittel |
| **Abhängigkeitslast für diesen Slice** | sehr gering — Standardbibliothek genügt plausibel | sehr gering bis gering | mittel | mittel bis hoch | mittel | mittel |
| **Wartbarkeit** | hoch — kleine Sprache, flache Lernkurve, stabile Kompatibilitätszusage | mittel — höhere Einstiegs- und Umbaukosten | hoch (Lesbarkeit), mittel (Betriebsstabilität) | mittel | mittel bis hoch | mittel bis hoch |
| **Testbarkeit** | hoch — Testwerkzeuge in der Standarddistribution | hoch | hoch | hoch | hoch | hoch |
| **Künftige Cross-Plattform-Implikationen** | hoch — Cross-Kompilierung ohne Fremdtoolchain | hoch — mit höherem Toolchain-Aufwand | mittel | mittel | hoch | hoch |
| **Supply-Chain-Implikationen** | günstig — wenige Abhängigkeiten, prüfsummengebundene Auflösung | günstig — Lockfile, aber tendenziell tiefere Abhängigkeitsbäume | ungünstig — breite, transitiv tiefe Ökosystemoberfläche | ungünstig — sehr breite, tiefe Ökosystemoberfläche | mittel | mittel |

### 8.3 Empfehlung

```text
RECOMMENDATION (language/runtime):  Go als bevorzugter Kandidat.
                                    Rust als stärkste Alternative.
STATUS:                             PROPOSED / UNACCEPTED
LANGUAGE / RUNTIME:                 NOT SELECTED
```

**Begründung.** Der erste Slice ist bewusst schmal: ein read-only Erhebungspfad, der Dateien liest, normalisiert und ein Ergebnis mit Provenance erzeugt. Für diese Aufgabe schneidet Go auf den CoreOps-tragenden Kriterien am besten ab — einzelnes selbstständiges Binary ohne Laufzeitinstallation (Standalone-First, Offline, minimaler Footprint), prüfsummengebundene Modulauflösung mit Vendoring (deterministische Paketierung, Supply-Chain), plausibel **null** Drittabhängigkeiten für diesen Slice (Dependency-Admission bleibt leer), Testwerkzeuge in der Standarddistribution und einfache Cross-Kompilierung für spätere Plattformerweiterung.

Rust ist bei Sicherheitsgarantien überlegen und bei Paketierung gleichwertig; die höheren Einstiegs-, Umbau- und Toolchain-Kosten stehen für einen ersten, absichtlich kleinen Slice in keinem günstigen Verhältnis. Python und TypeScript/Node.js scheitern nicht an der Aufgabe, sondern an den CoreOps-Randbedingungen: Laufzeitabhängigkeit, größerer Footprint, schwächere deterministische Paketierung und eine deutlich breitere Supply-Chain-Oberfläche. JVM und .NET sind tragfähig, bringen aber ohne Gegenwert für diesen Slice eine schwerere Laufzeit- und Paketierungsgeschichte mit.

**Diese Empfehlung ist keine Auswahl.** Sie ist ausdrücklich ADR-relevant und bleibt bis zu einer eigenen Human-Maintainer-Entscheidung `PROPOSED / UNACCEPTED`.

```text
language recommendation != language selection
Docker-first            != language constraint
```

## 9. Source-Tree-Vorschlag

**Status: `PROPOSED / UNACCEPTED`. Es wird kein Verzeichnis angelegt.** Der Vorschlag ist ausschließlich Markdown und beschreibt **Rollen**, keine Sprachidiome — das konkrete Verzeichnisidiom folgt der Sprachentscheidung aus §8 und ist hier bewusst nicht vorweggenommen.

```text
<repository root>
├─ docs/                      bestehend, unverändert
├─ project-brain/             bestehend, unverändert
├─ project-system/            bestehend, unverändert
└─ <application root>/        VORSCHLAG — nicht angelegt
   ├─ <entry point>/          Programmeinstieg; enthält keine Domänenlogik
   ├─ <contract>/             Vertragstypen des Observation Contract:
   │                          Feldidentität, observation_outcome, Zeitmodell,
   │                          Provenance, Freshness — ohne Erhebungswissen
   ├─ <collection>/           Erhebung; kapselt den P-3-Mechanismus vollständig
   │                          hinter der Vertragsgrenze
   ├─ <normalization>/        Raw auf kanonische Repräsentation; kennt keine Quelle
   └─ <tests>/                Testfälle des Test Envelope
```

**Optimierungsziele** und ihre Umsetzung im Vorschlag:

| Ziel | Umsetzung |
| ---- | --------- |
| Schmaler erster Slice | genau vier Rollen plus Einstieg; keine Schichten „auf Vorrat" |
| Klare Domänengrenzen | `<contract>` kennt keinen Mechanismus; `<collection>` kennt keine Normalisierung; `<normalization>` kennt keine Quelle |
| Künftige Testbarkeit | die `P-3`-Abhängigkeit liegt vollständig in `<collection>`; die übrigen Rollen sind ohne Ziel und ohne Umgebung prüfbar |
| Minimum an verfrühter Abstraktion | keine Plugin-, Adapter-, Registry-, Transport- oder Konfigurationsschicht; keine Datenbank; kein Service-Layer |
| Standalone-First | ein Anwendungswurzelverzeichnis, ein Einstiegspunkt, keine externe Laufzeitvoraussetzung im Entwurf |

**Ausdrücklich nicht Bestandteil des Vorschlags:** ein Frontend, eine API-Schicht, eine Persistenzschicht, ein Agent, ein Scheduler, ein Konfigurationssystem, eine CI-Definition, ein Containerartefakt.

```text
source-tree proposal != source tree created
directory named      != directory exists
```

## 10. Dependency-Admission-Gate

**Status: keine Abhängigkeit ist zugelassen. Keine wird durch dieses Dokument zugelassen.**

Das Gate **erweitert** die [Sovereignty and Dependency Policy](../architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md) §6 (deren zehn Zulassungskriterien unverändert gelten) um die Unterscheidung, die vor dem ersten produktiven Commit gebraucht wird.

### 10.1 Grundunterscheidung

| Klasse | Definition | Behandlung |
| ------ | ---------- | ---------- |
| **Standardbibliothek / Plattformfähigkeit** | Bestandteil der gewählten Sprach-/Runtime-Distribution selbst; kein separater Bezug, keine separate Versionierung, keine eigene Lieferkette | **kein** Admission-Verfahren; die Zulassung erfolgte implizit mit der Sprach-/Runtime-Entscheidung, deren Lieferkette gesondert zu bewerten ist |
| **Drittabhängigkeit** | jedes separat bezogene Paket, Modul, jede Bibliothek, jedes Werkzeug oder Artefakt mit eigener Herkunft, Versionierung und Lieferkette — **auch** wenn es klein, populär oder scheinbar trivial ist | **vollständiges** Admission-Verfahren je Abhängigkeit |

```text
popular       != vetted
small         != harmless
transitive    != out of scope
build-only    != outside the supply chain
```

Transitive und reine Build-/Test-Abhängigkeiten unterliegen demselben Verfahren; sie sind Teil der Lieferkette.

### 10.2 Admission-Kriterien je Drittabhängigkeit

| # | Kriterium | Anforderung |
| - | --------- | ----------- |
| 1 | **Notwendigkeit** | Warum ist die Standardbibliothek nicht ausreichend? Was ist der konkrete, benannte Bedarf? Bequemlichkeit ist kein Bedarf |
| 2 | **Herkunft** | Herausgeber, Repository, Bezugsweg, Verantwortlichkeit — identifizierbar und dokumentiert |
| 3 | **Pinning** | exakte Version, nicht Bereich; Auflösung an Prüfsumme oder Digest gebunden |
| 4 | **Lizenzierung** | Lizenz benannt, mit dem CoreOps-Veröffentlichungsmodell (§12) vereinbar, Weitergabepflichten benannt |
| 5 | **Offline-Verfügbarkeit** | über kontrolliertes lokales Repository oder Vendoring beziehbar; eine zwingend online auflösende Abhängigkeit ist für Kernfunktionen unzulässig |
| 6 | **Schwachstellenbehandlung** | benannter Prozess: Beobachtung, Bewertung, Reaktionsweg, Verantwortlicher |
| 7 | **Integrität** | Prüfsumme oder Digest verifizierbar; Signatur/Provenance sofern verfügbar; SBOM-Erfassung nach [Artifact Identity, Provenance and SBOM Model](../architecture/ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md) |
| 8 | **Reproduzierbarkeit** | derselbe Eingang erzeugt denselben Auflösungsstand; Auflösung ist nicht zeit- oder netzwerkabhängig |
| 9 | **Update-Policy** | wer aktualisiert wann, nach welcher Prüfung, mit welchem Rückfallweg |
| 10 | **Ersetzungs-/Entfernungskosten** | Exit-Strategie nach Sovereignty Policy §9: wie teuer ist der Ausbau bei Lizenz-, Sicherheits- oder Verfügbarkeitsproblem |

**Zulassung erfolgt ausschließlich durch eine ausdrückliche Human-Maintainer-Entscheidung, je Abhängigkeit, dokumentiert.** Es gibt keine pauschale Zulassung einer Ökosystem-Familie und keine stille Zulassung durch Verwendung.

### 10.3 Zielzustand für den ersten Slice

Der empfohlene Zuschnitt (§7 Option A/B, §8 Go) kommt plausibel **ohne** Drittabhängigkeit aus. Eine leere Abhängigkeitsliste ist für den ersten produktiven Commit ein **erklärtes Ziel**, keine bloße Möglichkeit.

## 11. Package-Manager- und Lockfile-Implikationen

Der [Foundation Scope Lock](FOUNDATION_SCOPE_LOCK.md) führte *Lockfiles von Package Managern, Dependencies* als *Forbidden Implementation Type* der **abgeschlossenen** Foundation-Phase. Mit `HM-F1` ist diese Phase geschlossen; die Regel bleibt für sie gültig und wird hier **nicht** aufgehoben. Für Post-Foundation-Arbeit gilt:

- Ein Lockfile ist ein **Lieferketten-Governance-Artefakt**, kein Nebenprodukt. Seine erstmalige Erzeugung ist ein eigenes Autoritätsereignis, kein Implementierungsdetail.
- Ein Lockfile darf erst entstehen, **nachdem** Sprache/Runtime entschieden und die Dependency-Admission-Policy in Kraft ist — sonst dokumentiert es Zulassungen, die nie erteilt wurden.
- Ein Lockfile mit **null** Drittabhängigkeiten ist ein zulässiges und erwünschtes Ergebnis; er belegt die leere Lieferkette ausdrücklich.
- Jede Lockfile-Änderung ist reviewpflichtig wie eine Codeänderung: sie kann Herkunft, Version und Integrität stillschweigend verschieben.
- Der Auflösungsvorgang muss **offline** reproduzierbar sein (Vendoring oder kontrolliertes lokales Repository).
- Ein Lockfile belegt **Auflösungsstand**, nicht Sicherheit.

```text
lockfile present    != dependencies vetted
lockfile present    != build reproducible
resolution pinned   != artifact integrity verified
zero dependencies   != zero supply chain (toolchain remains in scope)
```

**Derzeit existiert kein Package Manifest, kein Lockfile und kein Package Manager im Repository. Keines wird durch dieses Dokument autorisiert.**

## 12. Build- und Packaging-Implikationen

Vorbereitend, **nicht** entschieden:

| Thema | Anforderung an eine spätere Entscheidung |
| ----- | ---------------------------------------- |
| **Reproduzierbarkeit** | derselbe Quellstand erzeugt ein identifizierbar gleiches Artefakt; Build-Eingänge sind deklariert |
| **Offline-Build** | der Build muss ohne Netzwerkzugriff durchführbar sein (Sovereignty Policy §7) |
| **Artefaktidentität** | Artefakt-Identität, Provenance und SBOM nach dem bestehenden [Artifact Identity, Provenance and SBOM Model](../architecture/ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md); `SBOM vorhanden != vollständig != sicher` |
| **Toolchain** | die Build-Toolchain ist selbst Teil der Lieferkette und unterliegt Herkunfts-, Pinning- und Integritätsanforderungen |
| **Container** | `HM-2` setzt Docker-first als **Delivery Baseline** — `Docker-first != Docker-only != zwingende Runtime-Abhängigkeit`. Ein Containerartefakt ist eine **Auslieferungsform**, keine Architekturvorgabe und keine Voraussetzung des Slice |
| **CI** | das Manifest führt `ci: recommended` mit `Status: pending`. Eine CI-Definition ist **nicht** autorisiert und **nicht** Voraussetzung des ersten produktiven Commits |

```text
build succeeded     != artifact verified
artifact built      != artifact released
container image     != product release
CI configured       != quality established
```

**Derzeit existiert keine Build-Definition, kein Containerartefakt und kein CI-Workflow. Keines wird durch dieses Dokument autorisiert.**

## 13. `NEW-8` — `README.md` und `LICENSE`

`NEW-8` ist der in `CO-WP-031` **deferred** deklarierte Restpunkt: im öffentlichen Repository existiert weder `README.md` noch `LICENSE`. `CO-WP-031` hat ihn ausdrücklich nicht behoben und kein Risiko registriert. Dieser Abschnitt liefert die verlangte **Dispositionsempfehlung** — und **nur** diese.

### 13.1 Soll `README.md` vor dem ersten produktiven Source-Commit existieren?

```text
RECOMMEND: YES
STATUS:    PROPOSED / UNACCEPTED
```

**Begründung.** Ab dem ersten produktiven Source-Commit ändert sich, was ein Betrachter des öffentlichen Repositories vernünftigerweise annimmt: aus einem erkennbaren Dokumentationskorpus wird ein Repository mit Code. Ohne Einstiegspunkt ist die naheliegende Fehlannahme „hier gibt es ein benutzbares Produkt" — genau die Behauptung, die CoreOps durchgängig vermeidet. Der `README` ist hier ein **Ehrlichkeits-Instrument**, nicht Marketing. Die [Source-of-Truth-Hierarchie](REPOSITORY_GOVERNANCE_STANDARD.md) §7 führt `README and public summaries` auf Rang 19 — bewusst **unter** allen normativen Quellen: ein `README` fasst zusammen, er darf nichts überschreiben.

### 13.2 Welchen Mindestinhalt sollte er tragen?

| Element | Anforderung |
| ------- | ----------- |
| Was CoreOps ist | Projektname, Slogan, eine ehrliche Kurzbeschreibung |
| **Ehrlicher Status** | Foundation 0.1 geschlossen · Observe betreten mit Grenze · **kein** funktionaler Produktrelease · **kein** benutzbares Produkt · **0** implementierte Runtime-Capabilities · **0** `supported` Integrationen · **0** akzeptierte ADRs |
| Was das Repository **enthält** | Dokumentation, Governance, Register — und ab dem ersten produktiven Commit: welcher eng begrenzte Slice |
| Was es **nicht** enthält | Installation, Betrieb, Support, Zusicherung, Zeitplan |
| Einstiegspunkte | Verweise auf Project Brief, Scope Lock, Release-Taxonomie, Work Package Queue |
| Lizenz-/Nutzungshinweis | Verweis auf den tatsächlichen Lizenzstand — **ohne** Rechte zu behaupten, die nicht eingeräumt sind |
| Governance-Hinweis | Human-Maintainer-Vorbehalt; keine Beitragszusage ohne eigenen Prozess |
| Sprache | DE oder EN mit klarem Sprachstatus; **keine** Paritätsbehauptung ohne geprüfte Übersetzung |

**Nicht** in den `README`: Reife-, Sicherheits-, Zertifizierungs-, Compliance- oder Supportbehauptungen; Zeitpläne; Fertigstellungsversprechen; `v1.0`-Andeutungen; private Daten.

### 13.3 Muss die Lizenz-Disposition vor der Veröffentlichung produktiven Quellcodes erfolgen?

```text
RECOMMEND: YES
STATUS:    PROPOSED / UNACCEPTED
```

**Begründung.** Solange nur Dokumentation im öffentlichen Repository liegt, ist die fehlende Lizenz ein erklärter, dispositionierter Restpunkt. Mit veröffentlichtem produktivem Quellcode wird sie materiell: ein öffentliches Repository ohne Lizenz räumt Dritten **keine** Nutzungs-, Änderungs- oder Weitergaberechte ein — sichtbar ist nicht dasselbe wie nutzbar. Die Disposition muss also **vor** der Veröffentlichung getroffen werden, damit der veröffentlichte Zustand das ausdrückt, was tatsächlich gewollt ist. Eine bewusste Entscheidung, **keine** Lizenz zu vergeben, ist dabei eine vollwertige Disposition — sie muss nur getroffen und sichtbar sein statt zu unterbleiben.

> Dies ist eine **Governance-Empfehlung, keine Rechtsberatung.** Die Wahl eines Lizenzmodells hat rechtliche Wirkung; der Human Maintainer entscheidet sie eigenständig und zieht dafür bei Bedarf fachkundige Beratung hinzu.

### 13.4 Abgrenzung: privat/intern · Weitergabe · öffentliche Veröffentlichung

| Modus | Bedeutung | Lizenzbedarf |
| ----- | --------- | ------------ |
| **Privat / intern** | Nutzung ausschließlich durch den Rechteinhaber und Personen unter seiner Kontrolle; keine Weitergabe an Dritte | Lizenz für die eigene Nutzung nicht erforderlich; die Frage bleibt offen, nicht beantwortet |
| **Weitergabe (Redistribution)** | Dritte erhalten das Artefakt — auch bei Weitergabe an einen einzelnen Empfänger, auch ohne Entgelt | Eine ausdrückliche Rechteeinräumung ist erforderlich; ohne sie erhält der Empfänger keine Nutzungsrechte |
| **Öffentliche Veröffentlichung / Open Source** | jedermann kann beziehen, nutzen, ändern und weitergeben | Erfordert eine ausdrückliche Open-Source-Lizenz **und** eine Entscheidung über Beitragsannahme, Marken, Namensnennung und Haftungsausschluss |

```text
public repository != open source
publicly visible  != publicly licensed
no license        != permission granted
code published    != redistribution right granted
```

### 13.5 Welche Human-Maintainer-Autorität ist erforderlich?

| Artefakt | Erforderliche Autorität |
| -------- | ----------------------- |
| `README.md` | Ausdrückliche Human-Maintainer-Autorisierung der **Erstellung**, plus Autorisierung des konkreten Inhalts (der `README` trifft öffentliche Statusaussagen), plus die üblichen Staging-/Commit-/Push-Gates |
| `LICENSE` | Ausdrückliche Human-Maintainer-Entscheidung über das **Lizenzmodell** (eine Rechtsentscheidung des Rechteinhabers), **danach** Autorisierung der Dateierstellung, plus die üblichen Staging-/Commit-/Push-Gates. Zusätzlich zu dispositionieren: Beitragsannahme, Marken- und Namensnennung, Haftungs- und Gewährleistungsausschluss |

```text
README/LICENSE recommendation != artifact authorization
recommendation to create      != permission to create
```

**Weder `README.md` noch `LICENSE` wird durch dieses Work Package erstellt.** Beide bleiben nicht existent.

## 14. ADR- und Decision-Voraussetzungen

`CO-WP-032` erzeugt **keine** ADR-Datei, **keine** Decision-ID, **keine** Risk-ID und keine Zeile im Decision Index oder Risk Register. Es existieren weiterhin **0** ADR-Dateien und **0** akzeptierte ADRs. Die Entscheidungspakete in §7, §8 und §9 sind **ADR-Vorbereitung**, nicht ADR.

Vor einer produktiven Implementierung sind mindestens folgende Punkte zu dispositionieren — jeweils durch den Human Maintainer, jeweils unter der bestehenden Relevanzregel `HM-3`:

| Punkt | Gegenstand | Aktueller Stand |
| ----- | ---------- | --------------- |
| `P-3` | Erhebungsmechanismus und Transport für `CAP-DISCOVERY-004` — im Readiness Review ausdrücklich als ADR-pflichtig bezeichnet | `NOT SELECTED` |
| Sprache/Runtime | erste Implementierungsplattform | `NOT SELECTED` |
| Source Tree | Anwendungsstruktur | `PROPOSED / UNACCEPTED` |
| Dependency-Admission | Inkraftsetzung des Gates aus §10 | nicht in Kraft |
| `NEW-8` | `README` / `LICENSE` | `deferred`, Empfehlung liegt vor |
| Build/Packaging | Reproduzierbarkeit, Offline-Build, Artefaktidentität | nicht entschieden |

**Verhältnis zu den offenen `CCR`.** Die sechs offenen `CCR` bleiben offen und werden hier **nicht** geschlossen. `CCR-05` und `CCR-07` tragen `MUST CLOSE BEFORE DEPLOY`; das Readiness Review hat festgestellt, dass beide privilegierte beziehungsweise Offline-Ausführung betreffen und einen read-only Slice **nicht** berühren — sie sind ausdrücklich **keine** Observe-Precondition. Diese Feststellung wird hier nur referenziert, nicht neu getroffen.

```text
ADR candidate exists       != ADR required
ADR prepared               != ADR accepted
decision packet            != decision
foundation semantics established != ADR accepted
```

## 15. Künftige Testvoraussetzungen

Die Testvoraussetzungen verteilen sich auf **beide** Gates und dürfen nicht vermengt werden. Der [Test Envelope](../testing/OBSERVE_LOCAL_LINUX_HOST_TEST_ENVELOPE.md) §10 führt die Ausführungsvoraussetzungen vollständig.

### 15.1 Gate A — vor produktivem Code zu dispositionieren

Auf dieser Ebene ist ausschließlich verlangt, dass **definiert ist, was später geprüft und belegt werden muss** — nicht, dass es bereits geprüft oder belegt wäre.

1. **Teststrategie beziehungsweise künftige Validierungsanforderungen definiert** — für den Slice erfüllt durch den Test Envelope (zehn Mindestfälle mit Intent, erwartetem Ausgang, Datensemantik, Provenance-Verhalten, verbotener Inferenz und künftiger Evidenzanforderung).
2. **`P-2`-Evidenzmethode definiert** — erfüllt durch Test Envelope §8.3 (sieben kumulative Anforderungen). **Nicht** verlangt: dass die Evidenz vorliegt.
3. `P-3` entschieden — **offen**; ohne die Entscheidung ist kein Fall implementierbar.
4. Sprache/Runtime entschieden — **offen**; ohne die Entscheidung existiert kein Testwerkzeug.

### 15.2 Gate B — vor Testausführung und realer Beobachtung

5. Fixtures mit Identity, Revision und Provenance freigegeben — **existieren nicht**.
6. Lab-Environment mit deklarierter Environment Identity bereitgestellt — **nicht bereitgestellt**.
7. `P-1` erteilt, soweit reale Ziele berührt werden — **nicht erteilt**.
8. **Testausführung ausdrücklich autorisiert** — **nicht erteilt**.
9. `P-2` durch beobachtetes Schutzverhalten belegt — **nicht belegt**.

```text
test planned      != test implemented != test executed != test passed
test design       != test execution
test strategy defined != tests implemented
test result       != support status != production readiness
zero tests executed = the current state
```

**Keiner der Punkte 5 bis 9 ist eine Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren.**

## 16. Transition-Checklisten (fail-closed, zwei Gates)

Beide Checklisten sind **fail-closed**: `nicht dispositioniert` zählt als **nicht erfüllt** — nicht als „vermutlich in Ordnung". Sie sind **getrennt** zu bewerten; eine bestandene Gate-A-Checkliste sagt über Gate B nichts aus und umgekehrt.

### 16.1 Gate A — vor Produktivcode-/Implementierungsautorisierung

Eine spätere Produktivcode-Autorisierung gilt **nicht** als vorbereitet, solange auch nur eine dieser Zeilen offen ist.

| # | Punkt | Anforderung an die Disposition | Stand heute |
| - | ----- | ------------------------------ | ----------- |
| A-1 | **Observe-Phasenautorität** | Observe betreten und die Grenze von `HM-O1` ausdrücklich benannt | **erfüllt** — `HM-O1` `APPROVED WITH BOUNDARY` |
| A-2 | **Gewählter Value Slice** | ein benannter, eng begrenzter Slice | **erfüllt** — `HM-O2` `APPROVED` |
| A-3 | **`CO-WP-032` geschlossen** | Nova Review abgeschlossen **und** Human-Maintainer-Integration erfolgt | **erfüllt** — Nova Final Review `GO`, Human-Maintainer-Integration abgeschlossen (Commit `9999114`) |
| A-4 | **Nova Review** | Bewertung des Slice-Vertrags, des Test Envelope und dieses Gate-Dokuments liegt vor | **erfüllt** — Initial Review `REWORK — narrow semantic closure`, beide blockierenden Notes CLOSED, **Nova Final Review `GO`** |
| A-5 | **Repository-Integration** | Human-Maintainer-Staging, -Commit und -Push erfolgreich durchgeführt | **erfüllt** — Commit `9999114200bf18baaadfb508e8464720b75e352e`, gepusht, `origin/main` gleichauf |
| A-6 | **`P-3`-Entscheidung** | Erhebungsmechanismus und Transport ausdrücklich entschieden | **offen** — `NOT SELECTED` |
| A-7 | **Sprach-/Runtime-Entscheidung** | Plattform ausdrücklich entschieden | **offen** — `NOT SELECTED` |
| A-8 | **Source-Tree-Entscheidung** | Struktur ausdrücklich entschieden | **offen** — `PROPOSED / UNACCEPTED` |
| A-9 | **Dependency-Admission-Disposition** | Admission-Gate in Kraft; jede vorgesehene Abhängigkeit einzeln zugelassen oder ausdrücklich keine | **offen** — nicht in Kraft, **0** zugelassen |
| A-10 | **`NEW-8`-Disposition** | `README` und `LICENSE` entschieden — auch eine bewusste Nicht-Erstellung ist eine Disposition | **offen** — `deferred`, Empfehlung liegt vor |
| A-11 | **ADR-/Decision-Disposition** | die Punkte aus §14 entschieden oder ausdrücklich zurückgestellt | **offen** |
| A-12 | **Teststrategie / künftige Validierungsanforderungen definiert** | §15.1 Punkte 1 und 2 — einschließlich der **definierten `P-2`-Evidenzmethode**; ausdrücklich **nicht** `P-2` erfüllt | **teilweise** — Testdesign und `P-2`-Methode **definiert**; die Human-Maintainer-Disposition dieser Definition steht aus |
| A-13 | **Build-/Packaging-Disposition** | Reproduzierbarkeit, Offline-Build und Artefaktidentität entschieden, soweit §12 sie als erforderlich ausweist | **offen** |
| A-14 | **Ausdrückliche Implementierungs-Work-Package-Autorität** | ein eigenes, ausdrücklich autorisiertes implementierungsorientiertes Work Package mit benanntem Scope | **offen** — nicht erteilt |

**Bilanz Gate A: 5 von 14 erfüllt, 1 teilweise, 8 offen. Gate A ist nicht passiert.**

**Ausdrücklich nicht in dieser Checkliste:** `P-1` · `P-2` erfüllt · Testausführungsautorität · Fixtures · Lab. Sie sind **keine** Voraussetzung dafür, produktiven Quellcode zu schreiben oder zu autorisieren.

### 16.2 Gate B — vor realer Beobachtung / Zielausführung

Dieses Gate wird **erst nach** Gate A und **erst nach** einer vorhandenen Implementierung überhaupt bewertbar.

| # | Punkt | Anforderung an die Disposition | Stand heute |
| - | ----- | ------------------------------ | ----------- |
| B-1 | **`P-1` konkrete Zielautorisierung** | konkretes Ziel benannt (keine Zielklasse als Blankett) | **offen** — `NOT AUTHORIZED` |
| B-2 | **Exakt erlaubte Beobachtung** | welche Felder, aus welchen Quellklassen, mit welcher Erhebungsmechanismusklasse | **offen** |
| B-3 | **Exakte Ausführungsgrenze** | read-only; keine Rechteeskalation, keine Credentials, keine Secrets, kein Netzwerkziel; zeitliche Begrenzung, Widerrufbarkeit, Auditpflicht | **offen** |
| B-4 | **`P-2` No-Mutation-/Ausführungssicherheits-Evidenz** | beobachtetes geschütztes Verhalten nach Test Envelope §8.3 — **erst nach** Implementierung und Ausführung möglich | **offen** — `NOT SATISFIED` |
| B-5 | **Testausführungsautorität** | ausdrückliche Human-Maintainer-Autorisierung der Ausführung | **offen** — nicht erteilt |
| B-6 | **Fixtures** | Identity, Revision und Provenance freigegeben | **offen** — existieren nicht |
| B-7 | **Lab-Environment** | deklarierte Environment Identity bereitgestellt | **offen** — nicht bereitgestellt |
| B-8 | **Sonstige ausführungsspezifische Autorität** | jede weitere für den konkreten Zugriff erforderliche Freigabe | **offen** |

**Bilanz Gate B: 0 von 8 erfüllt. Realer Zielzugriff bleibt untersagt.**

```text
checklist documented != checklist satisfied
Gate A open          = productive code NOT AUTHORIZED
Gate B open          = target access NOT AUTHORIZED
Gate A passed        != Gate B passed
productive code authorization != target authorization
```

## 17. Kein automatischer Nachfolger

```text
Es erfolgt keine automatische Freigabe eines Folge-Work-Packages.
```

Der Abschluss von `CO-WP-032` — einschließlich Nova Review und Human-Maintainer-Integration — erzeugt **keinen** Nachfolger, reserviert **keine** Kennung und begründet **keine** Erwartung.

Zulässige Formulierung, und nur diese:

> **Ein späteres implementierungsorientiertes Work Package kann in Betracht gezogen werden, nachdem `CO-WP-032` geschlossen ist und die erforderlichen Human-Maintainer-Entscheidungen getroffen sind.**

```text
candidate successor != reserved work package
recommendation      != authorization
closure             != successor
```

Es wird **keine** Nachfolger-Kennung vergeben, **keine** reserviert und **keine** genannt. Es wird **kein** neuer primärer Work-Package-Typ eingeführt.

## 18. Compatibility

Additiv. **Keine** Änderung an Foundation Scope Lock, Release-Taxonomie, Decision Index, Risk Register, Lessons-Learned-Register, NDF-Feedback-Kandidaten, Capability Matrix, Initial Support Boundary, Roadmap oder Teststrategie. **Kein** ADR erzeugt oder akzeptiert. **Keine** Decision-, Risk- oder ADR-ID vergeben. **Keine** Technologie ausgewählt. **Kein** Artefakt außerhalb der neun autorisierten Pfade erzeugt. Breaking-Change-Potenzial: gering.

## 19. Next Decision

Der Nova Final Review dieses Dokuments ist erfolgt (`GO`) und die Human-Maintainer-Repository-Integration ist abgeschlossen (Commit `9999114200bf18baaadfb508e8464720b75e352e`, gepusht); die Gate-A-Punkte `A-1` bis `A-5` sind damit erfüllt. **Als Nächstes** — und jeweils als eigene, ausdrückliche Human-Maintainer-Entscheidung — die verbleibenden Gate-A-Punkte `A-6` bis `A-14` (§16.1). **Erst danach und erst mit einer vorhandenen Implementierung** werden die Gate-B-Punkte `B-1` bis `B-8` (§16.2) überhaupt bewertbar. Dieses Dokument trifft **keine** dieser Entscheidungen.
