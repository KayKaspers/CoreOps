# CoreOps – Domain Pack Support and Compatibility Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation support and compatibility model
> Implementation Status: Not implemented
> Support SLA: None defined
> Compatibility Validation: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-015 (docs-only / modular product-governance and compatibility foundation)

## 1. Status

Technologieunabhängiges Modell für **Support Levels, Maintenance Responsibility, Implementation/Validation/Evidence-Status und Compatibility Claims** von Domain Packs. Companion zu [DOMAIN_PACK_GOVERNANCE_MODEL.md](DOMAIN_PACK_GOVERNANCE_MODEL.md) und [DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md](../security/DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md). **Kein SLA definiert; keine Compatibility Validation durchgeführt.**

## 2. Purpose

Support- und Kompatibilitätsaussagen sind der häufigste Ort für Overclaim. Dieses Modell trennt Support Level, Maintenance, Implementation, Validation, Evidence und Compatibility strikt und bindet jeden Compatibility Claim an Version, Ziel-System, Profil und Evidenz — damit `supported ≠ validated`, `expected-compatible ≠ validated-compatible` und `support level ≠ SLA`.

## 3. Scope

Support-Konzepte/-Levels · Maintenance Responsibility · Implementation/Validation/Evidence Status · Compatibility Status/Dimensions · Claim Requirements · Known Limitations · Dependencies · Support Changes · Deprecation/Retirement · Community/External/Vendor-Support-Grenzen · Offline-Support-Grenze · Audit/Evidence.

## 4. Non-Goals

- Kein SLA, keine Reaktionszeit-Zusage.
- Keine Marketplace-/Packaging-/Update-/Dependency-Resolution-Technologie.
- Keine Behauptung durchgeführter Validierung/Zertifizierung.
- Keine Vendor-Support-Behauptung.

## 5. Support Concepts

**Support Level** (Zusage-Umfang), **Maintenance Responsibility** (wer pflegt), **Implementation Status** (Code vorhanden?), **Validation Status** (nachgewiesen?), **Evidence Status** (Nachweis aktuell?), **Compatibility Status** (mit was?) sind unabhängige Dimensionen. `evidence capability ≠ evidence available ≠ requirement satisfied` (aus CO-WP-004E).

## 6. Support Levels

- **SUP-0 — unsupported:** keine Support-Zusage; Nutzung auf eigenes Risiko; keine regelmäßige Pflege behauptet.
- **SUP-1 — community-best-effort:** Community-Unterstützung möglich; keine Reaktions-/Lösungszusage; kein SLA.
- **SUP-2 — project-maintained:** aktiver Maintainer und definierter Scope; Issues/Security-Meldungen nach Projektprozess; **keine** automatische Kompatibilitäts-/Validierungsgarantie; kein SLA, sofern nicht später ausdrücklich beschlossen.
- **SUP-3 — validated-supported-scope:** nur zulässig, wenn dokumentiert: supported pack version · CoreOps version range · Integration Contract version · target-system versions · deployment profiles · operating modes · validation evidence · known limitations · maintenance owner · security-review state · support exclusions. **SUP-3 ist kein Vendor-Endorsement und keine Zertifizierung.**
- **SUP-D — deprecated-limited-maintenance:** nur begrenzte Pflege (typischerweise kritische Fehler/Security-Fixes); dokumentiertes Enddatum/Review-Gate erforderlich; Migrationshinweis erforderlich.

```text
support level             ≠ SLA
project-maintained        ≠ validated
validated-supported-scope ≠ universal support
external vendor support   ≠ CoreOps support
```

## 7. Maintenance Responsibility

`unassigned · community-maintained · project-maintained · external-maintainer · joint-maintenance · maintenance-suspended · retired`. Beschreibt Verantwortung für Änderungen/Reviews — **keine** Aussage über Support Level, SLA, Validation, Vendor Support oder Security Certification. Getrennt vom Support Level (§6).

## 8. Implementation Status

`not-implemented · planned · partial · implemented · implementation-deprecated · removed`. Ein Pack kann als Governance-Modell **active** sein, ohne dass eine Implementierung existiert. **Keine Implementierungsbehauptung ohne überprüfbare Artefakte.**

## 9. Validation Status

`not-assessed · document-reviewed · contract-validated · integration-tested · scenario-validated · operationally-observed · invalid · superseded`.
```text
validation capability ≠ validation performed
test passed once      ≠ supported across all versions
community report      ≠ project validation
```

## 10. Evidence Status

`no-evidence · evidence-planned · partial-evidence · evidence-available · evidence-stale · evidence-invalid`.
```text
evidence available ≠ evidence current
```
**SUP-3 benötigt aktuelle und scopegebundene Evidenz.**

## 11. Compatibility Status

`not-assessed · expected-compatible · compatible-with-notes · validated-compatible · incompatible · deprecated-compatibility · unknown · conflicted`.
```text
same Contract version ≠ same capability set
newer version         ≠ automatically compatible
expected-compatible   ≠ validated-compatible
unknown               ≠ compatible
conflicted            ≠ safe for privileged enablement
```

