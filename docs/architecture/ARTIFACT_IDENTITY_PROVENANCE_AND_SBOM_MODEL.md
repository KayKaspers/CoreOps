# CoreOps – Artifact Identity, Provenance and SBOM Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation artifact identity, provenance and SBOM model
> Implementation Status: Not implemented
> Artifact Format: Not selected
> SBOM Format: Not selected
> Integrity Mechanism: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-022 (docs-only / artifact identity, software-supply-chain and trust-governance foundation)

## 1. Status

Technologieunabhängiges Modell für **Artifact-Identität, Version/Revision, Provenance, Integrity, Validation, Trust und SBOM/Components**. Companion zu [ARTIFACT_DEPENDENCY_COMPATIBILITY_AND_DISTRIBUTION_MODEL.md](ARTIFACT_DEPENDENCY_COMPATIBILITY_AND_DISTRIBUTION_MODEL.md) und [ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md](../security/ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md). Es wählt **keine** Artifact Registry, kein SBOM-Format, kein Signing/Hash-Verfahren, keinen Scanner, keine Build-Plattform.

## 2. Purpose

Ein Artifact ist Deployment-Input mit Supply-Chain-Risiko: verfügbar heißt nicht vertrauenswürdig, integer heißt nicht sicher, ein SBOM heißt nicht vollständig. Dieses Modell trennt Identität von Alias, Version von Revision, Provenance von Integrity von Validation von Trust — damit `artifact available ≠ trusted`, `integrity checked ≠ safe`, `SBOM available ≠ complete ≠ accurate` und `no known vulnerability ≠ secure`. Es baut auf [Deployment Artifact Binding](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md) (CO-WP-021) und [Evidence Model](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md) (CO-WP-018) auf (kein Parallelmodell).

## 3. Scope

Artifact-Klassen/-Identität · Version/Revision/Instance/Alias/Resolution · Lifecycle · Rollen · Source/Producer/Builder/Build Context · Provenance · Integrity · Validation · Trust Assessment · SBOM/Components/Dependencies · Vulnerability/Advisory · Deployment Binding · Offline · Audit. Dependency/Distribution-Detail (Companion 2), Trust/Quarantine/Revocation-Policy (Companion 3).

## 4. Non-Goals

- Keine Artifact Registry/Package Repository, kein Package-/Archiv-/SBOM-Format.
- Kein Signing-/Hash-Verfahren, kein Trust Anchor/PKI/Transparency Log.
- Kein Vulnerability-/Malware-Scanner, kein Build-System/CI, kein Dependency Resolver, keine Versionsnotation.
- Kein Runtime-Code, kein Artifact-Download/Scan; keine Behauptung implementierter/vertrauenswürdiger/sicherer Artifacts.

## 5. Concepts

Begriffe (mindestens): artifact · artifact class · artifact identity · version · revision · instance · alias · resolution · component · dependency · SBOM · provenance · integrity · validation · trust assessment · support state · compatibility · vulnerability · advisory · exploitability · quarantine · withdrawal · revocation · reinstatement · distribution · deployment binding.

**Grundregeln:**
```text
artifact available ≠ artifact trusted
artifact identity  ≠ artifact alias ≠ file name ≠ repository path
artifact version   ≠ artifact revision
provenance available ≠ provenance validated
integrity checked  ≠ artifact safe
SBOM available     ≠ SBOM complete ≠ SBOM accurate
vulnerability reported ≠ artifact exploitable in every context
no known vulnerability ≠ artifact secure
artifact revoked   ≠ historical evidence deleted
```

## 6. Artifact Classes

Je Klasse: Purpose · Owner · Producer · Typical source · Versioning/Provenance/Integrity/SBOM/Validation expectation · Target constraints · Disclosure sensitivity · Offline relevance · Threat references. **Keine konkreten Paketformate.**

| Artifact-Klasse | Provenance-Erwartung | SBOM-Erwartung | Threat refs |
|---|---|---|---|
| application / service artifact | hoch | hoch | THR-021, THR-022 |
| agent / adapter artifact | hoch | hoch | THR-008, THR-010 |
| integration artifact | mittel-hoch | mittel | THR-010, THR-023 |
| Domain-Pack artifact | mittel | mittel | THR-023 |
| deployment blueprint artifact | mittel | niedrig | THR-022 |
| configuration / policy artifact | mittel | niedrig | THR-006 |
| schema / migration artifact | mittel | niedrig | THR-014 |
| container/image-like artifact | hoch | hoch | THR-022, THR-023 |
| package/archive-like artifact | hoch | hoch | THR-021, THR-023 |
| script artifact | hoch | niedrig | THR-006 |
| firmware-like artifact | sehr hoch | mittel | THR-021 |
| documentation artifact | niedrig | n/a | THR-018 |
| evidence package | hoch | n/a | THR-016, THR-018 |
| offline distribution package | hoch | mittel | THR-024 |
| unknown artifact | quarantäne | quarantäne | THR-014 |

