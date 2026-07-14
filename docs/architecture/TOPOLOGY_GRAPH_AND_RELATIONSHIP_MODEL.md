# CoreOps – Topology Graph and Relationship Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation topology graph and relationship model
> Implementation Status: Not implemented
> Graph Technology: Not selected
> Discovery Technology: Not selected
> Visualization Technology: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-020 (docs-only / topology, relationship, evidence and manual-governance foundation)

## 1. Status

Technologie- und herstellerunabhängiges Modell für den **CoreOps Topology Graph**: Nodes, Edges, Relationship Assertions, Identität, Views/Snapshots und Grenzen. Companion zu [TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md](TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md) und [TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md](../security/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md). Es implementiert **keine** Graph-Datenbank, keine Discovery Engine, keine Netzwerkprotokolle, keine Visualisierung.

## 2. Purpose

Ein Topology-Graph verführt dazu, Darstellung mit Realität zu verwechseln: `graph presence ≠ resource currently reachable`, `edge ≠ verified connectivity`, `manual ≠ fact`. Dieses Modell trennt den Graph (Menge verwalteter Assertions) von der physischen Realität, trennt Node-/Edge-/Assertion-Identität und legt fest, dass Topologie weder autoritativen State noch Netzwerk-/Execution-Autorität gewährt.

## 3. Scope

Topology-Begriffe · Graph-Grenze · Node/Relationship Classes · Node/Edge/Assertion Identity · Assertion Origins · Producers/Sources · Canonical Identity/Aliases · Identity Resolution · Duplicate/Merge/Split · Temporal Validity · Views/Snapshots · Conflict Handling · State-/Event-/Evidence-/Execution-Grenzen · Workspace/Disclosure · Offline · Audit. Evidence/Confidence/Conflict-Detail (Companion 2), Manual Authority/Disclosure (Companion 3).

## 4. Non-Goals

- Keine Graph-Datenbank/Graph-Query-Language, keine Discovery Engine, keine CMDB.
- Keine SNMP-/LLDP-/CDP-/ARP-/NetFlow-/Nmap-/Netzwerk-/Agent-Protokoll-Auswahl.
- Keine Node-/Edge-ID-Technologie, keine Visualisierung/Layout Engine, kein Graph-Schema.
- Kein Runtime-Code, keine echte Netzwerk-Discovery; keine Behauptung implementierter/vollständiger Topologie.

## 5. Concepts

Begriffe (mindestens): topology · topology graph · topology view · topology snapshot · node · edge · relationship · relationship assertion · relationship type · node/edge identity · canonical identity · alias · identity resolution · duplicate candidate · merge · split · source · producer · observation · declaration · manual assertion · imported assertion · inference · confidence · evidence · validity interval · conflict · override · suppression · current/historical view.

**Grundregeln:**
```text
topology graph        ≠ authoritative physical reality
node                  ≠ managed resource automatically
edge                  ≠ verified connectivity automatically
relationship assertion ≠ validated relationship
discovered            ≠ trusted
observed              ≠ currently valid indefinitely
manual                ≠ automatically authoritative
inferred              ≠ independently observed
same display name     ≠ same node
missing edge          ≠ no relationship exists
```

## 6. Graph Boundary

Der Topology Graph ist eine **verwaltete Darstellung von Node-/Relationship-Assertions**. Er darf enthalten: mehrere konkurrierende Assertions · source-/workspace-spezifische Views · current/historical Views · Confidence/Evidence-Referenzen · Conflict-/Unresolved-States. Er ist **nicht** automatisch vollständige/aktuelle Realität, Netzwerkfreigabe, Execution Plan, Discovery-Beweis oder Security-Zonen-Zertifizierung.
```text
graph presence     ≠ resource currently reachable
graph path         ≠ authorised network path
topology relation  ≠ execution authorization
```

## 7. Node Classes

Je Klasse: Purpose · Identity expectation · Authoritative owner · Typical sources · Typical relationships · Workspace/scope · Temporal behavior · Evidence expectation · Disclosure sensitivity · Offline relevance · Threat references. **Keine vollständige Node-Instanzliste.**

