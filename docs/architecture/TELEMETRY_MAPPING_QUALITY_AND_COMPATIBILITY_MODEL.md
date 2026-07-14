# CoreOps – Telemetry Mapping, Quality and Compatibility Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation telemetry mapping and quality model
> Implementation Status: Not implemented
> Mapping Engine: Not selected
> Unit Library: Not selected
> Compatibility Validation: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-019 (docs-only / telemetry, observability and normalization foundation)

## 1. Status

Technologieunabhängiges Modell für **Source-Schemas, Canonical Fields, Normalization Profiles, Mapping/Transformation, Units, Quality/Confidence und Compatibility**. Companion zu [TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md) und [TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md](../security/TELEMETRY_TRUST_PRIVACY_AND_DISCLOSURE_POLICY.md). Keine Mapping-Engine, keine Unit-Library, kein Schemaformat.

## 2. Purpose

Mapping ist der Ort, an dem Semantik still verloren geht: `mapping succeeded ≠ semantic correctness verified`, `same unit ≠ same scale`, `normalization success ≠ semantic validation`. Dieses Modell trennt Profile-Identität von Version, Unit von Scale/Precision und Quality von Confidence und bindet jede Abbildung an Provenance und Lossiness.

## 3. Scope

Source Schemas · Canonical Fields · Normalization Profiles/Identity/Versioning · Mapping Classes · Transformation History · Units/Scale/Precision · Missing/Unknown Values · Lossiness · Quality/Confidence States · Validation · Compatibility · Schema/Profile Changes · Deprecation · Sampling/Aggregation Impact · Offline Mapping · Audit/Evidence.

## 4. Non-Goals

- Keine Mapping-Engine, keine Unit-Library, kein Schemaformat/Serialisierung.
- Keine Compatibility-Validation-Runtime, keine Behauptung durchgeführter Validierung.

## 5. Source Schemas

Ein Source Schema beschreibt die Struktur/Semantik der von einer Quelle gelieferten Telemetrie; es ist quellenspezifisch, versioniert und **nicht** autoritativ für CoreOps. `source field ≠ canonical field`.

## 6. Canonical Fields

Wie [Signal Model §18]: stabile Field-Identität, Semantik, Unit-Klasse, Scale, Precision, Value-Classification. `same label ≠ same semantics`; `mapped field ≠ authoritative field`.

## 7. Normalization Profiles

Ein Profile bildet ein Source-Schema auf Canonical Fields ab. Felder: stable profile identity · owner · source integration/class · source schema/version · target canonical field set · profile version · applicable resource classes · mapping rules · unit conversion rules · precision rules · lossiness classification · default/missing-value behavior · quality impact · compatibility state · validation state · known limitations · evidence reference · deprecation state.
```text
profile exists    ≠ mapping validated
mapping succeeded ≠ semantic correctness verified
new profile version ≠ automatically compatible
fallback mapping  ≠ safe mapping
```

## 8. Profile Identity

Stabile Profile-Identität (≠ display name/file). Retired Profile-IDs nicht still wiederverwendet; historische Profile/Evidenz bleiben erhalten.

## 9. Profile Versioning

Profile-Version getrennt von Source-Schema-Version, Canonical-Field-Version, Contract-/Produkt-Version (konsistent mit [Schema/Migration](SCHEMA_VERSIONING_AND_MIGRATION_MODEL.md)). Neuere Version ≠ automatisch kompatibel.

## 10. Mapping Classes

`direct mapping · renamed mapping · unit-converted mapping · scaled mapping · derived mapping · filtered mapping · aggregated mapping · default-substituted mapping · lossy mapping · unknown/unmapped`. Je Klasse: Compatibility-/Quality-/Lossiness-Auswirkung.

## 11. Transformation History

Erfasst: source field · source value classification · target canonical field · mapping rule reference · conversion · normalisation · filtering · redaction · aggregation · derivation · precision change · unit change · default substitution · quality impact · lossiness · producer/processor · time boundary · provenance reference. `transformation history ≠ legal chain of custody`; Raw Provenance nicht still entfernt.

## 12. Units

Unit Class · Source Unit · Canonical Unit · Conversion Rule · Unknown Unit · Conflicting Unit · Unit Version/Definition Context. `unknown unit ≠ canonical unit`; `conflicting unit ≠ safe automatic conversion` (fail-safe).

## 13. Scale

Scale getrennt von Unit; `same unit ≠ same scale`. Skalierungsfehler sind Quality-relevant.

