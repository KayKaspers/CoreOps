# CoreOps – Domain Pack Trust, Provenance and Lifecycle Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation Domain-Pack trust and lifecycle policy
> Implementation Status: Not implemented
> Pack Verification Mechanism: Not selected
> Cryptographic Mechanism: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-015 (docs-only / modular product-governance and compatibility foundation)

## 1. Status

Technologieunabhängige Policy für **Vertrauen, Herkunft, Integrität, Security-Response und Lebenszyklus** von Domain Packs. Companion zu [DOMAIN_PACK_GOVERNANCE_MODEL.md](../architecture/DOMAIN_PACK_GOVERNANCE_MODEL.md) und [DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md](../architecture/DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md). Kein Pack-Verification-Mechanismus, keine Kryptografie ausgewählt.

## 2. Purpose

Packs sind eine Supply-Chain-Fläche. Diese Policy legt fest, wie wenig einer Pack-Herkunft/-Metadaten vertraut wird, wie kompromittierte Packs/Maintainer suspendiert werden und warum `downloaded ≠ trusted`, `public availability ≠ verified provenance` und `security reviewed ≠ vulnerability free`.

## 3. Scope

Trust Boundary · Pack Source · Ownership/Maintainers · Provenance · Integrity · Dependency Trust · Community/External Packs · Security Review · Vulnerability/Incident Response · Suspension · Withdrawal/Revocation · Offline Distribution · Support-/Compatibility-Impact · Deprecation/Retirement · Audit/Evidence · Fail-Closed.

## 4. Non-Goals

- Kein Pack-Verification-/Signatur-/Trust-Anchor-Mechanismus, keine Kryptografieauswahl.
- Kein Marketplace/Packaging/Update/Distribution, keine Dependency-Resolution-Engine.
- Keine Behauptung implementierter/validierter/zertifizierter Kontrollen; kein SLA.

## 5. Trust Boundary

Ein Domain Pack — insbesondere von externer Herkunft — liegt jenseits einer Vertrauensgrenze. `enrolled ≠ trusted`; `trusted ≠ globally authorised`. Pack-Vertrauen ist scope-/version-gebunden und widerrufbar. Pack-Aktivierung gewährt **keine** Runtime-Autorität (Governance §18).

## 6. Pack Source

Erfasst: pack source · source reference · owner · maintainer · version · release reference · dependency references. `public repository ≠ verified provenance`; `popular pack ≠ trusted/supported pack`.

## 7. Ownership and Maintainers

Owner/Maintainer sind zurechenbar dokumentiert (ohne private Kontaktdaten in öffentlichen Beispielen). Maintainer-Wechsel sind auditierbar. `external maintainer ≠ CoreOps endorsement`.

## 8. Provenance

```text
downloaded    ≠ trusted
parsed        ≠ validated
pack metadata ≠ verified pack artifact
```
Provenance-Status ist explizit; fehlende Provenance ⇒ kein automatisches Vertrauen (fail-closed §21). Konsistent mit [Field Provenance](../architecture/FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) (kein Parallelmodell).

## 9. Integrity

```text
integrity checked ≠ security reviewed
```
Integritätsprüfung (Mechanismus deferred) belegt Unverändertheit, nicht Sicherheit. Kein konkreter Signatur-/Hash-Mechanismus ausgewählt (Bezug THR-021, THR-023).

## 10. Dependency Trust

Dependencies erben **kein** Vertrauen automatisch; jede Dependency trägt eigene Provenance/Validation. Optionale Dependencies werden nicht still zur Core-Pflicht; externe/cloud Dependencies bleiben optional außerhalb cloudbezogener Pack-Scopes (Bezug THR-023).

## 11. Community and External Packs

`community pack ≠ trusted pack`. Externe Packs sind standardmäßig SUP-0/SUP-1, nicht automatisch vertrauenswürdig/supported. Offenlegung: source · owner · maintainer model · governance status · support level · compatibility claims · validation state · known limitations · dependency provenance · security reporting path · deprecation policy.

## 12. Security Review

```text
security reviewed ≠ vulnerability free
```
Security-Review-Status ist eine eigene Dimension (Governance §8); ein Review zu einem Zeitpunkt gilt nicht unbegrenzt. Fehlender Review ⇒ kein SUP-3.

## 13. Vulnerability and Incident Response

Für aktive/unterstützte Packs dokumentiert: Security-Owner · Meldeweg · Triage-Status · betroffene Versionen/Capabilities · Workaround · Fix-/Mitigation-Status · Support-Level-Auswirkung · Compatibility-Auswirkung · Retirement-/Suspension-Entscheidung · Evidenzreferenz. **Keine SLA/Reaktionszeit versprochen.** Ein Security-Finding **muss** Support-/Lifecycle-Status ändern können (Bezug THR-019, THR-020, THR-023).

