# CoreOps – Offline Trust, Activation, Revocation and Transfer Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation offline trust, activation, revocation and transfer policy
> Implementation Status: Not implemented
> Signing Technology: Not selected
> Hash Technology: Not selected
> Trust Anchor: Not selected
> Transfer Technology: Not selected
> Offline Activation: Not implemented
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-023 (docs-only / restricted-operation, offline-distribution and CorePack governance foundation)

## 1. Status

Technologieunabhängige Policy für **Offline-Trust, Transfer-Autorität, Import-Quarantine, Assessment, Approval, Activation, Revocation/Reinstatement, Rollback/Recovery, Evidence Return und Fail-Closed** von CorePacks. Companion zu [RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md](../architecture/RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md) und [COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md). Kein Signing/Hash/Trust-Anchor/Transfer, keine Offline-Activation implementiert.

## 2. Purpose

Offline-Betrieb senkt nie die Governance-Anforderungen. Diese Policy legt fail-closed-Regeln fest und warum `network unavailable ≠ security controls optional`, `imported ≠ trusted`, `trusted ≠ approved for activation`, `quarantine released ≠ activation authorised`, `activation approved ≠ deployment authorised` und `offline authorization ≠ reusable authorization`.

## 3. Scope

Authority Model · CorePack/Content Trust · Target Binding · Transfer Authority · Receipt · Import Quarantine/Assessment · Approval · Activation/Execution Authorization · Policy/Identity Freshness · Revocation Snapshot/Delayed Revocation · Exceptions · Updates/Deltas · Partial/Unknown Outcome · Rollback/Recovery · Restricted/Degraded Operation · Evidence Return · Disclosure/Export · Workspace Isolation · Audit · Failure/Fail-Closed.

## 4. Non-Goals

- Kein Signing-/Hash-/Trust-Anchor-/Transfer-/Removable-Media-/Encryption-Mechanismus, keine Reconciliation-Engine, keine Offline-Activation.
- Keine Behauptung implementierter/validierter/zertifizierter Offline-Kontrollen; kein Runtime-Code; keine Air-Gap-/Compliance-Eignung; keine Eignung für klassifizierte Netze.

## 5. Authority Model

Getrennt: `central policy authority` · `local delegated policy authority` · `central approval authority` · `local approval authority` · `central trust authority` · `local bounded trust authority` · `central revocation authority` · `local emergency suspension authority` · `activation authority` · `deployment execution authority` · `export authority`.

```text
local administrator           ≠ unrestricted local policy authority
offline operation             ≠ permanent delegation
local activation authority    ≠ deployment execution authorization
central authority unavailable ≠ local authority expands automatically
```

Autorität bindet an [Policy Decision](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Approval Lifecycle](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md) und [Execution Authorization](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) (CO-WP-013) — **keine parallele Offline-Autorität**.

## 6. CorePack Trust Boundary

`CorePack imported ≠ trusted`; `trusted CorePack ≠ approved for activation`. Trust ist eine bewusste, use-/target-/scope-/version-/revision-/zeitgebundene Entscheidung, getrennt von Provenance/Integrity/Validation/Compatibility/Support ([Companion 2 §17](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md)).

## 7. Content Trust

CorePack-Contents erben **keinen** Trust vom Container: `CorePack integrity verified ≠ every content item trusted`. Enthaltene Artifacts/Domain Packs/Blueprints unterliegen ihren eigenen Trust-Grenzen ([CO-WP-022](ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md), [CO-WP-015](DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md)). Content-Trust ist an die konkrete gebundene Revision geknüpft, nicht an einen Mutable Alias.

## 8. Target Binding

Trust/Approval/Activation eines CorePacks sind an ein konkretes Target Binding gebunden (Companion 2 §19). `CorePack valid for one target ≠ valid for another isolated environment`. Target Binding bleibt durch Transfer, Import und Activation erhalten und wird nicht durch lokalen Rename ersetzt (Bezug THR-035).

