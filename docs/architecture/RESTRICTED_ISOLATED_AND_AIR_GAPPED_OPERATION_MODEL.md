# CoreOps – Restricted, Isolated and Air-Gapped Operation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation restricted, isolated and air-gapped operation model
> Implementation Status: Not implemented
> Transfer Technology: Not selected
> Offline Runtime: Not selected
> Trust Bootstrap Technology: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-023 (docs-only / restricted-operation, offline-distribution and CorePack governance foundation)

## 1. Status

Technologieunabhängiges Modell für den Betrieb von CoreOps unter **eingeschränkter, intermittierender, isolierter und air-gapped Konnektivität** sowie in Recovery-Szenarien. Companion zu [COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md](COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md) und [OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md](../security/OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md). Kein Transferdienst, keine Offline-Runtime, kein Trust-Bootstrap-Mechanismus implementiert.

## 2. Purpose

CoreOps ist offline-/air-gap-fähig als Produktrichtung ([DEC-P-02](../../project-system/DECISION_INDEX.md)). Dieses Modell legt fest, wie Konnektivitätsklassen, Autorität, Identität, Freshness und degradierte Betriebsmodi getrennt bleiben, und warum `offline ≠ air-gapped`, `isolated ≠ trusted`, `network unavailable ≠ security controls optional` und `central authority unavailable ≠ local authority expands`. Fehlende Konnektivität senkt niemals still die Governance-Anforderungen.

## 3. Scope

Operational Connectivity Classes · Authority Boundary (central/local) · Offline Identity · Policy/Authorization Freshness · Time/Clock Uncertainty · Restricted Operation · Degraded Modes · Import Boundary · Activation Boundary · Evidence Return · Reconciliation · Workspace/Environment Isolation · Failure/Unknown State · Security Invariants.

## 4. Non-Goals

- Kein Archiv-/Paketformat, kein Installer, keine Update-Engine, kein Transferdienst, keine Removable-Media-Technologie, keine Synchronisations-/Reconciliation-Engine, keine Signing-/Hash-/PKI-/Trust-Anchor-/Encryption-Technologie.
- Keine Behauptung implementierter Offline-Installation/Aktivierung, keine Air-Gap-/Security-/Compliance-Eignung, keine Eignung für klassifizierte Netze.
- Kein Runtime-Code; keine tatsächliche Offline-Operation.

## 5. Concepts

Mindestbegriffe: `restricted-connected operation`, `intermittently-connected operation`, `isolated operation`, `air-gapped operation`, `offline operation`, `local authority`, `central authority`, `CorePack`, `transfer`, `import`, `quarantine`, `assessment`, `approval`, `activation`, `recovery`, `result package`, `evidence-return package`, `trust snapshot`, `revocation snapshot`, `freshness assessment`, `clock uncertainty`.

```text
offline               ≠ air-gapped automatically
isolated              ≠ trusted automatically
network unavailable   ≠ security controls optional
central unavailable   ≠ local authority expands automatically
system running        ≠ fully governed operation available
```

CorePack, Transfer, Import und Activation sind vollständig in [COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md](COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md) modelliert; dieses Dokument definiert die Betriebsumgebung, in der diese Einheiten wirken.

## 6. Connectivity Classes

Getrennte Betriebsbedingungen (keine Eignungsaussage für klassifizierte Netze):

| Klasse | Kurzcharakter |
| ------ | ------------- |
| `connected` | zentrale Autorität regulär erreichbar |
| `restricted-connected` | eingeschränkte, gefilterte oder unidirektionale Pfade |
| `intermittently-connected` | zeitweise Konnektivität, Lücken erwartbar |
| `isolated` | keine reguläre Verbindung, kontrollierter Transfer |
| `air-gapped` | keine Netzverbindung, ausschließlich physischer/manueller Transfer |
| `recovery-only` | eingeschränkter Not-/Wiederherstellungsbetrieb |
| `unknown-connectivity` | Konnektivität nicht bestimmt → konservativ behandelt |