## 14. Suspension

Ein kompromittierter Pack oder Maintainer **muss suspendierbar** sein (`suspended`, Governance §9). Suspension stoppt neue Aktivierungen/Claims; bestehende Deployments werden gemäß Incident-Response behandelt. Auditierbar.

## 15. Withdrawal and Revocation

Ein Pack/eine Pack-Version kann zurückgezogen/widerrufen werden. Widerrufene Packs bleiben **nicht** autoritativ und erhalten keine neuen Claims. Revocation-Information muss auch offline zugestellt/berücksichtigt werden (§16, Revocation-Distribution-Challenge).

## 16. Offline Distribution

Offline Packs/Metadaten benötigen: pack identity · version · target environment · dependency set · contract compatibility · provenance/integrity status · validity/usage boundary · import quarantine · local activation decision · revocation-distribution challenge · reconciliation requirement · audit continuity. Nicht behauptet: implementierte Offline Distribution, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie (Bezug THR-024). Fail-closed bei unklarer Provenance/Autorität.

## 17. Support-Level Impact

Trust-/Security-Ereignisse wirken auf den Support Level: ein Security-Finding oder Provenance-Verlust kann SUP-3 → SUP-D/Suspension auslösen. `support level ≠ SLA` bleibt.

## 18. Compatibility Impact

Trust-/Security-Ereignisse können Compatibility Claims invalidieren (`compatibility → deprecated-compatibility/incompatible`). Claims bleiben version-/scope-gebunden.

## 19. Deprecation and Retirement

Wie [Governance §26 / Support §17-18]. `deprecated ≠ immediately removed`; `retired ≠ historical evidence deleted`; retired Pack-IDs nicht wiederverwendet; historische Identität/Evidenz erhalten.

## 20. Audit and Evidence

Erfasst wie [Governance §17]. Trennung `support claim ≠ validation evidence ≠ compatibility for every target ≠ vendor certification`. Audit-/Evidence-Historie wird nicht umgeschrieben (Bezug THR-016, THR-017).

## 21. Fail-Closed Rules

Bei sicherheitsrelevanter Unklarheit **keine** privilegierte Pack-Aktivierung/-Nutzung: fehlende/unklare Provenance · fehlgeschlagene/unbekannte Integrität · kompromittierter/suspendierter Maintainer · widerrufener Pack · unbekannte Ziel-Umgebung für Offline-Paket · `conflicted` Compatibility · unbekannte sicherheitsrelevante Dependency/Extension · fehlender Security-Review für SUP-3.

## 22. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Downloaded or public availability does not imply trust or verified provenance.
2. Pack metadata is not a verified pack artifact; parsed is not validated.
3. Integrity checked is not security reviewed; security reviewed is not vulnerability free.
4. Community or external origin does not imply trust, support or endorsement.
5. A compromised pack or maintainer must be suspendable and revocable.
6. Pack activation must not grant runtime authority automatically.
7. Offline pack use requires provenance, integrity, target binding and explicit activation.
8. A security finding must be able to change support and lifecycle status.
9. Deprecated/retired packs must not receive new unsupported claims; historical evidence must not be deleted.
10. Vendor-specific extensions must not override CoreOps security invariants.

## 23. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-008, THR-010, THR-011, THR-016, THR-017, THR-019, THR-020, THR-021, THR-022, THR-023, THR-024, THR-025, THR-026, THR-035. Keine Duplikation, kein Parallelregister.

## 24. Technology Boundary

Nicht ausgewählt/implementiert: Pack-Verification, Signatur/Hash/Trust-Anchor, Kryptografie, Marketplace/Packaging/Update/Distribution, Dependency-Resolution-Engine.

## 25. Compatibility

Konsistent mit Governance-/Support-Companion, [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md), [Integration Trust/Failure](INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [Machine Identity/Enrollment](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [Offline Reconciliation](OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md), [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md). Konkretisiert DEC-P-02 (Offline), DEC-G-05/06, die Sovereignty-/Dependency-Linie und die NDF-Feedback-Kandidaten zur Supply-Chain-Sicherheit.

## 26. Open Questions

- Konkreter Pack-Verification-/Signatur-Mechanismus (spätere ADR, mit CO-WP-022 verbunden).
- Offline-Revocation-Distribution im Detail.
- Formaler Suspension-/Reactivation-Ablauf.

## 27. Next Decision

Artifact Trust/SBOM/Provenance (CO-WP-022), Restricted/Air-Gapped Operation (CO-WP-023) und Self-Protection/Recovery (CO-WP-026) konkretisieren Trust-/Provenance-Aspekte. Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten.
