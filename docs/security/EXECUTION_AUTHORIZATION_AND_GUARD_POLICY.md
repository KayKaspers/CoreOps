# CoreOps – Execution Authorization and Guard Policy

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation execution-authorization policy
> Implementation Status: Not implemented
> Execution Engine: Not selected
> Authorization Mechanism: Not selected
> Replay Protection Mechanism: Not selected
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-013 (docs-only / security-baseline)

## 1. Status

Dieses Dokument definiert das technologieunabhängige Foundation-Modell für **Execution Intent, Execution Plan, Execution Authorization, Pre-Execution Guards, Execution Boundary, Replay-/Duplicate-Handling, Expiry/Revocation/Consumption, Break Glass, Offline Authorization, Execution Results, Verification und Closure**. Dritte Stufe der Kette Policy → Approval → Execution Authorization. Companion zu [POLICY_DECISION_AND_EVALUATION_MODEL.md](POLICY_DECISION_AND_EVALUATION_MODEL.md) und [APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md).

## 2. Purpose

Weder ein Policy-`permit` noch eine Approval genügen, um zu handeln. Execution Authorization ist eine **explizite, begrenzte** Autorisierung zur Ausführung eines **bestimmten Plans** oder einer klar definierten Aktion. Dieses Dokument legt fest, wie diese Autorisierung entsteht, woran sie gebunden ist, welche Guards vor privilegierter Ausführung greifen und warum `executed ≠ successful ≠ verified ≠ compliant`.

## 3. Scope

Execution Intent/Plan · Execution Authorization und Lifecycle · Scope-/Target-Bindung · Pre-Execution Guards · Execution Boundary · Replay/Duplicate · Expiry/Revocation/Consumption · Break Glass · Offline Authorization · Execution Results · Verification/Closure · Audit/Evidenz · Fail-Closed-Regeln.

## 4. Non-Goals

- Keine Execution Runtime, kein Queue-/Event-Bus, keine Workflow-Engine.
- Kein Autorisierungsartefakt-/Token-/Ticketformat.
- Kein Replay-Schutz-Mechanismus (Nonce/Idempotency-Key), keine Kryptografie-/Signaturauswahl.
- Keine tatsächliche Ausführung, kein Runtime-Code, keine Codegenerierung.
- Keine automatisierte Approval-Entscheidung.
- Keine Behauptung implementierter/validierter/zertifizierter Kontrollen.

## 5. Authority Model

- **Trennung:** Policy Evaluation (Companion 1) · Approval (Companion 2) · Execution Authorization · Execution · Verification sind getrennte Verantwortlichkeiten.
- Execution Authorization darf **nicht allein aus einem `permit`** entstehen; sie setzt eine anwendbare Policy-Decision **und** — wo verlangt — eine gültige Approval voraus.
- Ausführung erfolgt nur über die vorgesehenen Execution-/Adapter-/Agentengrenzen (§12), niemals durch Experience-, Policy- oder Evidence-Module.
- Bezug: DEC-G-07 (Control Plane nicht umgehbar), DEC-P-04 (Read-only vor Write), [Logical Module Architecture](../architecture/COREOPS_LOGICAL_MODULE_ARCHITECTURE.md), [Safe Remediation Policy](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md).

## 6. Execution Intent

Execution Intent beschreibt die **gewünschte autorisierte Wirkung** (was soll erreicht werden), unabhängig von den konkreten Schritten. Ein Intent ist noch kein Plan und keine Autorisierung.
```text
intent ≠ plan ≠ approval ≠ authorization
```

## 7. Execution Plan

Ein Execution Plan beschreibt die geplanten konkreten Schritte und referenziert mindestens:
```text
plan identity · intent reference · principal · target resources · action sequence or action classes ·
adapter or agent path · credential references · desired-state or remediation reference · preconditions ·
impact · rollback or recovery expectation · verification requirement · policy-evaluation reference ·
approval requirement · audit reference
```
Grundregeln:
```text
intent ≠ plan
plan   ≠ approval
approved plan ≠ executed plan
material plan change → policy re-evaluation
```

## 8. Execution Authorization

Execution Authorization ist eine explizite, begrenzte Autorisierung zur Ausführung eines bestimmten Plans/einer klar definierten Aktion, mindestens gebunden an:
```text
authorization identity · accountable human owner · executing principal · policy version · policy decision ·
approval decision · execution plan · action · target resources · workspace or environment · resource scope ·
validity period · usage count or consumption rule · credential-state expectation ·
desired/effective-state expectation · conflict state · offline state · revocation state · audit reference
```
**Grundregel:** Execution Authorization entsteht **nicht allein aus einem `permit`**. Fehlt eine erforderliche Approval, eine gültige Policy-Version oder ein bestimmter Plan, entsteht keine Autorisierung.

## 9. Authorization Lifecycle

