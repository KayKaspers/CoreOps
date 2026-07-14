# CoreOps – Artifact Dependency, Compatibility and Distribution Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation artifact dependency, compatibility and distribution model
> Implementation Status: Not implemented
> Artifact Registry: Not selected
> Dependency Resolver: Not selected
> Vulnerability Scanner: Not selected
> Compatibility Validation: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-022 (docs-only / artifact identity, software-supply-chain and trust-governance foundation)

## 1. Status

Technologieunabhängiges Modell für **Artifact Resolution, Component/Dependency Relationships, Compatibility, Vulnerability-Kontext, Withdrawal/Revocation-Distribution und Existing Deployments**. Companion zu [ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md](ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md) und [ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md](../security/ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md). Keine Registry, kein Dependency Resolver, kein Scanner.

## 2. Purpose

Auflösung und Abhängigkeiten sind der Ort für Dependency Confusion und Version-Drift: `same version label ≠ same revision`, `repository trusted ≠ every contained artifact trusted`, `vulnerability reference ≠ deployment exploitable`. Dieses Modell bindet privilegiertes Deployment an aufgelöste Identität/Version/Revision und trennt Vulnerability von Artifact- und Deployment-Impact.

## 3. Scope

Artifact Resolution · Version/Revision · Component Identity · Dependency Classes/Relationships/Scope · Compatibility · Target Constraints · Distribution/Repository Context · Validation · Vulnerability/Advisory · Withdrawal/Revocation/Replacement · Existing Deployments · Offline Distribution · Delayed Revocation · Audit.

## 4. Non-Goals

- Keine Registry/Repository/Dependency-Resolver/Scanner-Technologie, kein SBOM-Format.
- Keine Versionsnotation, keine Behauptung durchgeführter Auflösung/Validierung.

## 5. Artifact Resolution

Eine Resolution: resolution identity · requested artifact identity/alias · requested version constraint · resolved artifact identity · resolved version/revision · source · repository/distribution context · resolution time · resolver identity/component · provenance/integrity state · withdrawal/revocation state · compatibility state · known limitations · audit reference. **Mutable Aliases unterstützen Discovery; privilegiertes Deployment bindet aufgelöste Identität/Version/Revision** (Companion 3 §20). Keine Resolver-/Registry-Technologie.

## 6. Version and Revision

`artifact version ≠ artifact revision`; `same version label ≠ same revision`; `newer ≠ safer ≠ compatible ≠ authorised replacement`. Auflösung protokolliert version **und** revision.

## 7. Component Identity

Wie [Artifact Model §21]: `same name ≠ same component`; `same version label ≠ same revision`; `alias match ≠ confirmed identity`. Component-Identitätskonflikte bleiben sichtbar (Bezug THR-014).

## 8. Dependency Classes

`direct · transitive · build · runtime · optional · development/test · embedded · bundled · external service · unknown`.

## 9. Dependency Relationships

Je Relationship: source/target component · relationship class · version/compatibility scope · required/optional · source · evidence · validation state · freshness · known limitations. `transitive dependency ≠ visible by default`; Dependency Confusion (Bezug THR-023) bleibt sichtbar. Keine automatische Dependency Resolution.

## 10. Dependency Scope

Scope (build vs. runtime vs. dev/test) beeinflusst Impact/Vulnerability-Bewertung; eine Build-Dependency ist nicht automatisch Runtime-relevant und umgekehrt.

## 11. Compatibility

Klassen (not-assessed · expected-compatible · compatible-with-notes · validated-compatible · incompatible · deprecated-compatibility · unknown · conflicted). `same version ≠ same behaviour`; `unknown ≠ compatible`; Claims an Artifact-/Component-/Target-/CoreOps-Version gebunden (konsistent mit [Schema/Migration](SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md), [Deployment Blueprint](DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md)).

## 12. Target Constraints

Artifacts deklarieren Target-Class-/Version-Constraints; ein Deployment auf nicht-konforme Targets ist unzulässig (Bezug Deployment §12 Revalidation).

## 13. Distribution Context

Distribution: distribution identity · artifact identity/revision · source distribution · destination · repository context · consumer · workspace/environment · provenance/integrity/trust/revocation state · validity boundary · audit reference.
```text
available from approved repository ≠ artifact approved for deployment
repository trusted ≠ every contained artifact trusted
```

## 14. Repository Context

