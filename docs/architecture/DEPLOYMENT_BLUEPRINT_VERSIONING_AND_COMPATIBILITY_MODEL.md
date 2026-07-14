# CoreOps – Deployment Blueprint Versioning and Compatibility Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation deployment-blueprint versioning model
> Implementation Status: Not implemented
> Blueprint Format: Not selected
> Schema Language: Not selected
> Template Engine: Not selected
> Compatibility Validation: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-021 (docs-only / deployment architecture, targeting and execution-governance foundation)

## 1. Status

Technologieunabhängiges Modell für **Blueprint-Identität, -Versionierung, Inputs/Parameters, Environment Overlays, Effective Blueprint, Artifact-/Dependency-Bindung und Compatibility**. Companion zu [DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md) und [DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md](../security/DEPLOYMENT_TARGETING_EXECUTION_AND_RECOVERY_POLICY.md). Kein Blueprint-Format, keine Schema-Sprache, keine Template Engine.

## 2. Purpose

Ein Blueprint ist eine wiederverwendbare Deployment-Definition; sein Effective-Zustand entsteht erst durch Parameter/Overlays und darf Security nicht still schwächen. Dieses Modell trennt Blueprint-Identität von Version, Base von Effective und bindet jede Ableitung an Provenance — damit `default value ≠ universally safe value`, `overlay applied ≠ permission to weaken security` und `parameter accepted ≠ semantically validated`.

## 3. Scope

Blueprint Concepts/Identity/Lifecycle · Version Dimensions · Inputs/Parameters/Defaults/Secret References · Environment Overlays · Effective Blueprint · Target Constraints · Artifact/Dependency Binding · Compatibility · Validation · Known Limitations · Deprecation/Retirement · Migration Guidance · Domain-Pack Relationship · Offline · Audit.

## 4. Non-Goals

- Kein Blueprint-/Manifest-Format, keine Schema-Sprache (YAML/JSON/DSL), keine Template Engine.
- Keine Merge-/Templating-Technologie, keine Compatibility-Validation-Runtime.
- Keine Behauptung durchgeführter Validierung; keine Raw Secrets.

## 5. Blueprint Concepts

Ein Blueprint definiert eine Deployment-Wirkung parametrisiert; er ist keine Runtime und kein Deployment Plan (`blueprint ≠ deployment plan`, Companion 1).

## 6. Blueprint Identity

Ein Blueprint benötigt: stable blueprint identity · canonical name · owning module · owner · maintainer · Domain-Pack reference (falls anwendbar) · blueprint version · lifecycle/support/validation/compatibility state · input definitions · parameter classes · artifact/dependency references · target-class constraints · policy requirements · known limitations · deprecation state · audit reference.
```text
display name    ≠ blueprint identity
file name       ≠ blueprint identity
repository path ≠ blueprint identity
blueprint version ≠ CoreOps product version
```
**Blueprint-IDs werden nicht still wiederverwendet.**

## 7. Blueprint Lifecycle

`proposed · draft · preview · active · restricted · deprecated · security-maintenance-only · retired · archived` (konzeptionell).
```text
active     ≠ implemented ≠ supported ≠ validated
deprecated ≠ unavailable
retired    ≠ historical evidence deleted
```

## 8. Version Dimensions

Getrennt: blueprint version · effective-configuration version · artifact version · dependency version · schema/format version · Domain-Pack version · CoreOps product version. Keine Dimension ersetzt still eine andere (konsistent mit [Schema/Migration](SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md), [API Versioning](API_VERSIONING_COMPATIBILITY_AND_DEPRECATION_MODEL.md)). Keine Notation erzwungen (DEC-O-02).

## 9. Inputs and Parameters

Zu unterscheiden: `required · optional · defaulted · environment-derived · operator-supplied · policy-derived · secret reference · artifact reference · target-derived · computed · unknown input`. Je Input: stable input identity · semantic definition · value classification · source · owner · required/optional · default behavior · validation requirement · scope · sensitivity · compatibility impact · audit expectation.
```text
default value        ≠ universally safe value
parameter accepted   ≠ semantically validated
target-derived input ≠ target authorization
```

## 10. Defaults

Defaults sind explizit und provenance-erfasst; ein Default ist **kein** universell sicherer Wert. Fehlender Pflicht-Input → fail-closed (Companion 1 §15), keine stille Default-Substitution sicherheitsrelevanter Werte.

## 11. Secret References

`secret reference ≠ raw secret value`. Blueprints/Parameter enthalten **keine** Raw Secrets/Credential-Inhalte; nur Referenzen (Roh-Secret-Custody offen, CO-WP-024). Bezug THR-019, THR-020.

## 12. Environment Overlays

Overlays berücksichtigen Environment/Workspace/Site/Deployment Profile/Domain Pack/Target Class.
```text
overlay applied     ≠ base blueprint modified
overlay precedence  ≠ authority precedence
environment overlay ≠ permission to weaken security requirements
```
Jede effektive Konfiguration hat nachvollziehbare Herkunft/Transformationshistorie. Keine Merge-/Templating-Technologie.

