# CoreOps – Policy Decision and Evaluation Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation policy-decision model
> Implementation Status: Not implemented
> Policy Engine: Not selected
> Policy Language: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-013 (docs-only / security-baseline)

## 1. Status

Dieses Dokument definiert das technologieunabhängige Foundation-Modell für **Policy-Arten, Policy-Lifecycle, Policy Evaluation und Policy Decision Outcomes** in CoreOps. Es ist Companion zu [APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md) und [EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md). Es wählt **keine** Policy Engine, keine Policy-Sprache, kein Token- oder Protokollformat und beginnt keine Implementierung. Die Statuswerte in diesem Dokument sind konzeptionelle Guidance, keine implementierten Enums.

## 2. Purpose

CoreOps trennt **Policy Evaluation** (darf etwas grundsätzlich erlaubt sein?), **Approval** (hat eine berechtigte Instanz die konkrete Handlung genehmigt?) und **Execution Authorization** (existiert eine explizite, begrenzte Autorisierung, genau diesen Plan auszuführen?). Dieses Dokument liefert die erste Stufe: ein Modell, wie Policies klassifiziert, versioniert, ausgewertet werden und zu einem klar definierten, fail-closed-fähigen Decision Outcome führen — ohne dass ein `permit` allein jemals Ausführung begründet.

## 3. Scope

- Policy-Klassen und deren Governance-Attribute.
- Policy-Lifecycle (Draft bis Archived).
- Policy-Scope, -Subjects, -Resources, -Actions, -Context.
- Evaluation Inputs und deren Vollständigkeits-/Vertrauensanforderungen.
- Policy Decision Outcomes und deren Semantik.
- Default-Deny-Grenze, Konflikte und Präzedenz.
- Beziehung zu Exceptions, Break Glass und Offline-Auswertung.
- Audit-/Evidenzanforderungen der Evaluation.

Außerhalb des Scope: Approval-Lifecycle-Details (Companion 2), Execution Authorization und Guards (Companion 3), Technologieauswahl jeder Art.

## 4. Non-Goals

- Keine Auswahl einer Policy Engine (kein OPA/Rego, kein Cedar, kein XACML, kein anderes Produkt).
- Keine vollständige technische Policy-Sprache oder Grammatik.
- Kein Token-, Ticket- oder Autorisierungsartefaktformat.
- Keine kryptografische Signatur- oder Trust-Anchor-Auswahl.
- Kein API-, Datenbank- oder Ereignisschema.
- Keine Workflow-, Queue- oder Execution-Runtime.
- Keine automatisierte Approval-Entscheidung.
- Keine Behauptung einer implementierten oder validierten Kontrolle.

## 5. Concepts

Begriffe (mindestens zu unterscheiden):

| Begriff | Bedeutung |
|---|---|
| policy | Benannte Regelmenge, die für einen Scope Handlungen erlaubt, verbietet oder an Bedingungen knüpft. |
| policy rule | Einzelne Bedingung innerhalb einer Policy. |
| policy set | Gruppierung mehrerer Policies mit gemeinsamem Governance-Kontext. |
| policy version | Konkrete, historisch nachvollziehbare Fassung einer Policy. |
| policy scope | Der Geltungsbereich (Workspace/Environment/Resource-Klasse/Action). |
| policy subject | Principal, auf den die Policy angewandt wird (human/machine). |
| policy resource | Ziel-Ressource oder Ressourcenklasse. |
| policy action | Angefragte Handlung bzw. Handlungsklasse. |
| policy context | Zusätzliche Auswertungsfaktoren (State, Freshness, Trust, Offline, Zeit/Sequenz). |
| policy evaluation | Der Vorgang, für eine Anfrage ein Decision Outcome zu bestimmen. |
| policy decision | Das Ergebnis der Auswertung (`permit`/`deny`/…). |
| approval requirement | Eine durch Policy ausgelöste Forderung nach gesonderter Approval-Entscheidung. |
| approval request | Der erfasste Antrag auf eine Approval-Entscheidung (Companion 2). |
| approval decision | Die Entscheidung über einen Approval Request (Companion 2). |
| approver authority | Die Berechtigung einer Instanz, zu genehmigen (Companion 2). |
| execution intent | Die gewünschte autorisierte Wirkung (Companion 3). |
| execution plan | Die geplanten konkreten Schritte (Companion 3). |
| execution authorization | Explizite, begrenzte Autorisierung zur Ausführung (Companion 3). |
| execution request | Die Anfrage, einen autorisierten Plan auszuführen (Companion 3). |
| execution result | Das beobachtete Ergebnis der Ausführung (Companion 3). |
| verification result | Der gesonderte Nachweis, dass das Ergebnis den Zielzustand erreicht (Companion 3). |
| exception | Explizit erfasste, begründete, begrenzte Abweichung von einer Policy. |
| break-glass authorization | Außergewöhnliche, benannte, temporäre Autorisierungsgrundlage ([BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md)). |

