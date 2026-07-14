# CoreOps – Artifact Trust, Quarantine and Revocation Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation artifact trust, quarantine and revocation policy
> Implementation Status: Not implemented
> Signing Technology: Not selected
> Hash Technology: Not selected
> Trust Anchor: Not selected
> Scanner Technology: Not selected
> Revocation Distribution: Not implemented
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-022 (docs-only / artifact identity, software-supply-chain and trust-governance foundation)

## 1. Status

Technologieunabhängige Policy für **Artifact-Autorität, Quarantine, Trust Decision, Revocation/Reinstatement, Existing-Deployment-Response, Offline und Fail-Closed**. Companion zu [ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md](../architecture/ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md) und [ARTIFACT_DEPENDENCY_COMPATIBILITY_AND_DISTRIBUTION_MODEL.md](../architecture/ARTIFACT_DEPENDENCY_COMPATIBILITY_AND_DISTRIBUTION_MODEL.md). Kein Signing/Hash/Trust-Anchor/Scanner, keine Revocation-Distribution implementiert.

## 2. Purpose

Artifact-Trust ist die Supply-Chain-Grenze vor jedem Deployment: eine Trust Decision ist keine Execution Authorization, ein SBOM keine Sicherheit, eine Revocation keine automatische Entfernung. Diese Policy legt fail-closed-Regeln fest und warum `artifact available ≠ trusted`, `quarantine released ≠ deployment authorised` und `revocation → new deployment blocked unless explicit exception`.

## 3. Scope

Authority Model · Registration · Receipt/Quarantine · Identity Resolution · Source/Producer/Builder Trust · Provenance/Integrity Assessment · Validation · Trust Decision · SBOM Boundary · Component/Dependency Review · Vulnerability/Advisory Assessment · Compatibility · Distribution · Deployment Binding · Withdrawal/Revocation · Existing Deployment Response · Reinstatement · Exceptions · Offline · Delayed Revocation · Workspace Isolation · Audit · Failure/Fail-Closed.

## 4. Non-Goals

- Kein Signing-/Hash-/Trust-Anchor-/Scanner-/Registry-/Resolver-Mechanismus, keine Revocation-Distribution.
- Keine Behauptung implementierter/validierter/zertifizierter Trust-Kontrollen; kein Runtime-Code.

## 5. Authority Model

Getrennt: Artifact Owner · Validator · **Trust Decision Owner** · **Revocation Authority** · Distributor · Repository Operator · Deployment Consumer · Evidence Owner. `distributor ≠ trust authority`; `repository operator ≠ artifact owner`; `maintainer ≠ revocation authority automatically`. Trust Decisions/Deployment bleiben an [Execution Authorization (CO-WP-013)](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) gebunden — **keine parallele Autorität**.

## 6. Artifact Registration

Registrierung erfasst stabile Identität, Klasse, Source/Producer, Version/Revision; `registered ≠ trusted`. Unbekannte Identität/Source → Quarantine (§7).

## 7. Receipt and Quarantine

Quarantine gilt bei: unknown identity/source/provenance · integrity failure/conflict · malformed/incompatible metadata · revoked artifact · suspicious distribution · unresolved component identity · unknown offline package.
```text
quarantined        ≠ malicious proven
quarantine released ≠ deployment authorised
```
Quarantine benötigt Owner · Reason · Scope · Evidence · Review · Audit. Keine Quarantine-Technologie.

## 8. Identity Resolution

Auflösung bindet Identität/Version/Revision (Companion 2 §5); `mutable alias ≠ final privileged-deployment binding`; Identitätskonflikte → Quarantine (Bezug THR-014).

## 9. Source and Producer Trust

`producer stated ≠ producer verified`; Source Trust zeitgebunden; kompromittierte Source/Producer verlieren Trust, löschen aber historische Evidenz nicht (Bezug THR-023).

## 10. Builder and Build Context

Kompromittierter Builder/Build-Environment invalidiert Trust; `build metadata present ≠ provenance complete` (Bezug THR-023).

## 11. Provenance Assessment

Provenance-State ist explizit; `provenance available ≠ validated`; fehlende Provenance → Quarantine/fail-closed für privilegiertes Deployment.

## 12. Integrity Assessment

