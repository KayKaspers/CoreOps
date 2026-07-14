# CoreOps – Approval and Authorization Lifecycle

> Document Status: Implemented, pending Nova review
> Lifecycle Status: Foundation approval and authorization lifecycle
> Implementation Status: Not implemented
> Approval Engine: Not selected
> Authorization Artifact: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-013 (docs-only / security-baseline)

## 1. Status

Dieses Dokument definiert das technologieunabhängige Foundation-Modell für **Approval Requirements, Approval Requests, Approver Authority, Separation of Duties und Approval Decisions** sowie deren Bindung an Scope, Plan, Ablauf, Revocation und Consumption. Companion zu [POLICY_DECISION_AND_EVALUATION_MODEL.md](POLICY_DECISION_AND_EVALUATION_MODEL.md) und [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md). Es wählt **keine** Approval Engine, kein Autorisierungsartefakt und kein Tokenformat.

## 2. Purpose

Ein Policy-`permit` (Companion 1) bedeutet nicht, dass eine privilegierte Handlung genehmigt oder autorisiert ist. Dieses Dokument liefert die zweite Stufe: wie eine Approval-Anforderung entsteht, wer mit welcher Autorität genehmigt, wie Separation of Duties gewahrt bleibt und wie eine Approval-Entscheidung scope-, plan- und zeitgebunden, widerrufbar und einmalig-konsumierbar bleibt — als Voraussetzung, aber nicht als Ersatz für die Execution Authorization (Companion 3).

## 3. Scope

Approval Requirements · Approval Requests · Approver Authority · Separation of Duties · Approval Decisions und deren Lifecycle · Scope-/Plan-Bindung · Expiry · Revocation · Consumption · Replay-Grenze · Break-Glass- und Offline-Bezug · Audit/Evidenz.

## 4. Non-Goals

- Keine Approval Engine, kein Workflow-Produkt, keine Queue.
- Kein Autorisierungsartefakt-, Token- oder Ticketformat.
- Keine kryptografische Signatur-/Nonce-/Idempotency-Key-Auswahl.
- Keine automatisierte (nicht-menschliche) Genehmigung privilegierter Aktionen.
- Keine Behauptung einer implementierten oder validierten Kontrolle.
- Keine pauschale Zwei-Personen-Pflicht für alle Installationen.

## 5. Concepts

Aufbauend auf den Begriffen aus Companion 1 §5. Zusätzlich zentral: `approval requirement` (Forderung), `approval request` (Antrag), `approval decision` (Entscheidung), `approver authority` (Berechtigung), `separation of duties` (personelle/rolliche Trennung), `consumption` (Verbrauch der Wiederverwendbarkeit).

**Grundregeln:**
```text
approval        ≠ execution
approved        ≠ executed
expired approval ≠ valid authorization
revoked approval ≠ reusable approval
approval request ≠ approval decision
machine request  ≠ machine self-approval
```

## 6. Approval Requirements

Eine Approval Requirement definiert mindestens:

```text
requirement identity · triggering policy · subject · requested action · resource and scope ·
risk or impact · required approver role · minimum approver count · separation-of-duties requirement ·
expiry · single-use or reusable boundary · offline applicability · break-glass eligibility · evidence requirement
```

Grundregeln:
- Eine Approval Requirement entsteht aus einer Policy mit Outcome `approval-required` (Companion 1 §10).
- Es wird **keine** pauschale Zwei-Personen-Pflicht für alle Installationen behauptet.
- **Kleine oder einzelne Deployments** dürfen kompensierende Kontrollen verwenden, sofern die reduzierte Trennung **sichtbar und auditierbar** ist (§9).
- Risiko/Impact bestimmen Approver-Rolle, -Anzahl und SoD-Anforderung — qualitativ, ohne mathematische Risikoformel.

## 7. Approval Requests

Ein Approval Request erfasst mindestens:

```text
approval-request identity · requesting human or machine principal · human accountable owner ·
requested action · resource · scope · execution-plan reference · policy-decision reference ·
desired-state or drift reference · reason · impact · requested duration · requested usage count ·
request time · expiry · approver requirement · audit reference
```

