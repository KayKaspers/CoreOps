# CoreOps – Secrets, Configuration Vault and Custody Governance

> Document Status: Implemented, pending Nova review
> Policy Status: Foundation secret, sensitive-configuration and vault custody governance
> Implementation Status: Not implemented
> Vault Technology: Not selected
> Secret Store: Not selected
> Encryption Technology: Not selected
> Raw Secret Storage: Not decided
> Validation Status: Not performed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-024 (docs-only / security, configuration, secret-lifecycle and key-custody governance foundation)

## 1. Status

Technologieunabhängige Governance für **Secrets, sensitive configuration, Vault-Autoritätsgrenze, Secret-Lebenszyklus, Retrieval/Use, Disclosure, Profile und Audit**. Companion zu [KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md](KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md) und [CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md](../architecture/CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md). Baut auf der Modulgrenze [MOD-SEC-001 Credential and Secret Governance Boundary](../architecture/COREOPS_MODULE_CATALOG.md) und [Offline Credential and Rotation Governance (CO-WP-010)](OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md) auf; erweitert diese ohne Parallelmodell. Kein Vault, kein Secret Store, keine Encryption implementiert.

## 2. Purpose

Secrets sind die höchstprivilegierte Datenklasse. Diese Governance legt fest, warum `secret ≠ ordinary configuration`, `secret reference ≠ secret value`, `credential ≠ identity`, `credential possession ≠ authorization`, `retrieval permitted ≠ use permitted`, `use permitted ≠ export permitted` und `vault availability ≠ secret trust`. Sie definiert Begriffe, Zustände, Autoritäten, Bindungen und Fail-Closed-Regeln — ohne Vault-, KMS- oder Krypto-Technologie auszuwählen.

## 3. Scope

Concepts · Secret/Configuration Classification · Authority Model · Vault Governance · Secret Lifecycle · Secret Identity/Instances · Retrieval/Distribution/Injection/Use · Disclosure/Export · Workspace/Environment/Target Isolation · CorePack/Deployment Boundary · Audit/Evidence · Profiles · Failure/Fail-Closed. Key Custody, Rotation, Revocation, Recovery und Destruction im [Companion Key Policy](KEY_MATERIAL_ROTATION_REVOCATION_AND_RECOVERY_POLICY.md); Configuration Source of Truth und Secret References im [Companion Configuration Model](../architecture/CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md).

## 4. Non-Goals

- Kein Vault-Produkt/-Dienst, Secret Store, Datenbank, KMS, HSM, TPM, PKI, Trust Anchor, CA, Key-/Secret-/Token-Format, Hash-/Encryption-/Signatur-/KDF-Verfahren, RNG, Envelope Encryption, Secret Sharing, Threshold/Quorum, Injection/Sidecar/Mount, Config-Format, Policy-/Rotation-Engine, Scanner, Redaction-Engine, Backup-Encryption, Exportformat, Transfer-/Sync-Dienst, Cloud/IdP, Hardware Root of Trust.
- Keine Behauptung implementierter/validierter/zertifizierter Secret-Kontrollen; kein Runtime-Code; keine Compliance-/Einsatzreife-Behauptung; keine Rohsecret-Speicherentscheidung.

## 5. Concepts

`ordinary configuration` · `sensitive configuration` · `secret` · `secret reference` · `secret metadata` · `credential` · `authentication material` · `authorization material` · `key` · `key material` · `public key material` · `private/restricted key material` · `certificate reference` · `logical secret identity` · `secret version` · `secret value revision` · `secret value instance` · `credential instance` · `lease` · `retrieval instance` · `distribution instance` · `injection instance` · `use instance` · `rotation event` · `revocation event` · `suspension` · `reinstatement` · `destruction` · `vault` · `vault authority` · `custodian` · `owner` · `issuer` · `consumer` · `secret broker` · `rotation authority` · `revocation authority` · `recovery authority` · `export authority` · `break-glass authority` · `configuration source of truth` · `configuration snapshot` · `configuration overlay` · `runtime-effective configuration` · `configuration drift` · `configuration reconciliation`.

```text
vault = governte logische Schutz- und Autoritätsgrenze
vault ≠ ausgewähltes Produkt
```

## 6. Secret and Configuration Classification

Technologieunabhängige Klassen: `public configuration` · `internal configuration` · `sensitive configuration` · `secret configuration` · `authentication credential` · `machine credential` · `human credential reference` · `workload credential` · `API credential` · `temporary credential` · `recovery material` · `root/bootstrap material` · `signing/verification key reference` · `encryption/protection key reference` · `unknown-sensitive material`.