| Node-Klasse | Identity-Erwartung (kurz) | Threat refs |
|---|---|---|
| workspace / environment | logische Governance-Grenze | THR-035 |
| site or logical location | Standort/Zone | THR-018 |
| physical host / bare-metal | Hardware-Identität | THR-034 |
| virtual machine | VM-Identität | THR-034 |
| container / workload | Workload-Identität | THR-034 |
| cluster | Cluster-Identität | THR-034 |
| network device / interface | Netzobjekt | THR-039, THR-014 |
| network segment / subnet / VLAN | logisches Netz | THR-015 |
| route or routing domain | Routing | THR-015 |
| service / application | Dienst/App | THR-011 |
| data or storage resource | Datenobjekt | THR-014 |
| printer / print queue | Druck | THR-025 |
| integration / agent or relay | Vermittlung | THR-008, THR-010 |
| external system | extern | THR-011 |
| backup or recovery target | Backup | THR-033 |
| deployment target | Deployment | THR-022 |
| unknown resource | unbekannt | THR-014 |

`node present ≠ resource currently available`.

## 8. Relationship Classes

Je Class: directionality · symmetry · source/target classes · multiplicity · temporal behavior · authority expectation · evidence expectation · confidence expectation · security relevance · disclosure sensitivity. Klassen: `contains · member-of · hosts · runs-on · connects-to · linked-to · routes-to · depends-on · managed-by · observed-by · integrated-with · authenticates-via · deployed-by · backed-up-by · prints-to · attached-to · part-of · located-in · communicates-with · inferred-related-to · unknown`.
```text
connects-to  ≠ communication verified
depends-on   ≠ dependency currently healthy
managed-by   ≠ unrestricted administrative authority
observed-by  ≠ authoritative ownership
located-in   ≠ precise physical geolocation necessarily
```

## 9. Node Identity

20 Felder: stable topology-node identity · resource identity reference (falls verfügbar) · node class · canonical name · display name · owner · workspace/environment · lifecycle state · identity-resolution state · known aliases · source identities · first-/last-observed boundary · validity state · confidence state · evidence references · conflict state · merge/split history · disclosure classification · audit reference.
```text
display name    ≠ node identity
hostname        ≠ globally unique identity
IP address      ≠ stable resource identity
MAC address     ≠ universal trustworthy identity
external-system ID ≠ CoreOps canonical identity automatically
```
Keine ID-Technologie ausgewählt.

## 10. Edge and Assertion Identity

Eine Relationship Assertion benötigt: stable assertion identity · relationship class · source/target node · directionality · producer · source · source trust state · assertion origin · observation/declaration time · recording time · valid-from/valid-until (falls bekannt) · confidence state · validation state · evidence references · workspace/scope · conflict state · supersession reference · suppression state · manual-authority reference (falls anwendbar) · audit reference.
```text
edge identity        ≠ source node identity ≠ target node identity
same node pair       ≠ same relationship
same relationship type ≠ same validity period
new assertion        ≠ previous assertion deleted
```

## 11. Assertion Origins

`discovered · observed · declared by authoritative system · imported · manually asserted · derived · inferred · reconciled · historical · unknown`.
```text
declared   ≠ independently observed
imported   ≠ validated
manual     ≠ fact
derived    ≠ source observation
reconciled ≠ conflict-free automatically
```
**Keine Origin-Klasse legt allein Authority/Confidence/Validation fest** (Companion 2).

## 12. Producers and Sources

Getrennt: relationship producer · source · source system · source trust · source authority · recording component · transformation/reconciliation component (Detail Companion 2 §7-8). `producer ≠ source`; `source trust ≠ source authority`; `source authority ≠ global authority`; `recording component ≠ relationship owner`. Source Trust ist zeitgebunden; spätere Kompromittierung löscht historische Assertions nicht, löst aber Neubewertung aus.

## 13. Canonical Identity

