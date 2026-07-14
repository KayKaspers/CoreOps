# CoreOps – Topology Evidence, Confidence and Conflict Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation topology evidence, confidence and conflict model
> Implementation Status: Not implemented
> Conflict Resolution Engine: Not selected
> Identity Resolution Engine: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-020 (docs-only / topology, relationship, evidence and manual-governance foundation)

## 1. Status

Technologieunabhängiges Modell für **Assertion-Herkunft, Source Trust/Authority, Confidence, Validation, Evidence, Temporal Validity, Completeness und Conflict Handling** im Topology Graph. Companion zu [TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md](TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md) und [TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md](../security/TOPOLOGY_MANUAL_AUTHORITY_AND_DISCLOSURE_POLICY.md). Keine Conflict-/Identity-Resolution-Engine.

## 2. Purpose

Topologie-Wahrheit wird leicht überschätzt: `discovered ≠ trusted`, `high confidence ≠ authoritative truth`, `complete local graph ≠ complete physical topology`, `reconciled ≠ conflict-free`. Dieses Modell trennt Source Trust, Authority, Confidence, Validation, Evidence und Completeness und hält Konflikte sichtbar statt sie still aufzulösen.

## 3. Scope

Assertion Concepts/Origins · Source Trust/Authority · Confidence · Validation · Evidence References/Sufficiency · Source Independence · Temporal Validity · Freshness/Staleness · Completeness · Conflicts/Visibility/Precedence · Identity/Relationship Conflicts · Manual vs Observed · Supersession · Invalidations · Offline Evidence · Audit.

## 4. Non-Goals

- Keine Conflict-Resolution-/Identity-Resolution-Engine, kein Graph-Schema/Storage.
- Keine rechtliche Beweiskraft; keine Behauptung durchgeführter Validierung/Vollständigkeit.

## 5. Assertion Concepts

Eine Relationship/Node Assertion ist eine **behauptete** Aussage über Topologie mit Herkunft, Trust, Confidence, Validity und Evidence. `assertion ≠ validated relationship`.

## 6. Assertion Origins

Wie [Graph Model §11]: discovered · observed · declared · imported · manual · derived · inferred · reconciled · historical · unknown. **Keine Origin-Klasse legt allein Authority/Confidence/Validation fest.** `manual ≠ fact`; `inferred ≠ independently observed`.

## 7. Source Trust

Source Trust State (untrusted · limited · trusted-in-scope · unknown) ist zeitgebunden; kompromittierte Quellen (Bezug THR-008, THR-010, THR-012) verlieren Trust, löschen aber historische Assertions nicht. `discovered ≠ trusted`.

## 8. Source Authority

Source Authority (Recht, für eine Node-/Relationship-Klasse autoritativ zu sein) ist **scopegebunden** und getrennt von Trust. `source trust ≠ source authority`; `source authority ≠ global authority`.

## 9. Confidence

Confidence (not-assessed · low · medium · high · source-reported · derived · conflicted · unknown) je Assertion. `high confidence ≠ authoritative truth`; keine numerische Skala erzwungen.

## 10. Validation

Validation (not-assessed · unvalidated · validated-with-notes · validated · invalid · stale · superseded · conflicted · unknown). `validated assertion ≠ currently fresh`; Validation scopegebunden.

## 11. Evidence References

Assertions referenzieren Evidence über stabile Evidence References (konsistent mit [Evidence Model](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), kein Parallelmodell). `graph record ≠ validated relationship evidence`.

## 12. Evidence Sufficiency

Sufficiency (insufficient · partially-sufficient · sufficient-for-decision · sufficient-with-exception · conflicted · unknown) ist **decision-/scope-/zeitgebunden**. `evidence available ≠ evidence sufficient`.

## 13. Source Independence

`many assertions ≠ independent sources`; mehrere Assertions aus derselben Quelle/über denselben Adapter sind nicht unabhängig. Independence beeinflusst Confidence/Sufficiency.

## 14. Temporal Validity

Wie [Graph Model §19]: first/last observed, observation/recording time, valid-from/until, freshness, staleness, superseded, historical. `last observed ≠ currently valid indefinitely`; `newer timestamp ≠ higher authority`.

## 15. Freshness and Staleness

`recently recorded ≠ recently observed`; `not recently observed ≠ relationship absent`; `stale ≠ invalid automatically`. Freshness ist an Signal-/Source-Frequenz gebunden ([Telemetry Freshness](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md)).

## 16. Completeness

`complete local graph ≠ complete physical topology`; `collector healthy ≠ every source reported`. Completeness ist scopegebunden und kann `unknown` bleiben.