**Grundregeln (Boundary-Kette):**

```text
policy evaluation      ≠ approval
policy permit          ≠ execution authorization
approval               ≠ execution
execution authorization ≠ credential
execution request      ≠ execution result
execution result       ≠ verified success
exception              ≠ silent policy bypass
```

## 6. Policy Classes

Jede Klasse wird mit den geforderten Attributen (Purpose · Owner · Scope · Subjects · Resources · Actions · Inputs · Decision authority · Conflict owner · Offline applicability · Audit requirement · Threat references) beschrieben. Es wird **keine** vollständige technische Policy-Sprache definiert.

### 6.1 Platform Governance Policy
- **Purpose:** plattformweite, nicht umgehbare Grundregeln (z. B. Read-only vor Write, Human-Maintainer-Gates).
- **Owner:** Human Maintainer / Nova (Governance). **Scope:** gesamte Plattform. **Subjects:** alle Principals. **Resources:** alle. **Actions:** alle privilegierten Klassen.
- **Inputs:** Principal-Klasse, Action-Privilegstufe, Lifecycle-State. **Decision authority:** Governance. **Conflict owner:** Human Maintainer.
- **Offline applicability:** gilt auch offline. **Audit requirement:** jede Auswertung privilegierter Aktionen. **Threat references:** THR-002, THR-037.

### 6.2 Workspace Policy
- **Purpose:** Regeln innerhalb eines Workspace-Scopes (siehe [WORKSPACE_RBAC_AND_SCOPE_MODEL.md](WORKSPACE_RBAC_AND_SCOPE_MODEL.md)).
- **Owner:** Workspace Owner (kein globaler Approver). **Scope:** ein Workspace. **Subjects:** Workspace-Mitglieder/-Principals. **Resources:** Workspace-Ressourcen. **Actions:** Workspace-gebundene Aktionen.
- **Inputs:** Workspace, Rolle, Scope, Resource-Klasse. **Decision authority:** Workspace-Governance. **Conflict owner:** Platform-Ebene bei Kollision. **Offline applicability:** teilweise. **Audit requirement:** ja. **Threat references:** THR-002, THR-035.

### 6.3 Environment Policy
- **Purpose:** Regeln je Umgebung (z. B. produktiv vs. isoliert/offline). **Owner:** Nova/Environment-Governance. **Scope:** Environment. **Subjects:** alle im Environment. **Resources:** Environment-Ressourcen. **Actions:** deploymentnahe/privilegierte Aktionen.
- **Inputs:** Environment, Offline-State, Trust-State. **Decision authority:** Environment-Governance. **Conflict owner:** Platform. **Offline applicability:** zentral (siehe §15). **Audit requirement:** ja. **Threat references:** THR-024, THR-027.

### 6.4 Resource Policy
- **Purpose:** Regeln je Ressource/Ressourcenklasse. **Owner:** Resource-/Data-Owner (siehe Modul-Katalog). **Scope:** Ressource(nklasse). **Subjects:** zugreifende Principals. **Resources:** die Ressource. **Actions:** read/write/execute-Klassen.
- **Inputs:** Resource-Identity/-Klasse, Action-Privilegstufe. **Decision authority:** Resource-Owner. **Conflict owner:** Workspace/Platform. **Offline applicability:** teilweise. **Audit requirement:** bei privilegierten Aktionen. **Threat references:** THR-007, THR-014.