## 9. Transfer Authority

Transfer benötigt benannte Transfer-Autorität, Quell-/Zielumfeld, CorePack-Identity/Revision und Target Binding. `handling history ≠ legal chain of custody`. Transfer autorisiert weder Import noch Activation. Keine Transfer-/Removable-Media-Technologie (Bezug THR-022, THR-039).

## 10. Receipt

Receipt ist ein eigener Zustand: `received ≠ imported ≠ trusted ≠ activated`. Receipt bestätigt Empfang, nicht Vertrauen. `recent receipt ≠ recent assembly`.

## 11. Import Quarantine

Quarantine gilt bei: unknown identity/revision · target mismatch · manifest missing/conflicted · provenance unknown · integrity failure/unknown · revocation snapshot too stale · unsupported compatibility · unexpected contents · partial transfer · duplicate/replayed transfer (Bezug THR-026, THR-024).

```text
quarantine release ≠ activation authorization
```

Quarantine benötigt Owner · Reason · Scope · Evidence · Review · Audit.

## 12. Import Assessment

Assessment prüft: CorePack identity/version/revision · target binding · content population · manifest consistency · content resolutions · provenance · integrity · validation · trust · compatibility · policy/schema versions · revocation freshness · time uncertainty · known limitations · duplicate/replay context · local authority.

```text
import assessment passed ≠ activation approved
```

## 13. Approval

Approval ist explizit, zurechenbar, scope-/zeitgebunden und widerrufbar ([CO-WP-013](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md)). Machine Principals können nicht self-approven (Bezug THR-038). Approval bezieht sich auf ein konkretes CorePack, Target Binding und einen Effective Content Set.

## 14. Activation Authority

Activation Authority autorisiert die Aktivierung eines bewerteten, approbierten CorePacks in einem gebundenen Target. `activation authority ≠ deployment execution authorization`; `activation approved ≠ every contained deployment authorised`; `CorePack activated ≠ every content item installed/executed`; `activation completed ≠ desired state verified`.

## 15. Execution Authorization

Wo Activation ein Deployment/eine privilegierte Aktion auslöst, gilt zusätzlich [Execution Authorization (CO-WP-013)](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) und die [Deployment Control Plane (CO-WP-021)](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md) — action-/target-/scope-/plan-/zeitgebunden. Activation ersetzt Execution Authorization nicht.

## 16. Policy and Identity Freshness

`policy snapshot present ≠ policy current centrally`; `offline authorization ≠ reusable authorization`. Snapshots tragen Version, Scope, Validity Boundary, Consumption State. Offline Identity nutzt Identity-/Rollen-Snapshots mit expliziter Revocation-Freshness; Machine Principals imitieren keine Human Approval/Manual Authority (Bezug THR-001, THR-038). Verbrauchte/abgelaufene/widerrufene Autorisierung wird nicht wiederverwendet (Bezug THR-004, THR-026).

## 17. Revocation Snapshot

Ein Revocation Snapshot trägt Scope, Alter, Validity Boundary und bekannte Delivery-Lücke. Getrennte Revocation-Arten: content-artifact · CorePack · Domain-Pack · policy withdrawal · activation prohibition · distribution prohibition · local emergency suspension. `one contained artifact revoked ≠ CorePack automatically safe ≠ CorePack automatically wholly revoked`. Eine Impact-Bewertung bestimmt Contents, Targets, Activations und Existing Deployments.

## 18. Delayed Revocation

Offline gilt: last revocation snapshot · snapshot scope · snapshot age · validity boundary · known delivery gap · local suspension state · pending central confirmation · existing activation impact · new activation rule · exception state · audit reference.

```text
no local revocation entry ≠ not revoked centrally
```

Nach Ablauf der Freshness-Grenze wird neue privilegierte Activation — profilabhängig — blockiert, beschränkt oder benötigt eine eng begrenzte Exception (§19). Bestehende Activations werden gesondert bewertet (Bezug THR-027, THR-013).