Canonical Node Identity ist stabil und getrennt von Aliases/Display Name (Companion 2 §20). `alias match ≠ identity match`.

## 14. Aliases

Aliases (Hostnamen, externe IDs, Adressen, Inventarnummern, Herstellerkennungen, Agent-/Adapter-IDs, Drucker-/Queue-Namen) benötigen Source, Scope, Zeitbezug, Trust, Evidence.
```text
alias match     ≠ identity match
same hostname   ≠ same resource
same address    ≠ same resource over time
same external ID ≠ same workspace scope automatically
```

## 15. Identity Resolution

Status: `not-assessed · candidate · probable-match · confirmed-match · conflicted · rejected-match · split-required · merge-pending · merged · invalidated · unknown`. Resolution Decision benötigt: candidate nodes · matching/conflicting attributes · workspace/scope · source trust · temporal overlap · evidence references · confidence · decision owner · review trigger · audit reference. **Keine automatische Merge Engine.** (Bezug THR-014, THR-035.)

## 16. Duplicate Candidates

`duplicate candidate ≠ duplicate confirmed`. Duplicate-Erkennung basiert nicht auf Alias/Timestamp allein (Companion 2).

## 17. Merge

Merge benötigt: explizite Entscheidung · Scope · Owner · Evidence · Aliasbehandlung · Relationship-Neuzuordnung · Conflict-Behandlung · Reversibility/Forward-Recovery-Bewertung · Audit. `merge completed ≠ historical identities deleted`; `same canonical node ≠ every source assertion becomes authoritative`.

## 18. Split

Split benötigt entsprechende Zuordnung der Assertions und Evidence References. `split completed ≠ original evidence discarded`.

## 19. Temporal Validity

`first observed · last observed · observation time · recording time · valid from · valid until · freshness · staleness · superseded · historical · unknown validity`.
```text
recently recorded  ≠ recently observed
last observed      ≠ currently valid indefinitely
not recently observed ≠ relationship absent
newer timestamp    ≠ higher authority
```
**Keine stille Last-Write-Wins-Regel.**

## 20. Current Views

`current calculated view · source-specific view · workspace view` (Detail §21). `current view ≠ complete reality`; `source-specific view ≠ canonical view`.

## 21. Historical Views and Snapshots

`historical view · point-in-time snapshot · comparison view · conflict view · unknown-completeness view`.
```text
snapshot created ≠ every source observed at the same instant
snapshot         ≠ immutable evidence automatically
```
Ein Snapshot benötigt Scope · Sources · Time Boundaries · Completeness · Confidence · Conflict · Evidence References.

## 22. Conflict Handling

Zusammenfassung (Detail Companion 2 §17-22): keine automatische Last-Write-Wins; keine Timestamp-Priorität allein; Konflikte bleiben sichtbar; Source Authority scopegebunden; manuelle Assertion löscht konkurrierende Evidence nicht; **ungelöste sicherheitsrelevante Konflikte blockieren privilegierte Automatisierung**. Keine Conflict-Resolution-Engine.

## 23. State-Authority Boundary

```text
topology assertion ≠ authoritative resource state
node present       ≠ resource currently available
edge present       ≠ relationship currently operational
derived path       ≠ desired or effective state
manual topology declaration ≠ execution authorization
```
[SoT](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) und [State-Modell](OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md) bleiben autoritativ; **kein Parallel-State-Modell**.

## 24. Event and Evidence Boundary

Topology Changes können Events erzeugen/Evidence referenzieren, **aber**: `topology difference ≠ incident automatically`; `new edge ≠ verified connectivity event`; `graph snapshot ≠ sufficient evidence automatically`; `manual assertion ≠ validated evidence`. Übergänge brauchen Klassifizierung, Provenance, Scope ([Event/Audit](EVENT_AND_AUDIT_CORRELATION_MODEL.md), [Evidence](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md)).

## 25. Execution Boundary