## 13. Effective Blueprint

Der Effective Blueprint ist Base + aufgelöste Parameter + Overlays mit vollständiger Provenance. `effective configuration ≠ validated configuration`; Sicherheitsanforderungen der Base bleiben Untergrenze (§12).

## 14. Target Constraints

Blueprints deklarieren Target-Class-Constraints; ein Deployment auf nicht-konforme Target-Klassen ist unzulässig. Target Constraints sind Teil der Compatibility (§17).

## 15. Artifact Binding

Ein Artifact (Companion 1 §13): identity · version · class · source · owner · producer · provenance/integrity/validation status · compatibility state · target constraints · dependency references · known limitations · withdrawal/revocation state · audit reference. `artifact available ≠ trusted ≠ compatible`; `newer artifact ≠ authorised replacement`. Keine Registry-/Signing-/Hash-Technologie (CO-WP-022).

## 16. Dependency Binding

Dependencies (Companion 1 §14): identity · type · version/compatibility scope · required/optional · target relevance · availability · trust/validation state · failure behavior · offline availability · audit. `optional ≠ mandatory Core dependency`; keine automatische Dependency Resolution.

## 17. Compatibility

Compatibility-Klassen (not-assessed · expected-compatible · compatible-with-notes · validated-compatible · incompatible · deprecated-compatibility · unknown · conflicted). `same version ≠ same behaviour`; `newer ≠ automatically compatible`; `unknown ≠ compatible`. Claims an Blueprint-/Artifact-/Target-/CoreOps-Version gebunden.

## 18. Validation

Validation (format/schema/parameter/semantic/target-constraint/artifact-compatibility) ist eigenständig; `parameter accepted ≠ validated`; `effective ≠ validated`; scopegebunden.

## 19. Known Limitations

Blueprints mit Support ≥ dokumentieren bekannte Limitationen (analog Domain Pack Support); fehlende Angabe ist kein Vollständigkeitsnachweis.

## 20. Deprecation

Deprecated Blueprints erhalten Migrationshinweise; `deprecated ≠ unavailable`; historische Identität/Evidenz erhalten.

## 21. Retirement

`retired` erhält keine neuen Deployments/Claims; retired Blueprint-IDs nicht wiederverwendet; historische Evidenz erhalten.

## 22. Migration Guidance

Blueprint-/Parameter-/Artifact-Änderungen liefern Migrationshinweise (Version/Verhalten/Compatibility); Guidance ist keine Garantie erfolgter Migration.

## 23. Domain-Pack Relationship

Blueprints können von Domain Packs (CO-WP-015) referenziert/bereitgestellt werden; Pack-Status validiert einen Blueprint nicht automatisch (`pack activation ≠ blueprint validated`).

## 24. Offline Blueprints

Offline verfügbare Blueprints/Parameter/Artifacts benötigen Provenance/Integrität/Target-Binding/explizite Aktivierung (Companion 3 §25). Keine konfliktfreie automatische Synchronisation.

## 25. Audit and Evidence

Erfasst: blueprint creation/version · parameter/overlay resolution · effective configuration provenance · artifact/dependency binding · compatibility classification · validation · deprecation/retirement. `evidence available ≠ valid ≠ sufficient`.

## 26. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Blueprint identity is stable and separate from file name, repository path, schema format and product version.
2. Blueprint version, artifact version and dependency version remain separate.
3. Defaults are not universally safe values; parameter acceptance is not semantic validation.
4. Environment overlays must not weaken base security requirements silently.
5. Effective configuration provenance and transformation history must be preserved.
6. Secret references are not raw secret values; blueprints hold no raw secrets.
7. Artifact availability does not imply trust, integrity or compatibility.
8. A new blueprint/artifact version is not automatically compatible or an authorised replacement.
9. Retired blueprint identities must not be reused; historical evidence must not be deleted.

## 27. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-006, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024. Keine Duplikation, kein Parallelregister.

## 28. Technology Boundary

Nicht ausgewählt/implementiert: Blueprint-/Manifest-Format, Schema-Sprache, Template/Merge Engine, Compatibility-Validation-Runtime, Artifact Registry, Secret-Technologie.

## 29. Open Questions

- Blueprint-Format-/Schema-Wahl (spätere ADR).
- Overlay-Präzedenzregeln im Detail.
- Roh-Secret-Referenzmechanismus (CO-WP-024).

## 30. Next Decision

Companion 3 trägt Targeting/Execution/Recovery-Policy. Artifact Trust (CO-WP-022) und Secrets/Key Custody (CO-WP-024) konkretisieren Artifact-/Secret-Aspekte. Format-/Engine-Wahl bleibt einer späteren ADR-Runde vorbehalten.