`integrity metadata present ≠ verified`; `integrity verified ≠ safe`; Integrity-Failure/-Conflict → Quarantine (Bezug THR-021, THR-022). Keine Hash-/Signing-Technologie.

## 13. Validation

Validation ist eigenständig (Companion 1 §16); `validated ≠ trusted ≠ supported`; scopegebunden.

## 14. Trust Decision

Eine Trust Decision (18 Felder): artifact identity/revision · decision owner · consumer/intended use · target class · workspace/environment · deployment profile · provenance/integrity/validation/SBOM/vulnerability/compatibility state · known limitations · exceptions · validity boundary · review trigger · evidence references · audit reference.
```text
trusted-for-one-use ≠ trusted globally
```
**Trust Decisions ersetzen keine Execution Authorization.**

## 15. SBOM Boundary

`SBOM available ≠ complete ≠ accurate ≠ artifact trusted`; `component absent from SBOM ≠ component absent`; `multiple SBOMs ≠ independent sources`. SBOM-Fehlen blockiert keine Bewertung, aber `no SBOM` ist ein dokumentierter Trust-Faktor.

## 16. Component and Dependency Review

Component-/Dependency-Review (Companion 2) ist Teil der Trust Decision; `same name/version ≠ same component/revision`; Dependency Confusion sichtbar (Bezug THR-023).

## 17. Vulnerability and Advisory Assessment

`vulnerability reference exists ≠ artifact affected`; `component affected ≠ deployment exploitable`; `no scanner finding ≠ no vulnerability`; `severity ≠ deployment-context risk`. Assessment ist deployment-context-gebunden; keine Scanner-Technologie.

## 18. Compatibility

Compatibility ist getrennt von Trust (Companion 2 §11); `trusted ≠ compatible with every target`.

## 19. Distribution

`available from approved repository ≠ approved for deployment`; `repository trusted ≠ every artifact trusted`. Distribution erhält Provenance/Trust/Revocation-State.

## 20. Deployment Binding

Deployment bindet konkrete Resolution-Ergebnisse (Companion 1 §24); vor Execution/Wave-Expansion **Recheck**: artifact resolution unchanged · trust decision current · no material revocation/withdrawal · provenance/integrity current · SBOM/advisory not materially changed · target compatibility valid · authorization scope matches.
```text
deployment plan approved ≠ any later artifact resolution approved
```

## 21. Withdrawal

`withdrawal · support withdrawal · compatibility withdrawal · distribution removal · trust suspension · deployment prohibition` getrennt. `artifact withdrawn ≠ necessarily malicious`; `repository removal ≠ revocation`.

## 22. Revocation

Revocation (Companion 2 §18) benötigt Authority/Reason/Scope/Effective boundary/Affected consumers/Impact/Existing-deployment assessment/Offline impact/Follow-up/Evidence/Audit. `revocation issued ≠ delivered to every offline environment`; **`revocation received → new deployment blocked unless explicit governed exception applies`**. Keine Revocation-Distribution-Technologie.

## 23. Existing Deployment Response

`artifact revoked ≠ automatic destructive removal`; `existing deployment ≠ permission for new deployment`. Bestehende Deployments benötigen gesonderte Impact-/Mitigation-/Recovery-/Exception-Bewertung (rollback-Artifact evtl. selbst revoked → Forward Recovery, CO-WP-021).

## 24. Reinstatement

Reinstatement (10 Felder): prior withdrawal/revocation reference · owner/authority · reason · new evidence · updated provenance/integrity/vulnerability state · target scope · validity boundary · review trigger · audit reference.
```text
reinstated ≠ historical revocation erased ≠ globally trusted ≠ deployment authorised
```

## 25. Exceptions

Eine Governance-Exception (z. B. Deployment trotz Revocation) ist explizit, begründet, scope-/zeitgebunden, human-approved (CO-WP-013), auditiert; keine stille Umgehung. Machine Principals können keine Exception self-granten.

## 26. Offline Artifact Transfer

Offline-Artifact-Pakete benötigen: artifact identities/revisions · component/SBOM references · source/target environment · workspace/scope · distribution context · provenance/integrity/validation state · trust decision · withdrawal/revocation state · validity/usage boundary · import quarantine · explicit import · deployment binding · result/reconciliation package · delayed revocation challenge · clock/sequence uncertainty. Nicht behauptet: implementiert, beliebige Air-Gap-Stufen, konfliktfreie automatische Reconciliation, Klassifiziertnetz-Eignung, konkrete Signing-/Hash-/Trust-Anchor-Technologie (Bezug THR-024). Fail-closed bei unklarer Provenance/Revocation.