## 12. Compatibility Dimensions

Ein Compatibility Claim ist gebunden an mindestens: Domain Pack version · CoreOps product version/range · Integration Contract version · target-system version/range · adapter/agent version (falls anwendbar) · deployment profile · operating mode · required dependencies · known limitations · validation evidence · claim owner · claim date/review trigger.

## 13. Claim Requirements

Jeder Compatibility Claim benötigt Scope (§12) und Evidenz (§10); **keine universelle Compatibility-Aussage ohne Scope**. `unknown`/`conflicted` erlauben keine privilegierte Aktivierung. Ein `validated-compatible` gilt nur im getesteten Scope (`validation evidence must not be generalized beyond its tested scope`).

## 14. Known Limitations

Jedes Pack mit Support-Level ≥ SUP-2 dokumentiert bekannte Limitationen; SUP-3 dokumentiert sie verbindlich als Teil des Scopes. Fehlende Limitations-Angabe ist kein Nachweis von Vollständigkeit.

## 15. Dependencies

Wie [Governance §12](DOMAIN_PACK_GOVERNANCE_MODEL.md): dependency identity/type · required/optional · minimum/bounded compatibility · reason · security relevance · offline availability · fallback · deprecation impact · validation status. Ein Dependency-Update erweitert **nicht** still einen Support-/Compatibility-Claim; externe Cloud-Abhängigkeiten bleiben optional, sofern der Pack-Scope nicht ausdrücklich cloudbezogen ist.

## 16. Support Changes

Support-Level-Änderungen sind auditierbar (Grund, betroffene Versionen, Wirkung, Owner, Datum). Eine Höherstufung (z. B. SUP-2 → SUP-3) benötigt die vollständigen SUP-3-Anforderungen (§6); eine Herabstufung dokumentiert Migrationshinweis.

## 17. Deprecation

Wie [Governance §26 / Policy §19]: reason · affected versions · replacement/migration path · support-level change · maintenance boundary · security-fix boundary · compatibility impact · announcement reference · review/end date · owner.
```text
deprecated          ≠ immediately removed
replacement available ≠ migration completed
end of maintenance  ≠ all deployments updated
```

## 18. Retirement

`retired` erhält **keine** neuen Compatibility/Support Claims; historische Pack-IDs und Evidenz bleiben erhalten (`retired ≠ historical evidence deleted`); retired Pack-IDs werden **nicht** wiederverwendet.

## 19. Community and External Support

Community-/External-Support fällt standardmäßig unter SUP-0/SUP-1; eine Höherstufung ist nicht automatisch. `community report ≠ project validation`; `popular pack ≠ supported pack`. CoreOps stuft externe Packs nicht automatisch als SUP-2/SUP-3 ein.

## 20. Vendor-Support Boundary

`external vendor support ≠ CoreOps support`. Herstellersupport für ein Zielsystem ist keine CoreOps-Support-Zusage für ein Pack. Vendor-Bezüge erzeugen kein Endorsement (Governance §15).

## 21. Offline Support Boundary

Offline-Support-Aussagen sind an Ziel-Umgebung, Provenance, Integrität und Contract-Version gebunden (Companion 3 §16). Keine beliebige Air-Gap-/Klassifiziertnetz-Support-Behauptung.

## 22. Audit and Evidence

Erfasst: support-level change · validation activity · evidence reference · compatibility claim · deprecation/retirement. Trennung `support claim ≠ validation evidence ≠ compatibility for every target ≠ vendor certification`.

## 23. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Support level must not imply implementation or validation.
2. Support level must not imply an SLA.
3. SUP-3 requires version-, target-, profile-, limitation- and evidence-bound scope.
4. Compatibility claims must remain version-, target-, profile- and evidence-bound.
5. `expected-compatible` is not `validated-compatible`; `unknown` is not `compatible`.
6. Validation evidence must not be generalized beyond its tested scope.
7. Community or external origin must not imply support.
8. Vendor support is not CoreOps support.
9. Evidence availability does not imply current evidence.

## 24. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-011, THR-023, THR-030, THR-034. Keine Duplikation, kein Parallelregister.

## 25. Technology Boundary

Nicht ausgewählt/implementiert: SLA-/Support-Tooling, Compatibility-Validation-Runtime, Marketplace/Packaging/Update/Dependency-Resolution.

## 26. Open Questions

- Kriterien für den Übergang SUP-2 → SUP-3 im Detail.
- Formaler Umgang mit `conflicted` Compatibility über mehrere Ziel-Versionen.

## 27. Next Decision

Deployment Control Plane (CO-WP-021) und Foundation Readiness Review (CO-WP-030) referenzieren Support-/Compatibility-Aussagen. Tooling bleibt einer späteren ADR-Runde vorbehalten.