Je Klasse dokumentiert das Betriebsprofil: Connectivity expectation · Permitted inbound paths · Permitted outbound paths · Central-authority availability · Local-authority expectation · Identity and policy freshness · Revocation freshness · Time uncertainty · Evidence-return expectation · Update approach · Known limitations.

`connected ≠ restricted-connected ≠ intermittently-connected ≠ isolated ≠ air-gapped` — jede Klasse hat eigene Freshness-, Autorisierungs- und Evidence-Erwartungen. Keine Klasse behauptet Compliance oder Eignung für klassifizierte oder VS-Netze.

## 7. Authority Boundary

Getrennt: `central policy authority` · `local delegated policy authority` · `central approval authority` · `local approval authority` · `central trust authority` · `local bounded trust authority` · `central revocation authority` · `local emergency suspension authority` · `activation authority` · `deployment execution authority` · `export authority`.

```text
local administrator            ≠ unrestricted local policy authority
offline operation              ≠ permanent delegation
local activation authority     ≠ deployment execution authorization
central authority unavailable  ≠ local authority expands automatically
```

Delegation bleibt scope-, purpose-, profile-, version- und zeitgebunden. Sie erweitert weder Ziel, Aktion noch Scope über das ausdrücklich Delegierte hinaus. Autorität bindet an [Policy Decision (CO-WP-013)](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Approval Lifecycle (CO-WP-013)](../security/APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md) und [Execution Authorization (CO-WP-013)](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) — **keine parallele Offline-Autoritätsstruktur**.

## 8. Central Authority

Zentrale Autorität ist der reguläre Ursprung für Policy, Approval, Trust und Revocation. Ihre Unerreichbarkeit ist ein bekannter Betriebszustand, kein Freibrief: konsumierte, abgelaufene oder widerrufene Autorisierung bleibt ungültig; ein fehlender lokaler Revocation-Eintrag beweist keine zentrale Nicht-Revocation (`no local revocation entry ≠ not revoked centrally`). Zentrale Autorität delegiert ausdrücklich und zeitgebunden; sie delegiert nicht implizit durch Konnektivitätsverlust.

## 9. Local Delegated Authority

Lokale delegierte Autorität handelt nur innerhalb eines expliziten, zeitgebundenen Delegationsprofils (Policy Snapshot, Scope, Ziel, Ablauf, erlaubte Aktionen). `local administrator ≠ unrestricted local policy authority`. Lokale Emergency Suspension (defensiv, restriktiv) ist zulässig; lokale Ausweitung von Trust oder Deployment-Autorität ist es nicht. Break-Glass bleibt außergewöhnlich, benannt, zurechenbar, auditiert und ablaufpflichtig ([CO-WP-009](../security/WORKSPACE_RBAC_AND_SCOPE_MODEL.md)) — kein Parallelbetriebsmodus.

## 10. Offline Identity