## 7. Artifact Identity

26 Felder: stable artifact identity · artifact class · canonical name · owner · maintainer · producer · source · version · revision · lifecycle/support/compatibility/provenance/integrity/validation/trust state · SBOM reference · component/dependency references · target constraints · distribution/withdrawal/revocation state · known limitations · audit reference.
```text
display name ≠ artifact identity
file name    ≠ artifact identity
download URL or repository path ≠ artifact identity
artifact alias ≠ resolved artifact identity
```
**Artifact-IDs werden nicht still wiederverwendet.**

## 8. Version, Revision and Instance

Getrennt: artifact version · artifact revision · build instance · distribution instance · repository instance · deployment instance · alias · resolution result.
```text
same version label ≠ same artifact revision
same artifact bytes ≠ same provenance automatically
same alias         ≠ same resolved artifact over time
newer version      ≠ safer ≠ compatible ≠ authorised replacement
```
Keine Versionsnotation.

## 9. Aliases and Resolution

Detail Companion 2 §5. `alias match ≠ identity match`. **Mutable Aliases dürfen Discovery unterstützen, aber privilegiertes Deployment bindet die aufgelöste Artifact-Identität, -Version und -Revision** (§24, Companion 3 §20).

## 10. Lifecycle

Konzeptionelle Statuswerte:
```text
proposed → received → quarantine-pending → quarantined → assessment-pending → assessed →
validation-pending → validated → validation-failed → active → restricted → withdrawn → revoked →
reinstatement-pending → reinstated → deprecated → retired → archived → invalidated
```
```text
received  ≠ trusted
assessed  ≠ validated
validated ≠ supported
active    ≠ authorised for every target
withdrawn ≠ revoked
reinstated ≠ prior trust automatically restored
```

## 11. Roles

Getrennt: artifact owner · maintainer · source · producer · builder · build-environment owner · distributor · repository operator · validator · trust decision owner · revocation authority · deployment consumer · evidence owner.
```text
repository operator ≠ artifact owner
builder             ≠ source-code owner
distributor         ≠ trust authority
maintainer          ≠ revocation authority automatically
```

## 12. Source and Producer

Source (Ursprungscode/-quelle) und Producer (erzeugende Instanz) getrennt; `producer stated ≠ producer verified`. Source Trust ist zeitgebunden (konsistent mit [Domain Pack Trust](../security/DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md)).

## 13. Builder and Build Context

Builder-Identität und Build Context (Inputs, Dependencies, Environment, Boundary) sind Teil der Provenance (§14); kompromittierter Builder (Bezug THR-023) invalidiert Trust. `build metadata present ≠ provenance complete`.

## 14. Provenance

16 Felder: artifact identity/revision · source identity/version · producer · builder identity · build context/inputs · dependency inputs · build/assembly boundary · transformations · distribution steps · repository contexts · SBOM reference · validation activity · owner · known gaps · audit reference.
```text
producer stated       ≠ producer verified
build metadata present ≠ provenance complete
provenance valid      ≠ artifact free of vulnerabilities
handling history      ≠ legal admissibility
```
Konsistent mit [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md); keine Attestation-Technologie.

## 15. Integrity

`not-assessed · integrity-reference-present · integrity-check-pending · integrity-checked · integrity-failed · integrity-unknown · integrity-conflicted`.
```text
integrity metadata present ≠ integrity verified
integrity verified ≠ provenance verified ≠ artifact trusted ≠ artifact safe
```
Keine Hash-/Signing-/Verification-Technologie.

## 16. Validation

`not-assessed · document-reviewed · structure-validated · provenance-validated · integrity-validated · dependency-reviewed · scenario-validated · target-scope-validated · invalid · stale · superseded · conflicted`.

## 17. Trust Assessment

`not-assessed · untrusted · restricted · conditionally-trusted · trusted-for-bounded-scope · trust-review-required · suspended · revoked · conflicted · unknown`.
```text
validated  ≠ universally trusted
trusted    ≠ compatible with every target
supported  ≠ trusted
trusted-for-bounded-scope requires explicit scope, versions, evidence, owner and review trigger
```

## 18. SBOM Concepts

`SBOM capability · requirement · document/artifact · identity · version · generation context · component · component identity · component version · dependency relationship · completeness · accuracy · freshness · validation`.
```text
SBOM exists      ≠ every component known
component listed ≠ component actually present
component absent from SBOM ≠ component absent from artifact
SBOM validated   ≠ artifact trusted
multiple SBOMs   ≠ independent evidence sources
```
Keine SBOM-Spezifikation.

## 19. SBOM References