Konzeptionelle Statuswerte:
```text
requested · evaluation-pending · approval-pending · authorized · suspended · revoked · expired ·
consumed · execution-started · execution-completed · verification-pending · closed · invalidated
```
Grundregeln:
```text
authorized        ≠ executing
execution-started ≠ execution-completed
execution-completed ≠ successful
closed            ≠ successful ≠ verified
```
`closed` benötigt: Closure Reason · Final Outcome · Verification Status · verbleibende Exception/Risk · verantwortlichen Principal · Audit Reference.

## 10. Scope and Target Binding

Autorisierung ist mindestens gebunden an: action · resource identity · resource class (falls anwendbar) · workspace · environment · execution-plan version · credential expectation · applicable policy version · approval decision. **Nicht zulässig:**
- Autorisierung für Ressource A wird für Ressource B verwendet.
- Scope-Erweiterung ohne neue Entscheidung.
- Adapter wählt zusätzliche Targets selbst.
- Agent erweitert Action Scope.
- Globale Wirkung aus lokaler Workspace-Autorisierung.

## 11. Pre-Execution Guards

Vor privilegierter Ausführung wird mindestens geprüft:
```text
executing principal active · credential state acceptable · policy version valid · policy decision applicable ·
approval valid · authorization active · authorization not expired · authorization not revoked ·
authorization not consumed · plan version matches · targets match · scope matches ·
desired or remediation reference valid · effective state sufficiently determined · no blocking authority conflict ·
required provenance available · adapter or agent trust acceptable · offline package valid for target environment ·
audit start record possible
```
**Bei sicherheitsrelevanter Unklarheit gilt fail-closed** (§20). Die Effective-State-Bestimmung folgt dem [State-Modell](../architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md): `indeterminate`/`conflicted` blockiert privilegierte Ausführung.

## 12. Execution Boundary

Ausführung darf nur über die vorgesehenen Execution-/Adapter-/Agentengrenzen erfolgen. **Nicht zulässig:**
```text
Experience module executes directly
Policy module executes directly
Approval decision triggers target execution by itself
Adapter self-grants write authority
Agent expands its own scope
Evidence record becomes an execution trigger
Notification becomes a control command
Offline import executes immediately
```
Bezug: THR-005, THR-006, THR-007, THR-034, DEC-G-07.

## 13. Replay and Duplicate Handling

Mindestens berücksichtigen: repeated/delayed/duplicated execution request · replayed approval · replayed authorization · replayed offline package · restarted execution · retry after unknown result. Grundregeln:
- Retry ist **nicht** automatisch Replay.
- Ein **Unknown Result darf nicht automatisch erneut ausgeführt** werden.
- Wiederholte Ausführung benötigt nachvollziehbare Retry- oder neue Authorization-Governance.
- Ursprüngliche Request-/Authorization-/Result-Historie bleibt erhalten.
- Es wird **kein** konkreter Nonce-/Token-/Idempotency-Key-Mechanismus ausgewählt (Bezug THR-004, THR-026).

## 14. Expiry, Revocation and Consumption

- **Expiry** beendet die Gültigkeit nach einer definierten Grenze; Ablauf wird vor privilegierter Ausführung ausgewertet.
- **Revocation** beendet Autorität vor Ablauf; widerrufene Autorisierung bleibt **nicht** autoritativ.
- **Consumption** beendet/reduziert die Wiederverwendbarkeit gemäß Usage Rule; eine konsumierte Autorisierung autorisiert **keinen** Replay.
```text
single-use authorization → must not be reusable silently
revoked authorization    → must not remain authoritative
consumed authorization   → must not authorise a replay
expiry                   → must be evaluated before privileged execution
```
Konkrete technische Durchsetzung bleibt nicht implementiert.

## 15. Break Glass

Break Glass kann eine außergewöhnliche Autorisierungsgrundlage liefern, ersetzt aber nicht Identitätszurechenbarkeit, Scope-Bindung, Grund, Ablauf, Audit, Revocation oder Post-Event Review.
```text
break glass ≠ permanent permit ≠ anonymous execution ≠ removal of original policy history ≠ unlimited scope
```
Die bestehende [Break-Glass-Policy](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md) bleibt authoritative; kein Parallelmodell.

## 16. Offline Authorization

Offline/isolierte Umgebungen können vorab/lokal genehmigte Autorisierungen benötigen. Mindestens: target-environment binding · scope binding · action binding · plan binding · accountable human owner · expiry · usage boundary · provenance requirement · integrity requirement · local activation decision · revocation-distribution challenge · audit continuity · post-reconnection reconciliation. Nicht behauptet: implementierte Offline Authorization, beliebige Air-Gap-Stufen, Eignung für klassifizierte Netze, konkrete Signing-/Token-/Trust-Anchor-Technologie. Bezug THR-024, THR-027.

## 17. Execution Results