## 17. Conflicts

Konfliktquellen: different node identities · relationship types · directionality · target nodes · validity periods · workspace scopes · manual vs observed · authoritative declaration vs telemetry observation · offline vs current source.

## 18. Conflict Visibility

Konflikte bleiben **sichtbar** (`conflict-detected`/`conflicted`) und werden nicht still aufgelöst; keine automatische Last-Write-Wins; keine Timestamp-Priorität allein. Auflösung ist auditierbar.

## 19. Precedence Boundary

Präzedenz ist **pro Node-/Relationship-Klasse oder Scope** geregelt (z. B. authoritative declaration vor Telemetrie-Inferenz), nicht global-implizit. `reconciled ≠ conflict-free automatically`.

## 20. Identity Conflicts

`same display name/hostname/address/external ID ≠ same canonical node`; Identity-Konflikte (candidate/probable/conflicted) blockieren automatische Merges und — bei Sicherheitsrelevanz — privilegierte Automatisierung (Bezug THR-014, THR-035).

## 21. Relationship Conflicts

Konkurrierende Relationship-Assertions (z. B. `connects-to` vs. `not-connected`, verschiedene Targets) bleiben nebeneinander sichtbar mit Source/Confidence/Validity; keine stille Bevorzugung.

## 22. Manual versus Observed Assertions

`manual ≠ observed fact`; eine manuelle Assertion löscht konkurrierende Observation/Evidence **nicht** (Companion 3 §16). Manuelle und beobachtete Assertions koexistieren mit Herkunft/Confidence.

## 23. Supersession

Neuere Assertions können ältere ablösen (`superseded`); Supersession löscht keine historische Assertion/Evidence. Supersession-Historie bleibt nachvollziehbar.

## 24. Invalidations

Eine Assertion kann invalidiert werden (`invalidated`, z. B. widerrufene Quelle); invalidierte Assertions bleiben historisch sichtbar, sind aber nicht mehr autoritativ.

## 25. Offline Evidence

Offline-Topologie-Evidence folgt Companion 3 §23 (target-environment binding, provenance, integrity, import governance). `offline vs current source` bleibt als Konflikt sichtbar; keine konfliktfreie automatische Reconciliation (Bezug THR-024).

## 26. Audit and Evidence

Assertion-/Resolution-/Conflict-/Supersession-Aktivitäten werden auditiert. Trennung `graph record ≠ validated evidence ≠ proof of physical topology`.

## 27. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Discovered/observed/declared/imported/manual/derived/inferred assertions remain separate; origin alone sets no authority.
2. Source trust is not source authority; source authority is scope-bound, not global.
3. High confidence is not authoritative truth; many assertions are not independent sources.
4. Validated is not fresh; evidence available is not sufficient; sufficiency stays decision-/scope-/time-bound.
5. Complete local graph is not complete physical topology; completeness may remain unknown.
6. Conflicts remain visible; no silent last-write-wins or timestamp-only precedence.
7. Same name/address/external ID does not imply the same canonical node.
8. Supersession and invalidation must preserve historical assertions and evidence.
9. Unresolved security-relevant identity/relationship conflicts must block unsafe privileged automation.

## 28. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md): THR-008, THR-010, THR-012, THR-013, THR-014, THR-015, THR-024, THR-026, THR-035. Keine Duplikation, kein Parallelregister.

## 29. Technology Boundary

Nicht ausgewählt/implementiert: Conflict-/Identity-Resolution-Engine, Confidence-Scoring-Technologie, Graph-Schema/Storage, Runtime-Code.

## 30. Compatibility

Konsistent mit [Graph Model](TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md), [Evidence Model](EVIDENCE_REFERENCE_VALIDATION_AND_LINEAGE_MODEL.md), [SoT-Konfliktmodell](SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md), [Offline Reconciliation](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md), [Telemetry](TELEMETRY_SIGNAL_AND_NORMALIZATION_MODEL.md). Konkretisiert DEC-S-114 (Conflict visibility) und die Evidence-Sufficiency-Linie.

## 31. Open Questions

- Präzedenzregeln je Node-/Relationship-Klasse im Detail.
- Independent-Source-Kriterien für Topologie-Confidence.
- Automatisierte Konflikterkennung (spätere ADR).

## 32. Next Decision

Companion 3 trägt Manual Authority/Disclosure. Cross-Document Consistency Review (CO-WP-029) prüft Konsistenz. Resolution-Engine-Wahl bleibt einer späteren ADR-Runde vorbehalten.