## 14. Precision

Precision/Rounding getrennt; `unit converted ≠ precision preserved`; `rounded value ≠ source value`. Precision-Verlust wird als Lossiness erfasst.

## 15. Missing and Unknown Values

`zero ≠ missing ≠ unknown`; `missing sample ≠ zero value`. Default-Substitution ist explizit und quality-/provenance-erfasst; keine stille Null-Substitution.

## 16. Lossiness

Lossiness-Klassifikation (lossless · rounding-loss · unit-loss · aggregation-loss · filtering-loss · derivation-loss · unknown). Lossy Mapping bleibt sichtbar; `aggregated ≠ complete population`.

## 17. Quality States

`not-assessed · raw-unvalidated · normalized-unvalidated · validated-with-notes · validated · degraded · stale · invalid · conflicted · unknown`.
```text
normalization success ≠ semantic validation
validated ≠ fresh
fresh     ≠ complete
```

## 18. Confidence States

`not-assessed · low · medium · high · source-reported · derived · conflicted · unknown`.
```text
high confidence ≠ authoritative
multiple samples ≠ independent sources
```
Keine numerische Confidence-Skala erzwungen.

## 19. Validation

Validierung (format/schema/unit/semantic/referential) ist eigenständig; `normalized ≠ validated`; scopegebunden, nicht über geprüfte Signale generalisiert.

## 20. Compatibility

Compatibility-Klassen (not-assessed · source-compatible · canonical-compatible · behaviourally-compatible · incompatible · deprecated · unknown · conflicted). `unknown ≠ compatible`; Claims an Source-Schema-/Profile-/Field-Version gebunden.

## 21. Schema and Profile Changes

Change Classes (additive field · removed field · semantic change · unit change · scale change · mapping change · default-behaviour change · lossiness change · identifier change). Eine formal additive Änderung kann semantisch breaking sein; benötigt Compatibility-/Quality-/Deprecation-Bewertung.

## 22. Deprecation

Deprecated Profile/Field erhält Migrationshinweise; `deprecated ≠ removed`; historische Identität/Evidenz erhalten; retired IDs nicht wiederverwendet.

## 23. Sampling Impact

Sampling (Signal Model §24) beeinflusst Quality/Completeness der gemappten/aggregierten Werte; `sampled ≠ complete population`; Sampling-State bleibt in Mapping-/Evidence-Aussagen sichtbar.

## 24. Aggregation Impact

Aggregation (Signal Model §25) beeinflusst Lossiness/Authority; `aggregate value ≠ authoritative individual value`; Aggregations-Provenance erhalten.

## 25. Offline Mapping

Offline erzeugte/angewandte Mappings benötigen Provenance/Integrität/Target-Binding/explizite Aktivierung (Companion 3 §23). Keine konfliktfreie automatische Reconciliation behauptet.

## 26. Audit and Evidence

Mapping-/Profile-/Validation-Aktivitäten werden auditiert (konsistent mit [Audit Policy](../security/AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md)); `evidence available ≠ valid ≠ sufficient`.

## 27. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. A mapped field is not an authoritative field.
2. Mapping success is not semantic correctness verification.
3. Normalization must preserve source provenance and transformation history.
4. Unit, scale and precision remain separate; unknown/conflicting units block unsafe automatic conversion.
5. Zero, missing and unknown remain distinct; no silent default substitution.
6. Lossiness must remain visible.
7. A new profile/schema version is not automatically compatible; formally additive changes may be breaking.
8. High confidence is not authoritative truth; multiple samples are not independent sources.
9. Retired profile/field identities must not be reused silently.

## 28. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-011, THR-012, THR-013, THR-014, THR-024. Keine Duplikation, kein Parallelregister.

## 29. Technology Boundary

Nicht ausgewählt/implementiert: Mapping-Engine, Unit-Library, Schemaformat, Compatibility-Validation-Runtime, Storage, Runtime-Code.

## 30. Open Questions

- Standardisierte Semantik-Validierung je Field.
- Unit-/Scale-Konventionen im Detail.
- Independent-Source-Kriterien für Confidence.

## 31. Next Decision

Companion 3 trägt Trust/Privacy/Disclosure. Topology Graph (CO-WP-020) und Test Strategy (CO-WP-028) referenzieren Quality/Validation. Mapping-Engine-/Unit-Library-Wahl bleibt einer späteren ADR-Runde vorbehalten.