- Ein Execution Result ist das **beobachtete** Ergebnis, nicht der Nachweis von Erfolg.
```text
execution request ≠ execution result
execution result  ≠ verified success
```
- **Unknown Result** ist weder `successful` noch `failed`; es löst keine automatische erneute Ausführung aus (§13).
- **Partial Failure** bleibt sichtbar; ein Teilerfolg ist **kein** vollständiger Erfolg (Bezug THR-031).

## 18. Verification and Closure

- **Verification** ist ein gesonderter Nachweis, dass das Ergebnis den erwarteten Zielzustand erreicht — möglichst über eine **unabhängige, frische** Quelle (konsistent mit [Safe Remediation §12](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md)).
```text
executed  ≠ successful
successful ≠ verified
closed     ≠ successful ≠ verified
```
- **Rollback** gilt erst nach Verifikation als wiederhergestellt (`rollback completed ≠ restored unless verified`, Bezug THR-032).
- **Closure** erfordert die in §9 genannten Felder; `closed` behauptet weder Erfolg noch Compliance.

## 19. Audit and Evidence

Mindestens erfasst: policy identity/version · evaluation inputs/outcome · conflict status · approval requirement/request/approver/decision · execution plan · authorization · expiry · revocation · consumption · execution start · execution result · verification · break-glass use · offline activation · closure. Trennung:
```text
authorization evidence ≠ execution evidence ≠ success evidence ≠ verification evidence ≠ compliance
```
Audit-Records sind **keine** Ausführungsauslöser; Audit-/Evidence-Historie wird nicht umgeschrieben (Bezug THR-016, THR-017).

## 20. Fail-Closed Rules

Bei sicherheitsrelevanter Unklarheit wird **nicht** privilegiert ausgeführt, insbesondere bei: fehlendem/unklarem Evaluation-Input · `indeterminate`/`conflicted` Policy Decision · abgelaufener/widerrufener/konsumierter Autorisierung · Plan-/Target-Mismatch · unzureichend bestimmtem Effective State · blockierendem Autoritätskonflikt · fehlender Provenance · unklarer Adapter-/Agent-Trust · ungültigem Offline-Paket · unmöglichem Audit-Start-Record · unbekanntem/verspätetem Zeit-/Sequenzkontext.

## 21. Security Invariants

Als Designanforderungen (nicht implementierte Kontrollen):

1. Policy permit must not imply execution authorization.
2. Approval must not itself execute an action.
3. Machine principals must not self-approve privileged actions.
4. Execution authorization must remain action-, target-, scope-, plan- and time-bound.
5. Revoked, expired or consumed authorization must not remain usable.
6. Material plan or target change requires re-evaluation.
7. Unknown or conflicted authority must block privileged execution.
8. Break-glass authorization must remain named, temporary and auditable.
9. Offline authorization requires provenance, integrity and explicit activation.
10. Executed must not be interpreted as successful.
11. Successful must not be interpreted as verified; closed must not be interpreted as successful.
12. Audit evidence must not grant execution authority.

## 22. Threat References

Reale IDs aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md): THR-003, THR-004, THR-005, THR-006, THR-007, THR-008, THR-010, THR-016, THR-017, THR-024, THR-026, THR-027, THR-028, THR-029, THR-030, THR-031, THR-032, THR-034, THR-037, THR-038. Keine Duplikation, kein Parallelregister.

## 23. Technology Boundary

Nicht ausgewählt/implementiert: Execution Engine, Queue/Event-Bus, Workflow-Engine, Autorisierungsmechanismus/Token/Artefakt, Replay-Schutz (Nonce/Idempotency), Signatur-/Kryptografie-/Trust-Anchor-Technologie, API-/DB-Schema, Runtime-Code.

## 24. Compatibility

Konsistent mit Companion 1/2, [Safe Remediation Policy](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md) (Detection/Recommendation/Plan/Approval/Execution/Verification getrennt), [State-Modell](../architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [Trust/Deployment/Execution Boundaries](TRUST_DEPLOYMENT_AND_EXECUTION_BOUNDARIES.md), [Machine Identity](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [Break Glass](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md). Konkretisiert DEC-P-04, DEC-P-05, DEC-G-04, DEC-G-07.

## 25. Open Questions

- Verhältnis von Execution Authorization zu einem späteren Job-/Task-Modell (CO-WP-021).
- Konkreter Replay-Schutz (spätere ADR).
- Unabhängige Verifikationsquellen je Aktionsklasse.

## 26. Next Decision

Integration Contract (CO-WP-014), Deployment Control Plane (CO-WP-021), Artifact Trust (CO-WP-022) und Secrets/Key Custody (CO-WP-024) konkretisieren einzelne hier referenzierte Grenzen. Engine-/Mechanismuswahl bleibt einer späteren ADR-Runde vorbehalten; in diesem WP wird keine Technologie ausgewählt.