Grundregeln:
- Eine **Machine Identity darf einen Request erzeugen**, darf sich aber **nicht selbst genehmigen** (`machine request ≠ machine self-approval`).
- Bei privilegierten Machine Requests bleibt die **menschliche Verantwortlichkeit** (`human accountable owner`) nachvollziehbar.
- Ein Request referenziert die konkrete `execution-plan reference` und `policy-decision reference`; ändert sich der Plan/Target, ist eine erneute Entscheidung nötig (§10).

## 8. Approver Authority

Eine Approver Identity benötigt:

```text
valid human identity · active account and membership · applicable role · applicable scope ·
current authority · no conflicting lifecycle state · required separation from requester where applicable · audit attribution
```

Grundregeln:
```text
workspace owner        ≠ automatic global approver
platform administrator ≠ universal risk acceptor
Human Maintainer       ≠ runtime approver by default
machine principal      ≠ human approval authority
approval role name     ≠ unlimited approval scope
```
Die Approver-Autorität ist stets **scope-gebunden** (Workspace/Environment/Resource-Klasse) und wird gegen die Identitäts-/RBAC-Governance ([HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md), [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](WORKSPACE_RBAC_AND_SCOPE_MODEL.md)) geprüft.

## 9. Separation of Duties

Mindestens zu prüfen:
```text
requester        ≠ approver
policy author    ≠ final policy approver
execution planner ≠ sole approver
approver         ≠ executor where required
break-glass user ≠ exclusive post-event reviewer
audit subject    ≠ exclusive audit reviewer
```

**Kleine Deployments:**
- Reduzierte personelle Trennung ist zulässig.
- Die Abweichung muss **sichtbar** sein.
- Kompensierende Kontrollen müssen **dokumentiert** sein.
- Es wird **keine** automatische Behauptung vollständiger Separation of Duties erhoben.

Bezug: THR-002, THR-037 (Privilege-/Insider-Abuse).

## 10. Approval Decisions

Konzeptionelle Statuswerte:
```text
pending · approved · denied · expired · revoked · superseded · consumed · cancelled · indeterminate
```

Grundregeln:
```text
approved        ≠ executed
expired approval ≠ valid authorization
revoked approval ≠ reusable approval
scope change    → new approval decision
plan change     → re-evaluation and potentially new approval
target change   → new approval
```
- Approval-Historie darf **nicht** durch spätere Entscheidungen umgeschrieben werden.
- `indeterminate` (z. B. unklare Autorität/Inputs) begründet **keine** gültige Approval.
- `consumed` markiert eine verbrauchte Single-Use-Approval (§14).

## 11. Scope and Plan Binding

Eine Approval Decision ist gebunden an: Action · Resource/Resource-Klasse · Workspace · Environment · Execution-Plan-Version · Policy-Version · angefragte Dauer/Usage. Grundregeln:
- Eine Approval für Ressource A gilt **nicht** für Ressource B.
- Eine **materielle** Plan- oder Target-Änderung erfordert Re-Evaluation (Companion 1) und ggf. eine neue Approval.
- Scope-Erweiterung ohne neue Entscheidung ist unzulässig.

## 12. Expiry

- Jede Approval hat einen definierten Ablauf.
- Eine abgelaufene Approval ist **keine** gültige Autorisierungsgrundlage.
- Ablauf wird **vor** privilegierter Nutzung geprüft (Companion 3 §11).
- Bei unsicherer/false Zeitquelle gilt fail-closed (Bezug THR-027).

## 13. Revocation

- Eine Approval kann **vor** Ablauf widerrufen werden.
- Eine widerrufene Approval bleibt **nicht** wiederverwendbar und **nicht** autoritativ.
- Revocation-Information muss auch in offline/isolierten Umgebungen zugestellt/berücksichtigt werden (§17, Revocation-Distribution-Challenge).

## 14. Consumption

- Eine Single-Use-Approval wird durch Nutzung **konsumiert** und ist danach nicht mehr verwendbar (`consumed authorization must not authorise a replay`).
- Eine reusable Approval bleibt nur im deklarierten `usage count`/Zeitfenster gültig.
- Konsum/Verbrauch werden auditiert; die technische Durchsetzung bleibt **nicht** implementiert.