### 6.5 Identity and Access Policy
- **Purpose:** wer welche Identitäts-/Zugriffsrechte hat. **Owner:** Security/Governance. **Scope:** Identity/Access. **Subjects:** human & machine principals. **Resources:** Identitäts-/RBAC-Objekte. **Actions:** Rollen-/Scope-/Membership-Änderungen.
- **Inputs:** Identity-Klasse, Rolle, Scope, Lifecycle-State. **Decision authority:** Security-Governance. **Conflict owner:** Platform. **Offline applicability:** teilweise. **Audit requirement:** ja. **Threat references:** THR-001, THR-002, THR-038.

### 6.6 Approval Policy
- **Purpose:** legt fest, wann eine Handlung eine gesonderte Approval-Entscheidung erfordert und mit welcher Autorität. **Owner:** Governance. **Scope:** approval-pflichtige Aktionsklassen. **Subjects:** Antragsteller & Approver. **Resources:** betroffene Ressourcen. **Actions:** privilegierte/risikoreiche Aktionen.
- **Inputs:** Risiko/Impact, Action-Klasse, Separation-of-Duties-Anforderung. **Decision authority:** Governance. **Conflict owner:** Platform. **Offline applicability:** siehe Companion 2 §17. **Audit requirement:** ja. **Threat references:** THR-003, THR-004.

### 6.7 Execution Policy
- **Purpose:** unter welchen Bedingungen eine autorisierte Ausführung überhaupt erfolgen darf (Guards). **Owner:** Security-Governance. **Scope:** Execution-Pfade. **Subjects:** ausführende Principals/Adapter/Agents. **Resources:** Ziel-Ressourcen. **Actions:** Execution-Klassen.
- **Inputs:** Authorization-State, Plan-Version, Credential-State, Effective-State. **Decision authority:** Execution-Governance. **Conflict owner:** Platform. **Offline applicability:** ja. **Audit requirement:** ja. **Threat references:** THR-005, THR-006, THR-026, THR-034.

### 6.8 Deployment Policy
- **Purpose:** Regeln für Deployment-/Änderungsaktionen. **Owner:** Deployment-Governance. **Scope:** Deployment-Aktionen. **Subjects:** Deployer/Automation. **Resources:** Deployment-Targets/Artefakte. **Actions:** deploy/rollback/update.
- **Inputs:** Artefakt-Provenance, Environment, Approval-State. **Decision authority:** Deployment-Governance. **Conflict owner:** Platform. **Offline applicability:** ja. **Audit requirement:** ja. **Threat references:** THR-021, THR-022, THR-024.

### 6.9 Integration Policy
- **Purpose:** Regeln für Adapter/Integrationen (insb. read≠write). **Owner:** Integration-Governance. **Scope:** Integrationen. **Subjects:** Adapter/Integrationen. **Resources:** Managed Resources. **Actions:** read/write über Integration.
- **Inputs:** Adapter-Trust, Read/Write-Klassifikation, Scope. **Decision authority:** Integration-Governance. **Conflict owner:** Platform. **Offline applicability:** teilweise. **Audit requirement:** ja. **Threat references:** THR-007, THR-010, THR-011.

### 6.10 Offline Policy
- **Purpose:** Regeln, die in offline/isolierten Umgebungen gelten (Aktivierung, Provenance, Revocation-Verteilung). **Owner:** Environment-/Security-Governance. **Scope:** offline/air-gapped Betrieb. **Subjects:** lokale Principals. **Resources:** lokale Ressourcen. **Actions:** lokal aktivierte Aktionen.
- **Inputs:** Offline-State, Provenance, Integrity, Revocation-Distribution-State. **Decision authority:** Environment-Governance. **Conflict owner:** Platform bei Reconnection. **Offline applicability:** definitionsgemäß. **Audit requirement:** ja, mit Audit-Continuity. **Threat references:** THR-024, THR-027.