Repository-Trust ist scopegebunden und wird **nicht** von jedem enthaltenen Artifact geerbt. Keine Registry-/Repository-Technologie.

## 15. Validation

Validation (format/structure/provenance/integrity/dependency/scenario/target-scope) je Artifact/Component; scopegebunden, nicht über geprüfte Menge generalisiert.

## 16. Vulnerability and Advisory Context

Getrennt: vulnerability reference · advisory · affected-component/-version claim · artifact impact · deployment-context impact · exploitability · remediation/mitigation/exception · review trigger.
```text
vulnerability reference exists ≠ artifact affected
component affected ≠ deployment exploitable
no scanner finding ≠ no vulnerability
fixed version available ≠ compatible ≠ authorised
severity ≠ deployment-context risk automatically
```
Keine Scanner-/Advisory-Technologie.

## 17. Withdrawal

`withdrawal · support withdrawal · compatibility withdrawal · distribution removal`. `artifact withdrawn ≠ necessarily malicious`; `repository removal ≠ artifact revocation`.

## 18. Revocation

Revocation (Companion 3 §22): artifact identity/revision · authority · reason · scope · effective boundary · affected consumers/targets · distribution/deployment impact · existing-deployment assessment · offline-distribution impact · required follow-up · evidence · audit reference.
```text
artifact revoked ≠ every existing deployment automatically removed
revocation issued ≠ revocation delivered to every offline environment
```
Keine Revocation-Distribution-Technologie.

## 19. Replacement

`fixed/newer version available ≠ compatible ≠ authorised replacement`; ein Replacement durchläuft dieselbe Resolution/Trust/Compatibility-Bewertung. Kein sicheres Replacement kann fehlen (§20).

## 20. Existing Deployments

`artifact revoked ≠ automatic destructive removal`; `existing deployment ≠ permission for new deployment`; `revocation received → new deployment blocked unless explicit exception applies`. Bestehende Deployments benötigen gesonderte Impact-/Mitigation-/Recovery-/Exception-Bewertung (Companion 3 §23, Bezug THR-024).

## 21. Offline Distribution

Offline-Artifact-Distribution folgt Companion 3 §26 (target-environment binding, provenance, integrity, import quarantine, explicit import). Keine konfliktfreie automatische Reconciliation.

## 22. Delayed Revocation

`revocation not yet delivered · artifact already deployed/staged/running · rollback artifact also revoked · offline target unreachable · replacement incompatible/absent`. Delayed Revocation bleibt sichtbar; `revocation received → new deployment blocked unless explicit exception`.

## 23. Audit and Evidence

Resolution-/Dependency-/Compatibility-/Vulnerability-/Withdrawal-/Revocation-Aktivitäten werden auditiert (Companion 3 §29). `artifact evidence ≠ trust decision ≠ execution authorization`.

## 24. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Mutable aliases must not serve as final privileged-deployment binding; resolution binds identity, version and revision.
2. Same version label may resolve to different revisions; resolution records both.
3. Repository trust is not inherited by every contained artifact.
4. Vulnerability reference does not automatically imply artifact or deployment exploitability.
5. Missing scanner finding or SBOM component does not prove absence.
6. Withdrawal, revocation and repository removal remain separate.
7. Revocation blocks new deployment unless an explicit governed exception applies; existing deployment is not permission for new deployment.
8. Revocation and reinstatement must preserve historical evidence.
9. Offline artifact distribution requires provenance, integrity, target binding and explicit import governance.

## 25. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-013, THR-014, THR-021, THR-022, THR-023, THR-024, THR-026. Keine Duplikation, kein Parallelregister.

## 26. Technology Boundary

Nicht ausgewählt/implementiert: Artifact Registry/Repository, Dependency Resolver, Scanner, SBOM-Format, Revocation-Distribution, Versionsnotation.

## 27. Open Questions

- Dependency-Confusion-Schutz je Ecosystem (spätere ADR).
- Revocation-Distribution/Delayed-Revocation-Mechanismus (mit CO-WP-023).
- Deployment-Context-Exploitability-Bewertung im Detail.

## 28. Next Decision

Companion 3 trägt die Trust-/Quarantine-/Revocation-Policy. Restricted/Air-Gapped Operation (CO-WP-023) und Self-Protection/Recovery (CO-WP-026) konkretisieren Offline-/Revocation-Aspekte. Resolver-/Scanner-Wahl bleibt einer späteren ADR-Runde vorbehalten.