Berücksichtigt: local human identity · machine identity · identity snapshot · role/membership snapshot · credential reference · revocation freshness · accountable owner · break-glass authority · identity conflict · delayed central changes. Identitäts-/Rollen-Snapshots sind zeitpunktbezogen und können veraltet sein; Freshness bleibt explizit. Machine Principals dürfen **keine Human Approval oder Manual Authority imitieren** (Bezug THR-038, THR-001). Keine Credential-/Identity-Technologie ausgewählt (verweist auf [Machine Identity (CO-WP-010)](../security/MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [Human Identity (CO-WP-009)](../security/HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md)).

## 11. Policy and Authorization Freshness

`policy snapshot present ≠ policy current centrally`; `offline authorization ≠ reusable authorization`. Jeder Policy-/Authorization-Snapshot trägt Version, Scope, Validity Boundary und Consumption State. Nach Ablauf der Freshness-Grenze wird neue privilegierte Aktivierung — abhängig vom Profil — blockiert, beschränkt oder benötigt eine eng begrenzte, human-approved Exception. Verbrauchte, abgelaufene oder widerrufene Autorisierung wird nicht wiederverwendet (Bezug THR-004, THR-026).

## 12. Time and Clock Uncertainty

Getrennt: central time · local time · package assembly time · transfer time · receipt time · import time · activation time · last synchronization · freshness assessment · clock uncertainty.

```text
local clock appears valid ≠ central policy or revocation state current
recent receipt            ≠ recent assembly
```

Zeitunsicherheit bleibt sichtbar und fließt in jede Freshness- und Revocation-Bewertung ein (Bezug THR-027). Keine Clock-/Synchronisationstechnologie ausgewählt.

## 13. Restricted Operation

Restricted Operation ist ein bewusst begrenzter Betrieb bei eingeschränkter Autorität/Freshness. Er verengt Fähigkeiten, statt Autorität zu erweitern: read-only bevorzugt vor write, keine privilegierte Aktivierung ohne aktuelle Autorisierung, keine stille Scope-Erweiterung. `system running ≠ fully governed operation available`.

## 14. Degraded Modes

Betriebszustände: `fully operational` · `restricted operation` · `read-only operation` · `evidence-only operation` · `recovery-only operation` · `activation-blocked` · `deployment-blocked` · `trust-review-required` · `unknown operational state`. Ein degraded mode benötigt: reason · owner · permitted capabilities · prohibited capabilities · workspace/scope · validity boundary · review trigger · audit expectation · exit condition. Ein degradierter Modus ist benannt, begrenzt, zurechenbar und reviewbar; er ist kein dauerhafter Ersatz für governten Vollbetrieb.

## 15. Import Boundary

Import in eine restricted/isolated/air-gapped Umgebung erfolgt ausschließlich über die CorePack-Kette (Transfer → Receipt → Quarantine → Assessment → Approval → Activation; [Companion](COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md)). `received ≠ imported ≠ trusted ≠ activated`. Import darf Scope nicht erweitern und Target Binding nicht durch lokalen Rename ersetzen (Bezug THR-024, THR-014). Ein Pack für Environment A wird nicht still in Environment B importiert.

## 16. Activation Boundary

Activation setzt bestandene Assessment und explizite Approval voraus; `import assessment passed ≠ activation approved`; `quarantine release ≠ activation authorization`; `activation approved ≠ every contained deployment authorised`. Deployment innerhalb einer Activation bindet weiterhin an [Execution Authorization (CO-WP-013)](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) und [Deployment Control Plane (CO-WP-021)](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md). Details in [OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md](../security/OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md).

## 17. Evidence Return

Ergebnisse aus isolierten Umgebungen kehren als Result- bzw. Evidence-Return-Package zurück (source/destination environment · CorePack/activation references · per-content/per-target results · unknown outcomes · verification evidence · audit records · local exceptions · policy/trust snapshots · redaction/disclosure state · provenance · integrity · time uncertainty · known gaps). `evidence returned ≠ evidence complete ≠ evidence sufficient`. Ausgehende Pakete benötigen eigene Export-Autorität (`local read access ≠ export authority`) und enthalten keine Rohsecrets/wiederverwendbaren Credentials (Bezug THR-020, THR-025).

## 18. Reconciliation

Reconciliation gleicht zurückgeführte Ergebnisse mit dem zentralen Zustand ab. Unbekannte Ausgänge führen nicht zu automatischem Retry, sondern zu expliziter Reconciliation. `missing return package ≠ no operation occurred` (Bezug THR-028, THR-029). Keine automatische Reconciliation-Engine ausgewählt.

## 19. Workspace and Environment Isolation

Workspace und Environment bleiben durch CorePack, Transfer, Import und Activation erhalten. Ein Pack für Environment A darf nicht still in Environment B aktiviert werden; Cross-Workspace-Inhalte benötigen explizite Governance; Target Binding wird nicht durch lokalen Rename ersetzt; geteilte Transfer-Infrastruktur erzeugt keine gemeinsame Autorität; Import/Activation erweitern Scope nicht (Bezug THR-035).

## 20. Failure and Unknown State

`partial assembly · manifest incomplete · transfer interrupted · receipt incomplete · import failed · integrity unknown · provenance unknown · revocation freshness unknown · target binding conflicted · activation partial · activation outcome unknown · evidence return incomplete · reconciliation conflicted`.

```text
transfer failed        ≠ source CorePack invalid automatically
import failed          ≠ no content entered target environment
activation failed      ≠ no side effects
missing return package ≠ no operation occurred
```

`unknown ≠ safe`; unklare Zustände lösen fail-closed und Reconciliation aus (Bezug THR-031, THR-032).

## 21. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Offline operation must not expand central or local authority automatically.
2. Connectivity classes remain separate operating conditions with distinct freshness and authorization expectations.
3. Central authority unavailability must not implicitly delegate or expand local authority.
4. Delegated local authority remains scope-, purpose-, profile-, version- and time-bound.
5. Offline authorization remains action-, target-, scope-, version-, purpose- and time-bound and is not reusable.
6. Identity, policy and revocation freshness and clock uncertainty remain explicit.
7. Import must not expand scope, and activation must not authorise unbound deployment.
8. Target binding must be preserved through transfer, import and activation.
9. Degraded and restricted modes remain named, bounded, attributable and reviewable.
10. Evidence-return and export require their own authority and must not include secrets.
11. Workspace and environment isolation must survive transfer, import and activation.
12. Unknown operational, transfer, import or activation state must fail closed and require reconciliation, not silent retry.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 22. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md) (keine erfundenen IDs): manipuliertes Offline-Paket THR-021/THR-022/THR-024; falsches Target Environment THR-022/THR-035; Replay/Duplicate Import THR-026/THR-004; stale Policy/Revocation THR-013/THR-027; kompromittiertes Artifact THR-021/THR-022/THR-023; Identifier Confusion THR-014; Privilege Escalation THR-002; Secret Exposure THR-019/THR-020/THR-025; Audit Manipulation THR-016/THR-017; Partial Failure THR-031/THR-032; Result Misrepresentation THR-028/THR-029/THR-030; Resource gegen CoreOps THR-034/THR-036; Machine-Imitation THR-038/THR-001.