### 6.11 Evidence and Audit Policy
- **Purpose:** was auditiert/als Evidenz erfasst wird und wie Evidenz geschützt bleibt. **Owner:** Security-Governance. **Scope:** Audit/Evidence. **Subjects:** alle. **Resources:** Audit-/Evidence-Objekte. **Actions:** record/export.
- **Inputs:** Ereignistyp, Datenklasse, Redaction-Anforderung. **Decision authority:** Security-Governance. **Conflict owner:** Platform. **Offline applicability:** ja. **Audit requirement:** ja (Selbstbezug). **Threat references:** THR-016, THR-017, THR-018.

### 6.12 Emergency-Access Policy
- **Purpose:** außergewöhnlicher Zugriff (Break Glass) unter strikten Bedingungen. **Owner:** Security-Governance / Human Maintainer. **Scope:** Notfall. **Subjects:** benannte Notfall-Principals. **Resources:** betroffene Ressourcen. **Actions:** notwendige privilegierte Aktionen.
- **Inputs:** Emergency-State, benannte Identität, Ablauf, Scope. **Decision authority:** Emergency-Governance. **Conflict owner:** Human Maintainer. **Offline applicability:** ja. **Audit requirement:** verschärft, mit Post-Event Review. **Threat references:** THR-037. Die bestehende [Break-Glass-Policy](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md) bleibt autoritativ; dieses Dokument erzeugt kein Parallelmodell.

## 7. Policy Lifecycle

Konzeptionelle Statuswerte (keine implementierten Enums):

```text
draft → review-pending → approved → active → suspended → superseded → withdrawn → expired → conflicted → archived
```

Grundregeln:
- Nur `approved` **und** `active` Policies dürfen für aktuelle normative Entscheidungen verwendet werden.
- `draft` ist **nicht** wirksam.
- `approved` bedeutet **nicht** automatische Aktivierung (`active` ist ein separater Schritt).
- Policy-Versionen bleiben historisch nachvollziehbar; Supersession löscht **keine** frühere Decision-Evidenz.
- Eine laufende Authorization bleibt an die **tatsächlich ausgewertete** Policy-Version gebunden.
- `withdrawn`/`expired`/`suspended` Policies dürfen **keine** neue Autorisierung begründen.
- `conflicted` markiert nicht aufgelöste Präzedenz (siehe §12) und begründet keinen Trust.

## 8. Policy Scope

Ein Policy-Scope bindet mindestens: Subject-Klasse · Resource/Resource-Klasse · Action/Action-Klasse · Workspace · Environment · Policy-Version. Grundregeln:
- Keine Policy gilt breiter als ihr deklarierter Scope.
- Eine Workspace-Policy erzeugt **keine** globale Wirkung.
- Scope-Zuordnung, die unklar ist, führt zu `indeterminate` (§10), nicht zu `permit`.
- Überlappende Scopes werden über Präzedenz (§12) geregelt, nicht durch stillen Vorrang.

## 9. Evaluation Inputs

Für eine Auswertung werden mindestens berücksichtigt:

```text
principal identity · principal class · human or machine identity · workspace · environment ·
resource identity · resource class · action · action privilege level · requested scope ·
current lifecycle state · credential state · desired-state reference · drift or remediation reference ·
policy version · approval state · exception state · break-glass state · freshness state ·
trust state · conflict state · offline state · time or sequence context
```

**Grundregel:** Ein fehlender oder unklarer sicherheitsrelevanter Input darf **nicht** still positiv interpretiert werden. Fehlt ein für die Entscheidung notwendiger Input, ist das Outcome `indeterminate` (bei privilegierten Aktionen fail-closed), nicht `permit`. Herkunft/Freshness/Trust von Inputs stammen aus dem [Source-of-Truth-Modell](../architecture/SOURCE_OF_TRUTH_AND_STATE_AUTHORITY_MODEL.md) und dem [State-Modell](../architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md); `evidence capability ≠ evidence available ≠ requirement satisfied`.

## 10. Decision Outcomes

```text
permit · deny · approval-required · not-applicable · indeterminate · conflicted
```

- **`permit`** — Die Policy verbietet die angefragte Handlung unter den geprüften Bedingungen nicht.
  ```text
  permit ≠ approved ≠ execution-authorized
  ```