## 15. Replay Boundary

Mindestens berücksichtigen: wiederholter/verzögerter/duplizierter Request · replayed approval · replayed offline package. Grundregeln:
- Ein **Retry** ist nicht automatisch ein **Replay**.
- Eine replayed oder duplizierte Approval begründet **keine** neue Autorität.
- Ursprüngliche Request-/Decision-Historie bleibt erhalten.
- Es wird **kein** konkreter Nonce-/Token-/Idempotency-Key-Mechanismus ausgewählt (Bezug THR-004, THR-026).

## 16. Break Glass

Break Glass kann eine außergewöhnliche Approval-/Autorisierungsgrundlage liefern, ersetzt aber nicht Identitätszurechenbarkeit, Scope-Bindung, Grund, Ablauf, Audit, Revocation oder Post-Event Review.
```text
break glass ≠ permanent permit
break glass ≠ anonymous execution
break glass ≠ removal of original approval history
break glass ≠ unlimited scope
```
Die bestehende [Break-Glass-Policy](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md) bleibt authoritative; kein Parallelmodell.

## 17. Offline Approval

Offline/isolierte Umgebungen können vorab oder lokal genehmigte Approvals benötigen. Mindestens: target-environment binding · scope binding · action binding · plan binding · accountable human owner · expiry · usage boundary · provenance requirement · integrity requirement · local activation decision · revocation-distribution challenge · audit continuity · post-reconnection reconciliation. Nicht behauptet werden: implementierte Offline-Approval, beliebige Air-Gap-Eignung, Eignung für klassifizierte Netze, konkrete Signing-/Token-/Trust-Anchor-Technologie.

## 18. Audit and Evidence

Mindestens erfasst: approval requirement · approval request · requesting principal · human accountable owner · approver identity · approval decision · scope/plan reference · expiry · revocation · consumption · break-glass use · offline activation. Trennung:
```text
authorization evidence ≠ execution evidence ≠ success evidence ≠ verification evidence ≠ compliance
```
Approval-/Audit-Historie wird nicht umgeschrieben; Audit-Records sind keine Ausführungsauslöser.

## 19. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Approval must not itself execute an action.
2. Machine principals must not self-approve privileged actions.
3. Privileged machine requests must retain traceable human accountability.
4. Approver authority remains scope-bound; a role name does not grant unlimited approval scope.
5. A material plan or target change requires re-evaluation and potentially a new approval.
6. Expired, revoked or consumed approvals are not valid authorization.
7. Approval history must not be rewritten by later decisions.
8. Separation of Duties applies where required; reduced separation must be visible and compensated.
9. Break-glass approvals remain named, temporary, scope-bound and auditable.
10. Offline approvals require target binding, provenance, integrity and explicit activation.
11. A replayed or duplicated approval grants no new authority.

## 20. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-002, THR-003, THR-004, THR-026, THR-027, THR-037, THR-038. Keine Duplikation, kein Parallelregister.

## 21. Technology Boundary

Nicht ausgewählt/implementiert: Approval Engine, Workflow-/Queue-Runtime, Autorisierungsartefakt/Token/Ticket, Signatur-/Nonce-/Idempotency-Mechanismus, API-/DB-Schema, automatisierte Genehmigung.

## 22. Compatibility

Konsistent mit Companion 1/3, [Human Identity/RBAC](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md), [Break Glass](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md), [Machine Identity](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [Safe Remediation](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md). Konkretisiert DEC-P-05, DEC-G-07 und die Approval-Aspekte von DEC-S-84 (sensible Aktionen mit Reauth/Approval).

## 23. Open Questions

- Mindest-Approver-Anzahl je Risikoklasse (qualitativ vs. profilabhängig).
- Formalisierung „sichtbarer" reduzierter SoD in kleinen Deployments.
- Offline-Revocation-Distribution als spätere ADR.

## 24. Next Decision

Die Execution Authorization und die Pre-Execution Guards (Companion 3) konsumieren Approval Decisions als eine notwendige, aber nicht hinreichende Bedingung. Engine-/Artefaktwahl bleibt einer späteren ADR-Runde vorbehalten.