## 19. Exceptions

Eine Governance-Exception (z. B. Activation trotz zu alter Revocation-Freshness oder trotz Revocation) ist explizit, begründet, scope-/zeitgebunden, human-approved (CO-WP-013) und auditiert. Keine stille Umgehung; Machine Principals können keine Exception self-granten.

## 20. Updates and Deltas

`delta available ≠ target baseline confirmed`; `delta applied ≠ resulting CorePack verified`; `same base version ≠ same base revision`. Delta-Activation erfordert eine **bestätigte exakte Ausgangsrevision**; unbekannte/konflikthafte Baseline blockiert automatische Delta-Aktivierung (Companion 2 §27).

## 21. Partial Activation

Per-Content-/Per-Target-Ergebnisse bleiben sichtbar: activated · skipped · rejected · failed · outcome-unknown · rollback-pending · recovery-pending. `partial activation ≠ complete success`. Historie bleibt erhalten (Bezug THR-031).

## 22. Unknown Outcome

`unknown activation outcome → no automatic retry → reconciliation required`. `activation failed ≠ no side effects`; `missing return package ≠ no operation occurred` (Bezug THR-029, THR-028).

## 23. Rollback

Rollback erfordert aktuelle Trust-/Revocation-/Compatibility-Bewertung: `previously activated ≠ currently safe rollback target`. Ein Rollback-Ziel-CorePack kann selbst revoked/inkompatibel sein → Forward Recovery statt unsicherem Rollback (Bezug THR-032).

## 24. Recovery

Recovery Packs erhalten keine automatische globale Recovery Authority; Recovery ist scope-/target-/zeitgebunden, approbiert und auditiert. Recovery erhält historische Evidenz und bestätigt Ergebnis über Verifikation, nicht über Annahme (Bezug THR-033).

## 25. Restricted and Degraded Operation

Bei eingeschränkter Freshness/Autorität gelten degradierte Modi (Companion 1 §14): read-only/evidence-only/recovery-only/activation-blocked/deployment-blocked/trust-review-required. Ein degradierter Modus ist benannt, begrenzt, zurechenbar, reviewbar und erweitert keine Autorität. `system running ≠ fully governed operation available`.

## 26. Evidence Return

Result-/Evidence-Return-Packages tragen Provenance/Integrity/Time-Uncertainty und known gaps. `evidence returned ≠ evidence complete ≠ evidence sufficient`. Keine automatische Reconciliation-Engine; unbekannte Ausgänge lösen explizite Reconciliation aus.

## 27. Disclosure and Export

```text
local read access             ≠ export authority
CorePack activation authority ≠ evidence-export authority
```

Ausgehende Pakete benötigen eigene Export-Autorität und enthalten keine Rohsecrets, wiederverwendbaren Credentials oder unnötigen sensitiven Daten (Bezug THR-020, THR-025).

## 28. Workspace and Environment Isolation

Trust/Approval/Activation sind workspace-/environment-gebunden. Ein Pack für Environment A wird nicht still in Environment B aktiviert; Cross-Workspace-Inhalte benötigen explizite Governance; geteilte Transfer-Infrastruktur erzeugt keine gemeinsame Autorität; Import/Activation erweitern Scope nicht (Bezug THR-035).

## 29. Audit and Evidence

Erfasst: transfer · receipt · import · quarantine · assessment · approval · activation · partial results · update/delta · revocation · reinstatement · rollback · recovery · evidence return · exception · closure. `offline audit record ≠ trust decision ≠ activation authorization`; Historie wird nicht umgeschrieben (Bezug THR-016, THR-017).

## 30. Failure and Unknown State

`partial assembly · manifest incomplete · transfer interrupted · receipt incomplete · import failed · integrity unknown · provenance unknown · revocation freshness unknown · target binding conflicted · activation partial · activation outcome unknown · evidence return incomplete · reconciliation conflicted`.