- **`deny`** — Die Handlung ist im geprüften Scope nicht erlaubt.
- **`approval-required`** — Policy verlangt eine gesonderte Approval-Entscheidung (Companion 2), bevor überhaupt eine Authorization entstehen kann.
- **`not-applicable`** — Die Policy ist für den angefragten Scope/Kontext nicht zuständig.
- **`indeterminate`** — Notwendige Inputs, Autorität oder Auswertung sind nicht ausreichend bestimmbar.
- **`conflicted`** — Mehrere anwendbare Policy Decisions stehen in nicht aufgelöstem Konflikt (§12).

**Grundregel:**
```text
indeterminate or conflicted → no privileged execution authorization
```
Es gibt **keine** implizite Default-Permit-Regel (§11).

## 11. Default-Deny Boundary

- Für privilegierte Aktionen gilt: ohne eindeutiges, anwendbares und ausreichend belegtes `permit` **plus** erfüllte Approval- und Authorization-Bedingungen entsteht **keine** Ausführungsberechtigung.
- Das Standard-Outcome für privilegierte Aktionen bei fehlender Klarheit ist `deny` oder `indeterminate`, niemals `permit`.
- `not-applicable` einer einzelnen Policy ist **kein** `permit`; es verschiebt die Entscheidung auf zuständige Policies oder auf Default-Deny.
- Ein `permit` erfüllt nur die Policy-Ebene; Approval (Companion 2) und Execution Authorization (Companion 3) sind eigenständige, zusätzliche Bedingungen.

## 12. Conflicts and Precedence

Konfliktquellen: mehrere Policy-Ebenen; Workspace- vs. Plattformregeln; aktive Exception; Break Glass; Offline-Policy; widersprüchliche Policy-Versionen; unklare Scope-Zuordnung.

Grundregeln:
- Keine universelle „Last-Rule-Wins"-Regel.
- Keine automatische Priorität allein durch neueren Timestamp.
- `deny` darf **nicht** ohne explizite Ausnahme- oder Emergency-Governance überschrieben werden.
- Break Glass entfernt **keine** ursprüngliche Policy.
- Präzedenz ist **pro Policy-Klasse oder Scope** geregelt (z. B. Platform Governance vor Workspace), nicht global-implizit.
- Ungelöste Konflikte bleiben als `conflicted` **sichtbar** und blockieren privilegierte Autorisierung.
- Jede Konfliktauflösung ist auditierbar.

Es wird **keine** konkrete Conflict-Resolution-Engine ausgewählt.

## 13. Exceptions

- Eine Exception ist eine **explizit erfasste, begründete, scope-gebundene und zeit-/review-begrenzte** Abweichung von einer Policy.
- Eine Exception ist **kein** stiller Policy-Bypass (`exception ≠ silent policy bypass`).
- Exceptions haben einen zurechenbaren menschlichen Verantwortlichen, einen Ablauf und eine Review-Pflicht.
- Eine Machine Identity darf **nicht** alleinige Risikoakzeptanz für eine Exception tragen.
- Exceptions werden auditiert und propagieren **nicht** automatisch auf andere Ressourcen oder Scopes.
- Die Exception-Semantik ist konsistent mit dem [Drift-/Exception-Modell](../architecture/DRIFT_DETECTION_AND_CONVERGENCE_MODEL.md) (dortige Divergence Exceptions).

## 14. Break-Glass Relationship

Break Glass kann eine außergewöhnliche Auswertungs-/Autorisierungsgrundlage liefern, ersetzt aber nicht Identitätszurechenbarkeit, Scope-Bindung, Grund, Ablauf, Audit, Revocation oder Post-Event Review.

```text
break glass ≠ permanent permit
break glass ≠ anonymous execution
break glass ≠ removal of original policy history
break glass ≠ unlimited scope
```

Die bestehende [Break-Glass-Policy](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md) bleibt authoritative; dieses Dokument referenziert sie und erzeugt kein Parallelmodell.

## 15. Offline Evaluation

- In offline/isolierten Umgebungen erfolgt die Auswertung gegen die **lokal gültige, provenance-geprüfte** Policy-Version.
- Eine offline nicht verfügbare Revocation-Information darf **nicht** still als „keine Revocation" interpretiert werden; unklare Autorität führt zu `indeterminate`/fail-closed.
- Offline getroffene Auswertungen sind **audit-continuity-pflichtig** und werden nach Reconnection reconciled (siehe [Offline Reconciliation Policy](OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md)).
- Es wird **keine** Offline-Auswertung als implementiert oder für beliebige Air-Gap-Stufen geeignet behauptet; keine Signing-/Trust-Anchor-Technologie wird ausgewählt.

