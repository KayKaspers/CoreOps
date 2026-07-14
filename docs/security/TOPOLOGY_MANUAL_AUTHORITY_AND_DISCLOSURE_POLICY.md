# CoreOps – Topology Manual Authority and Disclosure Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation topology manual-authority and disclosure policy
> Implementation Status: Not implemented
> Manual Workflow Technology: Not selected
> Graph Storage: Not selected
> Disclosure Automation: Not implemented
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-020 (docs-only / topology, relationship, evidence and manual-governance foundation)

## 1. Status

Technologieunabhängige Policy für **Manual Authority, Overrides, Suppression, Merge/Split-Autorität, Workspace-Isolation, Disclosure/Export und Offline** im Topology Graph. Companion zu [TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md](../architecture/TOPOLOGY_GRAPH_AND_RELATIONSHIP_MODEL.md) und [TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md](../architecture/TOPOLOGY_EVIDENCE_CONFIDENCE_AND_CONFLICT_MODEL.md). Keine Manual-Workflow-/Graph-Storage-/Disclosure-Technologie.

## 2. Purpose

Manuelle Topologie-Eingriffe sind mächtig und gefährlich: sie dürfen konkurrierende Beobachtungen nicht löschen, keine globale Autorität erzeugen und keine Ausführung freigeben. Diese Policy legt fest, dass `manual ≠ fact`, `manual authority ≠ global topology authority`, `suppressed from view ≠ relationship does not exist` und `topology ≠ execution authorization`.

## 3. Scope

Authority Model · Manual Authority · Permitted Actions · Manual Assertions/Corrections · Overrides · Suppression · Merge/Split Authority · Approval · Review/Expiry · Machine-Principal Boundary · Competing Observations · Conflict Handling · Execution Boundary · Workspace Isolation · Sensitive Topology Data · Disclosure/Export · Offline · Audit · Fail-Closed.

## 4. Non-Goals

- Keine Manual-Workflow-/Graph-Storage-/Disclosure-/Export-Technologie.
- Keine Behauptung implementierter/validierter Kontrollen; kein Runtime-Code.

## 5. Authority Model

Getrennt: Topology Owner (je Node-/Relationship-Klasse/Scope) · Manual Authority (human) · Source Authority (System, scopegebunden, Companion 2 §8) · Recording Component · Execution Authority (CO-WP-013). `manual authority ≠ global topology authority`; `platform administrator ≠ automatic relationship authority`.

## 6. Manual Authority

Manual Authority ist gebunden an: human principal · accountable owner · workspace/scope · node/relationship classes · permitted manual actions · reason · evidence expectation · validity boundary · review trigger · approval requirement (falls anwendbar) · audit reference. **Human-attributable, scope-bound, temporär/reviewgebunden.**

## 7. Permitted Manual Actions

`create assertion · correct metadata · declare relationship · suppress display · override calculated view · merge nodes · split node · reject identity match · mark confidence · mark known limitation · withdraw manual assertion`. Jede Aktion ist scope-/klassengebunden und auditiert.

## 8. Manual Assertions

`manual assertion ≠ observed fact`; manuelle Assertions tragen Herkunft `manual`, Owner, Reason, Evidence-Erwartung; sie koexistieren mit beobachteten Assertions (Companion 2 §22).

## 9. Manual Corrections

Metadaten-/Alias-Korrekturen sind explizit, begründet, auditiert; `manual correction ≠ permission to delete source evidence`.

## 10. Overrides

Ein Override benötigt: override identity · owner · scope · affected node/assertion · original calculated/observed state · replacement/presentation decision · reason · evidence · approval (wo nötig) · validity boundary · review trigger · conflict state · audit reference.
```text
override active   ≠ observed assertion deleted
override           ≠ source evidence invalidated automatically
override expiry    ≠ observed relationship becomes authoritative automatically
```
Overrides sind temporär/reviewgebunden; dauerhafte semantische Änderungen brauchen Governance-Entscheidung.

## 11. Suppression

Suppression beeinflusst Darstellung/Alarmierung/Consumer-Views, darf aber **nicht** still: Source Assertions löschen · Audit History entfernen · Cross-Workspace-Leakage erlauben · Resource Identity verändern · Relationship Authority erzeugen.
```text
suppressed from view ≠ relationship does not exist
```
Suppression benötigt Owner · Scope · Reason · Consumer/View · Validity · Review · Audit.

## 12. Merge and Split Authority

Merge/Split (Graph Model §17-18) benötigen explizite human Entscheidung, Scope, Owner, Evidence, Conflict-Behandlung, Reversibility/Forward-Recovery-Bewertung, Audit. `merge completed ≠ historical identities deleted`; `split completed ≠ original evidence discarded`.

## 13. Approval Requirements

Sicherheitsrelevante manuelle Aktionen (z. B. Merge über Workspace-Grenzen, Override sicherheitsrelevanter Relationships) benötigen Approval (CO-WP-013); Approval ist scope-/plan-gebunden.

## 14. Review and Expiry

Manuelle Assertions/Overrides/Suppressions sind review-/ablaufgebunden; abgelaufene manuelle Eingriffe machen beobachtete Assertions **nicht** automatisch autoritativ (§10).

## 15. Machine-Principal Boundary

**Machine Principals dürfen keine menschliche Manual Authority imitieren**; automatisierte Prozesse erzeugen keine manuellen „Fakten". Machine-erzeugte Assertions tragen Herkunft `derived`/`inferred`/`observed`, nicht `manual`.

## 16. Competing Observations