Je Klasse dokumentiert: Zweck · Owner · Custodian · erlaubte Verbraucher · erlaubter Scope · erlaubte Speicherung · Exportgrenze · Auditgrenze · Rotationserwartung · Revocation-Verhalten · Offline-Verhalten · Recovery-Grenze · Known Limitations. `unknown-sensitive material` wird konservativ als Secret behandelt (fail-closed). **Keine realistischen Secrets/Credentials als Beispiele.**

## 7. Authority Model

Getrennt: configuration policy authority · configuration approval authority · secret policy authority · secret owner · secret custodian · secret issuer · secret registration authority · secret retrieval authority · secret use authority · secret distribution authority · secret export authority · rotation authority · revocation authority · suspension authority · reinstatement authority · recovery authority · destruction authority · break-glass authority · audit authority.

```text
owner              ≠ custodian
custodian          ≠ unrestricted user
retrieval authority ≠ use authority
use authority      ≠ export authority
rotation authority ≠ reinstatement authority
backup operator    ≠ recovery authority
local administrator ≠ unrestricted secret authority
machine principal  ≠ human approval authority
break-glass operator ≠ permanent policy authority
```

Delegation bindet an: accountable human owner · action · secret/key identity · target · consumer · workspace · environment · purpose · profile · version/revision · validity period · max use/use count (falls anwendbar) · review trigger · audit reference. **Offline-Betrieb erweitert keine Autorität.** Autorität bindet an [Policy Decision](POLICY_DECISION_AND_EVALUATION_MODEL.md), [Approval Lifecycle](APPROVAL_AND_AUTHORIZATION_LIFECYCLE.md), [Execution Authorization](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) (CO-WP-013) — keine parallele Autorität.

## 8. Vault Governance

Ein Vault ist ausschließlich eine **logische Governance-Grenze**, kein Produkt. Mindestattribute: vault identity · scope · workspace binding · environment binding · target/consumer scope · profile · accountable owner · custodian · policy version · supported secret classes · authority model · import boundary · export boundary · offline boundary · recovery boundary · audit boundary · known limitations.

```text
vault availability   ≠ secret availability for every consumer
vault membership     ≠ use authorization
vault administrator  ≠ secret owner
vault backup         ≠ recovery authorization
vault replication    ≠ trust equivalence
vault import         ≠ secret activation
vault container trust ≠ secret value validity
```

Keine Vault-Technologie ausgewählt.

## 9. Secret Lifecycle

Getrennte Zustände/Ereignisse: `proposed · generated/externally-issued · imported · registered · quarantined · assessed · approved · available · retrievable · distributed · injected · active-for-bounded-use · rotation-pending · rotated · superseded · suspended · revoked · recovery-pending · recovered · destruction-pending · destroyed · expired · outcome-unknown`.

```text
generated  ≠ registered
registered ≠ approved
approved   ≠ retrievable
retrievable ≠ usable
distributed ≠ injected
injected   ≠ successfully consumed
rotated    ≠ every consumer updated
revoked    ≠ removed from every runtime
destroyed  ≠ historical evidence erased
expired    ≠ safely removed
unknown    ≠ safe
```

Unknown Outcomes bleiben sichtbar, blockieren automatische Wiederholung, verlangen Reconciliation und berücksichtigen mögliche Side Effects (§17).

## 10. Secret Identity and Instances

Getrennt: stable logical secret identity · secret version · secret value revision · generated value instance · credential instance · lease instance · retrieval instance · distribution instance · injection instance · use instance · rotation event · revocation event · recovery instance · destruction event.

```text
same logical identity ≠ same value
same version          ≠ same distribution history
same value bytes      ≠ same authorization
same credential       ≠ same use instance
same key material     ≠ same custody state
same reference        ≠ same resolved value
```

Mutable Aliases sind **keine** finale privilegierte Bindung. Deployments/privilegierte Aktionen zeigen auf eine konkrete, bewertete Secret-/Credential-Revision bzw. eine eindeutig auflösbare gebundene Referenz.

## 11. Retrieval, Distribution, Injection and Use

Getrennte Zustände: retrieval request · retrieval authorization · value resolution · distribution · receipt · injection · workload consumption · use completion · cleanup · outcome verification. Erfasst: request identity · human/machine requester · secret identity/revision · consumer · target · purpose · approval · authorization · validity · max use · distribution instance · injection instance · outcome · cleanup state · audit reference.