## 16. Audit and Evidence

Mindestens erfasst werden: Policy-Identität und -Version · Evaluation Inputs (referenziert/klassifiziert) · Evaluation Outcome · Conflict-Status · ob `approval-required` ausgelöst wurde · Zeit-/Sequenzkontext · zurechenbarer Principal.

Trennung:
```text
authorization evidence ≠ execution evidence
execution evidence     ≠ success evidence
success evidence       ≠ verification evidence
verification evidence  ≠ compliance
```
Audit-Records sind **keine** Ausführungsauslöser (`audit evidence must not grant execution authority`) und werden nicht durch spätere Auswertung umgeschrieben.

## 17. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Policy permit must not imply execution authorization.
2. Policy evaluation must not itself perform or trigger an action.
3. A missing or unclear security-relevant input must not be interpreted as permit.
4. There is no implicit default-permit for privileged actions.
5. `indeterminate` or `conflicted` must block privileged execution authorization.
6. Policy conflicts must remain visible until explicitly and auditable resolved.
7. `deny` must not be overridden without explicit exception or emergency governance.
8. Only `approved` and `active` policy versions may drive current normative decisions.
9. A running authorization stays bound to the policy version actually evaluated.
10. Break-glass must not remove original policy history.
11. Offline evaluation must not treat missing revocation as absence of revocation.
12. Audit evidence must not grant execution authority.

Keine dieser Invarianten ist als implementierte Kontrolle dargestellt.

## 18. Threat References

Bestehende Szenarien aus dem [Threat Scenario Register](THREAT_SCENARIO_REGISTER.md) (nur reale IDs; keine Duplikation, kein Parallelregister):
THR-001, THR-002, THR-003, THR-004, THR-005, THR-006, THR-007, THR-010, THR-011, THR-014, THR-016, THR-017, THR-018, THR-024, THR-026, THR-027, THR-034, THR-037, THR-038.

## 19. Technology Boundary

Nicht ausgewählt und nicht implementiert: Policy Engine, Policy-Sprache/-Grammatik, Autorisierungsartefakt/Token, kryptografische Signatur/Trust Anchor, API-/DB-/Event-Schema, Workflow-/Queue-/Execution-Runtime, automatisierte Approval-Entscheidung. Dieses Dokument ist rein konzeptionell.

## 20. Compatibility

Konsistent mit: [Logical Module Architecture](../architecture/COREOPS_LOGICAL_MODULE_ARCHITECTURE.md) (Policy/Control/Execution getrennt), [Module Catalog](../architecture/COREOPS_MODULE_CATALOG.md) (Datenownership), [Human Identity/RBAC](HUMAN_IDENTITY_AND_ACCESS_GOVERNANCE.md), [Machine Identity](MACHINE_IDENTITY_AND_PRINCIPAL_GOVERNANCE.md), [State-Modell](../architecture/OBSERVED_DESIRED_AND_EFFECTIVE_STATE_MODEL.md), [Safe Remediation](SAFE_REMEDIATION_AND_STATE_CHANGE_POLICY.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Ergänzt DEC-P-05 (KI beratend) und DEC-G-07 (Control Plane nicht umgehbar).

## 21. Open Questions

- Präzise Präzedenzordnung zwischen Platform-, Environment-, Workspace- und Resource-Policy (ADR-Kandidat, später).
- Formaler Umgang mit `conflicted` in gemischten Offline-/Online-Auswertungen.
- Granularität der Evaluation-Input-Erfassung im Audit unter Datenminimierung.

## 22. Next Decision

Approval- und Authorization-Lifecycle (Companion 2) sowie Execution Authorization und Guards (Companion 3) bauen auf diesem Decision-Modell auf. Eine spätere ADR-Runde (frühestens CO-WP-029/Readiness) kann Policy-Engine- und Präzedenzfragen adressieren. In diesem WP wird keine Technologie ausgewählt.