Manuelle Eingriffe löschen konkurrierende Observations/Evidence **nicht**; diese bleiben mit Herkunft/Confidence sichtbar (Companion 2 §22). `manual override ≠ observed assertion deleted`.

## 17. Conflict Handling

Ungelöste sicherheitsrelevante Identity-/Relationship-Konflikte bleiben sichtbar und **blockieren privilegierte Automatisierung** (Companion 2 §17-21). Keine automatische Auflösung durch manuelle Autorität ohne Evidence/Audit.

## 18. Execution Boundary

```text
topology graph        ≠ execution plan
dependency path       ≠ approval to remediate
connectivity assertion ≠ permission to configure network
node selection        ≠ authorised target scope
```
Topology unterstützt Analyse/Planung; jede Write-/Execution-Aktion bleibt an [CO-WP-013](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) gebunden. Conflicted/unresolved Scope blockiert privilegierte Ausführung.

## 19. Workspace Isolation

Workspace Scope an Nodes/Assertions gebunden; Cross-Workspace-Beziehungen/-Merges nur mit expliziter Governance; Correlation/Shared Nodes erzeugen keine globale Disclosure/Authority (Bezug THR-035).

## 20. Sensitive Topology Data

Netzwerk-/Standort-/Abhängigkeitsdaten gelten als sensitiv; precise Location/personenbezogene Daten nur bei Notwendigkeit; Ressourcenexistenz nicht unautorisiert offengelegt (auch nicht via Beziehungen/Shared Nodes). Bezug THR-018, THR-040.

## 21. Disclosure

Disclosure scope-/authority-gebunden; `read permission ≠ export permission`; consumer-safe/administrative Views getrennt.

## 22. Export

Topology-Export benötigt eigene Autorität, Scope, Redaction-Anforderung, Provenance-/Workspace-Erhalt, Integrity-Status, Validity-Boundary, Audit (konsistent mit [Audit Disclosure](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md)). `export evidence ≠ recipient authority`. Keine Export-/Redaction-Technologie ausgewählt.

## 23. Offline Topology

Offline-Topology-Pakete benötigen: producer/source identity · target environment · workspace/scope · node/relationship assertions · source schema · observation boundary · clock uncertainty · sequence context · provenance/integrity/confidence/validation status · partial package state · duplicate/replay context · import quarantine · explicit import · conflict detection · reconciliation. Nicht behauptet: implementiert, konfliktfreie automatische Reconciliation, beliebige Air-Gap-Stufen, Klassifiziertnetz-Eignung, konkrete Signing-Technologie (Bezug THR-024). Fail-closed bei unklarer Provenance.

## 24. Audit and Evidence

Manuelle Aktionen/Overrides/Suppressions/Merges/Splits/Exports/Offline-Importe sind auditiert; `manual decision record ≠ proof of physical topology`; Historie nicht umgeschrieben (Bezug THR-016, THR-017).

## 25. Fail-Closed Rules

Keine privilegierte Automatisierung/Offenlegung bei: unresolved/conflicted Identity oder Relationship Scope · unklarer Source/Manual Authority · fehlender Approval für sicherheitsrelevante manuelle Aktion · unklarer Workspace-Grenze · fehlender Offline-Provenance · Topology-getriebener Ziel-Scope-Erweiterung ohne Execution Authorization (CO-WP-013).

## 26. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Manual authority remains human-attributable, scope-bound and reviewable.
2. Machine principals must not imitate human manual authority.
3. Manual assertions are not observed facts.
4. Manual overrides must not delete competing observations or evidence.
5. Suppression from a view must not imply relationship absence.
6. Merge and split must preserve historical identities, assertions and evidence.
7. Topology assertions, graph paths and node selection do not grant state, network, configuration or execution authority.
8. Unresolved identity or relationship conflicts block unsafe privileged automation.
9. Cross-workspace relationships, disclosure and export require explicit authority.
10. Topology exports must preserve workspace, provenance and disclosure scope.
11. Offline topology import requires provenance, integrity, target binding and explicit governance.
12. Manual corrections must not delete source evidence or audit history.

## 27. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-002, THR-008, THR-010, THR-012, THR-014, THR-016, THR-017, THR-018, THR-024, THR-026, THR-035, THR-037, THR-040. Keine Duplikation, kein Parallelregister.

## 28. Technology Boundary

Nicht ausgewählt/implementiert: Manual-Workflow-Technologie, Graph-Storage, Disclosure-/Export-/Redaction-Automation, Runtime-Code.

## 29. Compatibility

Konsistent mit Graph-/Evidence-Companion, [Human Identity/RBAC](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md)/[Workspace RBAC](WORKSPACE_RBAC_AND_SCOPE_MODEL.md), [Machine Identity](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [Policy/Approval/Execution](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Audit Disclosure](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md), [Threat Model](COREOPS_FOUNDATION_THREAT_MODEL.md), [Public Neutrality/Disclosure](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md). Konkretisiert DEC-P-04, DEC-G-06, DEC-G-07, DEC-P-08.

## 30. Open Questions

- Approval-Schwellen für sicherheitsrelevante manuelle Topologie-Aktionen.
- Precise-Location-Disclosure-Governance (CO-WP-025).
- Manual-Workflow-Ansatz (spätere ADR).

## 31. Next Decision

Data Classification/Retention (CO-WP-025) und Cross-Document Consistency Review (CO-WP-029) konkretisieren Disclosure-/Konsistenzaspekte. Workflow-/Storage-Wahl bleibt einer späteren ADR-Runde vorbehalten.