```text
retrieved         ≠ delivered
delivered         ≠ injected
injected          ≠ consumed
consumed          ≠ action successful
cleanup requested ≠ value absent
```

Keine Injection-Technologie ausgewählt.

## 12. Disclosure and Export

Getrennt: local read · secret retrieval · secret use · diagnostic visibility · support visibility · evidence visibility · secret export · key export · configuration export.

```text
local administrator access ≠ export authority
support access             ≠ secret access
audit access               ≠ secret value access
configuration export       ≠ secret export
activation authority       ≠ credential export authority
```

Minimierung und Zweckbindung gelten; kein Compliance-Claim. Ausgänge enthalten keine Rohsecrets/wiederverwendbaren Credentials (Bezug THR-020, THR-025, THR-040).

## 13. Workspace, Environment and Target Isolation

Secret-/Config-/Key-Bindungen bleiben durch Registration, Reference Resolution, Retrieval, Distribution, Injection, Rotation, Recovery und Offline Use erhalten. Ein Secret aus Workspace A wird nicht still in Workspace B verwendet; Environment Binding nicht durch Rename ersetzbar; geteilte Infrastruktur erzeugt keine gemeinsame Secret Authority; `same secret bytes ≠ same authorization`; ein Credential für Target A ist nicht automatisch für Target B gültig; Cross-Workspace-/Cross-Environment-Use ist explizit governed (Bezug THR-035).

## 14. CorePack and Deployment Boundary

Bindet an [CorePack (CO-WP-023)](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md), [Deployment (CO-WP-021)](../architecture/DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md), [Artifact Trust (CO-WP-022)](ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md):

```text
CorePack trust                  ≠ secret trust
artifact trusted                ≠ secret trusted
deployment approved             ≠ secret use approved
deployment execution authorization ≠ secret export authorization
```

CorePacks/Blueprints enthalten **keine** Raw Secrets/wiederverwendbaren Credentials — nur Secret References/Requirements/Governance-Metadaten; CorePack Activation autorisiert keine Secret Retrieval/Use; Secret Resolution erweitert Target Binding/Workspace Isolation nicht; Evidence-Return-CorePacks exportieren keine Secret Values; Recovery-CorePacks erhalten keine automatische Secret/Key Recovery Authority. Details im [Configuration Model §14–§15](../architecture/CONFIGURATION_SOURCE_OF_TRUTH_AND_SECRET_REFERENCE_MODEL.md).

## 15. Audit and Evidence

Audit/Evidence legen **keine** Secret Values offen. Sichere Metadaten: secret/key identity · version · request identity · requester · consumer · target · purpose · policy decision · approval · authorization · retrieval/distribution/injection/rotation/revocation/recovery/destruction event · outcome · freshness · audit reference.

```text
audit metadata available    ≠ secret value recorded
redacted output             ≠ guaranteed safe disclosure
hash or fingerprint reference ≠ secret value
access logged               ≠ access authorized
missing audit event         ≠ no secret use occurred
```

Records gehören MOD-EVD-001 ([CO-WP-018](AUDIT_INTEGRITY_RETENTION_AND_DISCLOSURE_POLICY.md)); Historie nicht umgeschrieben (Bezug THR-016, THR-017, THR-019). Keine Redaction-/Logging-/Hash-Technologie.

## 16. Profiles

`Standard`, `Hardened`, `Government` sind Governance-/Kontrollprofile mit unterschiedlicher Stärke (Dimensionen: Freigabetiefe · Rollentrennung · Offline-Freshness · Gültigkeitsdauer · erlaubte Exportwege · Break-Glass-Anforderungen · Rotationserwartungen · Recovery-Grenzen · Audit-Tiefe · Unknown-State-Verhalten).

```text
Government profile ≠ government certification
Hardened profile   ≠ proven secure
Standard profile   ≠ security controls optional
```

Keine konkrete technische Ausprägung vorgeschrieben; keine Compliance/Zertifizierung/Einsatzreife behauptet.

## 17. Failure and Unknown State

`secret reference unresolved · ambiguous resolution · secret unavailable · credential expired/revoked · rotation partially completed · old version still in use · distribution/injection/cleanup outcome unknown · offline freshness unknown · policy conflict · target/workspace mismatch · recovery/destruction outcome unknown · audit gap · unauthorized export suspected`.

```text
unknown          ≠ safe
failure          ≠ no side effects
retry            ≠ automatically permitted
cleanup unknown  ≠ value absent
missing evidence ≠ operation did not occur
```