Ein SBOM Reference (20 Felder): stable SBOM identity · subject artifact identity/revision · producer · generation context/time · source inputs · component population scope · dependency scope · format reference · schema version reference · freshness/completeness/accuracy/validation/provenance/integrity state · known exclusions · supersession reference · audit reference. **Ein Reference ersetzt nicht das SBOM-Artefakt.**

## 20. Components

Ein Component: stable component identity · component class · canonical name · source/supplier · version · revision (falls bekannt) · relationship to subject artifact · scope · license-information reference (falls verfügbar) · provenance reference · integrity/validation state · vulnerability references · known aliases · identity-conflict state.

## 21. Component Identity

```text
same component name ≠ same component
same version label  ≠ same revision
alias match          ≠ confirmed component identity
```
Keine Package-URL-/CPE-/andere Identifier-Technologie.

## 22. Dependency Relationships

Zusammenfassung (Detail Companion 2 §9): direct/transitive/build/runtime/optional/dev-test dependency, embedded/bundled component, external service dependency, unknown. Keine automatische Dependency Resolution.

## 23. Vulnerability and Advisory Context

Getrennt (Detail Companion 2 §16): vulnerability reference · advisory · affected-component claim · affected-version claim · artifact impact · deployment-context impact · exploitability · remediation availability · mitigation/exception state · review trigger.
```text
vulnerability reference exists ≠ artifact affected
component affected ≠ deployment exploitable
no scanner finding ≠ no vulnerability
fixed version available ≠ replacement compatible ≠ authorised
severity ≠ deployment-context risk automatically
```

## 24. Deployment Binding

Ein Deployment Plan bindet konkrete Resolution-Ergebnisse (Companion 3 §20): artifact identity · resolved version/revision · provenance/integrity/validation state · trust decision · SBOM reference · vulnerability-assessment reference · compatibility state · withdrawal/revocation state. **`deployment plan approved ≠ any later artifact resolution approved`**; Pre-Execution/Wave-Recheck erforderlich.

## 25. Offline Artifacts

Offline-Artifact-Pakete folgen Companion 3 §26 (target-environment binding, provenance, integrity, bounded trust, import quarantine, explicit import, delayed revocation). Nicht behauptet: implementiert, beliebige Air-Gap-Stufen, konfliktfreie automatische Reconciliation, Klassifiziertnetz-Eignung, konkrete Signing-/Hash-/Trust-Anchor-Technologie (Bezug THR-024).

## 26. Audit and Evidence

Erfasst wie Companion 3 §29. Trennung `artifact evidence ≠ trust decision ≠ execution authorization`. `evidence available ≠ valid ≠ sufficient` (CO-WP-018).

## 27. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Artifact availability must not imply trust, integrity, compatibility or deployment authorization.
2. Artifact identity, version, revision, instance and alias remain separate.
3. Mutable aliases must not serve as final privileged-deployment binding.
4. Source, producer, builder, distributor, repository operator and trust decision owner remain separate roles.
5. Provenance, integrity, validation, trust, support and compatibility remain separate dimensions.
6. SBOM existence must not imply completeness, accuracy or artifact trust.
7. Missing component or vulnerability information must not prove absence.
8. Trust decisions remain use-, target-, scope- and time-bound.
9. Artifact revocation must preserve historical evidence.

## 28. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-006, THR-008, THR-010, THR-013, THR-014, THR-016, THR-018, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024, THR-026.

## 29. Technology Boundary

Nicht ausgewählt/implementiert: Artifact Registry/Package Repository, Package-/Archiv-/SBOM-Format, Hash/Signing/Trust-Anchor/Transparency Log, Vulnerability-/Malware-Scanner, Build-System/CI, Dependency Resolver, Versionsnotation, Runtime-Code.

## 30. Compatibility

Konsistent mit [Deployment](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)/[Blueprint](DEPLOYMENT_BLUEPRINT_VERSIONING_AND_COMPATIBILITY_MODEL.md), [Integration Trust](../security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [Domain Pack Trust](../security/DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md), [Evidence Model](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [Field Provenance](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md), [Policy/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Konkretisiert DEC-P-02, DEC-G-05, die Sovereignty-/Dependency-Linie und die adoptierten NDF-Supply-Chain-Kandidaten.

## 31. Open Questions

- Artifact-/SBOM-Format-Wahl (spätere ADR).
- Integrity-/Attestation-Mechanismus (spätere ADR).
- Component-Identifier-Konventionen.

## 32. Next Decision

Companion 2 (Dependency/Distribution) und Companion 3 (Trust/Quarantine/Revocation-Policy). Restricted/Air-Gapped Operation (CO-WP-023) und Secrets/Key Custody (CO-WP-024) konkretisieren Offline-/Signing-Aspekte. Format-/Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