## 27. Delayed Revocation

`revocation not yet delivered · artifact already deployed/staged/running · rollback artifact also revoked · offline target unreachable · replacement incompatible/absent`. Delayed Revocation sichtbar; neues Deployment blockiert bis explizite Exception; bestehende Deployments gesondert bewertet.

## 28. Workspace Isolation

Artifact-Trust/-Revocation ist workspace-/environment-gebunden; Cross-Workspace-Trust-Übertragung nur mit expliziter Autorität; keine unautorisierte Offenlegung privater Repos/Artifacts (Bezug THR-035, THR-019).

## 29. Audit and Evidence

Erfasst: artifact registration · resolution · receipt · quarantine · provenance/integrity assessment · validation · trust decision · SBOM registration · component/dependency decision · vulnerability/advisory assessment · withdrawal · revocation · reinstatement · distribution · offline import · deployment binding · exception · closure. `artifact evidence ≠ trust decision ≠ execution authorization`; Historie nicht umgeschrieben (Bezug THR-016, THR-017).

## 30. Failure and Unknown State

`source/provenance/integrity unknown · SBOM missing/incomplete · component identity unresolved · revocation delivery unknown · offline reconciliation incomplete`. `missing information ≠ safe`; `unknown ≠ trusted`.

## 31. Fail-Closed Rules

Keine privilegierte Deployment-Bindung/-Nutzung bei: unbekannter Identität/Source/Provenance · Integrity-Failure/-Conflict · unresolved Component/Identity-Konflikt · aktiver/unklarer Revocation ohne Exception · fehlender Trust Decision für Ziel/Scope · unbekanntem Offline-Provenance · fehlendem Audit-Start · Alias-only-Binding für privilegiertes Deployment.

## 32. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Artifact availability must not imply trust, integrity, compatibility or deployment authorization.
2. Source, producer, builder, distributor, repository operator and trust decision owner remain separate roles.
3. Mutable aliases must not serve as final privileged-deployment binding.
4. SBOM existence must not imply completeness, accuracy or artifact trust; missing information proves no absence.
5. Vulnerability reference does not automatically imply artifact or deployment exploitability.
6. Trust decisions remain use-, target-, scope-, version- and time-bound and do not replace execution authorization.
7. Quarantine release must not imply deployment authorization.
8. Withdrawal, revocation, repository removal and support withdrawal remain separate.
9. Revocation blocks new deployment unless an explicit governed exception applies; existing deployment is not permission for new deployment.
10. Reinstatement preserves historical revocation evidence and does not imply global trust.
11. Offline artifact transfer requires provenance, integrity, target binding, bounded trust and explicit import governance.
12. Machine principals must not self-grant trust exceptions.

## 33. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-008, THR-010, THR-013, THR-014, THR-016, THR-017, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024, THR-026, THR-035. Keine Duplikation, kein Parallelregister.

## 34. Technology Boundary

Nicht ausgewählt/implementiert: Signing/Hash/Trust-Anchor/Transparency Log, Scanner, Registry/Repository, Dependency Resolver, Revocation-Distribution, Runtime-Code.

## 35. Compatibility

Konsistent mit Artifact-Model-/Dependency-Companion, [Deployment](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Domain Pack Trust](DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md), [Integration Trust](INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [Audit Policy](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md). Konkretisiert DEC-P-02, DEC-G-05, DEC-G-06 und die Supply-Chain-Sovereignty-Linie.

## 36. Open Questions

- Signing-/Integrity-/Trust-Anchor-Mechanismus (spätere ADR, mit CO-WP-024).
- Revocation-Distribution/Delayed-Revocation im Detail (mit CO-WP-023).
- Deployment-Context-Exploitability-Bewertung.

## 37. Next Decision

Restricted/Air-Gapped Operation (CO-WP-023), Secrets/Key Custody (CO-WP-024) und Self-Protection/Recovery (CO-WP-026) konkretisieren Offline-/Signing-/Revocation-Aspekte. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