## 23. Technology Boundary

Nicht ausgewählt: Archiv-/Paketformat · Manifestformat · Installer · Update-Engine · Transfer Appliance/-Dienst · Removable-Media-Technologie · Signing/Hash/PKI/Trust Anchor/Encryption · Clock-/Synchronisations-/Reconciliation-Engine · Offline-Runtime. Alle bleiben `deferred` bis zu einer späteren, ADR-gestützten Entscheidung.

## 24. Compatibility

Konsistent mit [DEC-P-02](../../project-system/DECISION_INDEX.md) (Offline First), [Deployment Control Plane (CO-WP-021)](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [Artifact Trust (CO-WP-022)](../security/ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md), [Policy/Approval/Execution (CO-WP-013)](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) und den Offline-Grenzen aus [CO-WP-011](SOURCE_OF_TRUTH_AND_FIELD_PROVENANCE_MODEL.md)/[CO-WP-016](DATA_OWNERSHIP_AND_PERSISTENCE_MODEL.md). Keine bestehende Foundation-Invariante wird geschwächt.

## 25. Open Questions

- Konkrete Freshness-Grenzwerte je Connectivity Class (deferred, profilabhängig).
- Genaue Exit-Kriterien und Autoritätsanforderungen je Degraded Mode (deferred).
- Verhältnis lokaler Emergency Suspension zu zentraler Reinstatement-Reihenfolge (Companion Policy §17–§18).

## 26. Next Decision

`CO-WP-024 – Secrets, Configuration Vault and Key Custody` gemäß lokaler [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md). Zuerst Nova Review von CO-WP-023, danach Human-Maintainer-Commit.