```text
topology graph        ≠ execution plan
dependency path       ≠ approval to remediate
connectivity assertion ≠ permission to configure network
node selection        ≠ authorised target scope
```
Topology unterstützt Planung/Analyse/Policy-Evaluation; jede Write-/Execution-Aktion bleibt an [CO-WP-013](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) gebunden. Unresolved/conflicted Identity oder Relationship Scope blockiert privilegierte automatische Ausführung.

## 26. Workspace and Disclosure

Workspace Scope bleibt an Nodes/Assertions gebunden; Cross-Workspace-Beziehungen brauchen explizite Governance; Netzwerk-/Standort-/Abhängigkeitsdaten sind sensitiv; Ressourcenexistenz nicht unautorisiert offengelegt; Correlation/Shared Nodes erzeugen keine globale Disclosure Authority; Export mit eigener Autorität; precise Location/personenbezogene Daten nur bei Notwendigkeit; consumer-safe/administrative Views getrennt (Detail Companion 3, Bezug THR-035, THR-018).

## 27. Offline Topology

Offline-Topology-Pakete folgen Companion 3 §23 (target-environment binding, provenance, integrity, import governance, conflict detection, reconciliation). Nicht behauptet: implementiert, konfliktfreie automatische Reconciliation, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-Technologie (Bezug THR-024).

## 28. Audit and Evidence

Erfasst: node/assertion identity · source/producer · workspace/scope · origin class · observation/recording time · identity-resolution decision · validation · confidence · evidence references · conflict · manual action · override · suppression · merge · split · export · offline import · supersession · closure. Trennung:
```text
graph record          ≠ validated relationship evidence
manual decision record ≠ proof of physical topology
export evidence       ≠ recipient authority
```

## 29. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Topology graph presence must not imply current resource availability.
2. Relationship assertions must not be treated as validated facts automatically.
3. Discovered, observed, declared, imported, manual and inferred assertions remain separate.
4. Same display name, address or alias must not imply the same canonical node.
5. Identity merge and split must preserve historical identity and evidence.
6. Timestamps must not provide silent last-write-wins conflict resolution.
7. Manual authority must remain human-attributable, scope-bound and reviewable.
8. Manual overrides must not delete competing observations or evidence.
9. Suppression from a view must not imply relationship absence.
10. Unresolved identity or relationship conflicts must block unsafe privileged automation.
11. Topology exports must preserve workspace, provenance and disclosure scope.
12. Offline topology import requires provenance, integrity, target binding and explicit governance.

## 30. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-002, THR-008, THR-010, THR-011, THR-012, THR-013, THR-014, THR-015, THR-016, THR-017, THR-018, THR-022, THR-024, THR-025, THR-026, THR-029, THR-030, THR-033, THR-034, THR-035, THR-037, THR-039, THR-040.

## 31. Technology Boundary

Nicht ausgewählt/implementiert: Graph-DB/Query-Language, Discovery Engine, CMDB, Netzwerk-/Agent-Protokoll, ID-Technologie, Identity-/Conflict-Resolution-Engine, Visualisierung/Layout, Graph-Schema, Runtime-Code.

## 32. Compatibility

Konsistent mit [Telemetry (topology observation)](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md), [Event/Audit](EVENT_AND_AUDIT_CORRELATION_MODEL.md)/[Evidence](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [SoT/Provenance/State](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Integration Contract](COREOPS_INTEGRATION_CONTRACT_V0_1.md), [Module Catalog](COREOPS_MODULE_CATALOG.md), [Data Ownership](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md), [Policy/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Konkretisiert die Topology-Aspekte von MOD-TOP und CAP-TOPOLOGY-*.

## 33. Open Questions

- Graph-/Discovery-Technologie (spätere ADR).
- Identity-Resolution-Kriterien im Detail.
- Precise-Location-Disclosure-Governance (CO-WP-025).

## 34. Next Decision

Companion 2 trägt Evidence/Confidence/Conflict, Companion 3 Manual Authority/Disclosure. Deployment Control Plane (CO-WP-021) referenziert deployment targets. Graph-/Discovery-Wahl bleibt einer späteren ADR-Runde vorbehalten.