```text
transfer failed        ≠ source CorePack invalid automatically
import failed          ≠ no content entered target environment
activation failed      ≠ no side effects
missing return package ≠ no operation occurred
```

`missing information ≠ safe`; `unknown ≠ trusted`.

## 31. Fail-Closed Rules

Keine privilegierte Activation/Deployment bei: unbekannter CorePack-Identity/Revision · Target Mismatch · unbekannter/failed Provenance/Integrity · zu alter/unklarer Revocation-Freshness ohne Exception · fehlender Trust Decision für Ziel/Scope · fehlender/abgelaufener/konsumierter Authorization · unbestätigter Delta-Baseline · unbekanntem Offline-Provenance · fehlendem Audit-Start · Alias-only-Binding für privilegierte Activation · Scope-Erweiterung durch Import/Activation.

## 32. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Offline operation must not expand central or local authority automatically.
2. CorePack transfer, receipt, import, trust, approval and activation remain separate states.
3. CorePack contents do not inherit trust, validation or compatibility from the container automatically.
4. Mutable aliases must not serve as final CorePack content binding.
5. Target binding must be preserved through transfer, import and activation.
6. Quarantine release must not imply activation authorization.
7. Activation authority must not imply deployment execution authorization.
8. Delta activation requires a confirmed exact base revision.
9. Delayed revocation and time/clock uncertainty remain explicit; stale revocation freshness must not permit new privileged activation without a bounded exception.
10. Partial and unknown activation outcomes remain visible and block unsafe retry.
11. Offline authorization remains action-, target-, scope-, version-, purpose- and time-bound and is not reusable.
12. Evidence-return and export require their own authority and must not include secrets; rollback and recovery require current trust, revocation and compatibility assessment.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 33. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](THREAT_SCENARIO_REGISTER.md): manipuliertes Offline-Paket THR-021/THR-022/THR-024; falsches Environment THR-022/THR-035; Replay/Duplicate Import THR-026/THR-004; stale Policy/Revocation THR-013/THR-027; kompromittiertes Artifact/Dependency THR-021/THR-022/THR-023; Identifier Confusion THR-014; Privilege Escalation THR-002; Secret Exposure THR-019/THR-020/THR-025; Audit Manipulation THR-016/THR-017; Partial Failure THR-031/THR-032; Result Misrepresentation THR-028/THR-029/THR-030; Resource gegen CoreOps THR-034/THR-036; Machine-Imitation THR-038/THR-001. Keine erfundenen IDs.

## 34. Technology Boundary

Nicht ausgewählt: Signing · Hash · Trust Anchor · PKI · Encryption · Transfer-/Removable-Media-Technologie · Installer · Update-Engine · Synchronisations-/Reconciliation-Engine · Manifest-/Pack-Format · Offline-Runtime/-Activation. Alle bleiben `deferred`.

## 35. Compatibility

Konsistent mit [DEC-P-02](../../project-system/DECISION_INDEX.md) (Offline First), [CO-WP-013](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [CO-WP-021](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [CO-WP-022](ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md) und den Offline-Grenzen aus [CO-WP-011](../architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md)/[CO-WP-016](DATA_MIGRATION_INTEGRITY_AND_RECOVERY_POLICY.md)/[CO-WP-018](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md). Keine bestehende Foundation-Invariante wird geschwächt.

## 36. Open Questions

- Profilabhängige Freshness-Schwellen für Policy-/Revocation-Snapshots (deferred).
- Verhältnis lokaler Emergency Suspension zu zentraler Reinstatement-Reihenfolge (Companion 2 §29).
- Exception-Governance-Grenzen für Delta-Activation ohne festgelegte Baseline-Bestätigungstechnologie (deferred).

## 37. Next Decision

`CO-WP-024 – Secrets, Configuration Vault and Key Custody` gemäß lokaler [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md). Zuerst Nova Review von CO-WP-023, danach Human-Maintainer-Commit.