## 18. Fail-Closed Rules

Keine privilegierte Secret-Retrieval/-Use bei: unresolved/ambiguous/conflicted Secret Reference · fehlender Approval/Authorization für Ziel/Consumer/Scope · abgelaufener/widerrufener/suspendierter Secret-Version · Target-/Workspace-/Environment-Mismatch · unbekannter Offline-Freshness ohne Exception · fehlendem Audit-Start · Alias-only-Binding für privilegierte Aktion · vermutetem unautorisiertem Export. Kein stiller Fallback auf ein anderes Secret.

## 19. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. Raw secrets do not appear in ordinary configuration, logs, events, CorePacks, deployment blueprints or ordinary evidence packages.
2. Secret reference and secret value remain separate.
3. Secret, credential and key identities remain separate from versions and instances.
4. Credential possession does not imply authorization; authentication material ≠ authorization material.
5. Retrieval, distribution, injection and use remain separate states.
6. Owner, custodian, issuer, consumer and the distinct authorities remain separate; retrieval ≠ use ≠ export.
7. Vault is a logical governance boundary; vault availability/membership does not imply secret trust or use authorization.
8. Offline operation does not expand authority; offline secret authorizations are bounded and non-reusable.
9. Unknown retrieval/injection/cleanup outcomes remain visible and block unsafe retry.
10. Audit and evidence contain no secret values; redaction is not asserted as guaranteed safety.
11. CorePack trust does not imply secret trust; deployment authorization does not imply secret-use authorization.
12. Secret and configuration bindings to workspace, environment and target are preserved and not replaced by rename.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 20. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](THREAT_SCENARIO_REGISTER.md) (lokal verifiziert, alle ≤040): Secret Exposure THR-019/THR-020/THR-040; offline export sensitive THR-025; offline import THR-024; replay THR-026; false time THR-027; reuse expired approval THR-004; stolen admin identity THR-001; privilege escalation THR-002; audit deletion/manipulation THR-016/THR-017; tenant boundary THR-035; insider THR-037; automation client THR-038; managed resource vs CoreOps THR-034. Keine erfundenen IDs; kein Parallelregister.

## 21. Technology Boundary

Nicht ausgewählt: Vault-Produkt/-Dienst · Secret Store · Datenbank · Config Store · KMS · HSM · TPM · PKI · CA · Trust Anchor · Key-/Secret-/Token-Format · Hash-/Encryption-/Signatur-/KDF-Verfahren · RNG · Envelope Encryption · Secret Sharing · Threshold/Quorum · Injection/Sidecar/Mount · Config-Format · Policy-/Rotation-Engine · Scanner · Redaction-Engine · Backup-Encryption · Export-/Transfer-/Sync-Technologie · Cloud/IdP · Hardware Root of Trust. Alle bleiben `deferred`; Rohsecret-Speicherung nicht entschieden.

## 22. Compatibility

Konsistent mit MOD-SEC-001, MOD-STA-001, MOD-EVD-001, [Machine Credential Lifecycle (CO-WP-010)](OFFLINE_CREDENTIAL_AND_ROTATION_GOVERNANCE.md), [Break-Glass (CO-WP-009)](BREAK_GLASS_AND_EMERGENCY_ACCESS_POLICY.md), [Policy/Approval/Execution (CO-WP-013)](EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md), [Artifact Trust (CO-WP-022)](ARTIFACT_TRUST_QUARANTINE_AND_REVOCATION_POLICY.md), [CorePack/Offline (CO-WP-023)](../architecture/COREPACK_IDENTITY_CONTENT_AND_LIFECYCLE_MODEL.md). Additiv; keine bestehende Foundation-Invariante geschwächt; kein Parallelmodell.

## 23. Open Questions

- Ob CoreOps Rohsecretmaterial speichert (spätere Sicherheitsarchitektur-Entscheidung, offen seit CO-WP-010).
- Empfohlene Default-Rotationsintervalle/Overlap-Fenster (Companion Key Policy, deferred).
- Zuverlässige Offline-Revocation-Distribution/-Reconciliation (Companion, deferred).

## 24. Next Decision

Nächstes lokal registriertes Work Package gemäß [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md): `CO-WP-025 – Data Classification, Retention and Redaction`. Zuerst Nova Review von CO-WP-024, danach Human-Maintainer-Commit. Vault-/KMS-/Krypto-Technologie bleibt separate spätere ADR-gestützte Arbeit.
