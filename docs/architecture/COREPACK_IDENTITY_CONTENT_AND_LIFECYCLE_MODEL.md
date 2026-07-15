# CoreOps – CorePack Identity, Content and Lifecycle Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation CorePack identity, content and lifecycle model
> Implementation Status: Not implemented
> Pack Format: Not selected
> Manifest Format: Not selected
> Installer Technology: Not selected
> Update Technology: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-023 (docs-only / restricted-operation, offline-distribution and CorePack governance foundation)

## 1. Status

Technologieunabhängiges Modell für **CorePack** als governte CoreOps-Distributionseinheit: Identität, Version/Revision/Instanzen, Manifest, Content Population/Resolution, Assembly, Provenance/Integrity/Validation/Trust/Compatibility, Target Binding, Lifecycle, Transfer/Receipt/Import/Quarantine/Activation, Update/Delta, Revocation/Reinstatement, Rollback/Recovery, Evidence Return. Companion zu [RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md](RESTRICTED_ISOLATED_AND_AIR_GAPPED_OPERATION_MODEL.md) und [OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md](../security/OFFLINE_TRUST_ACTIVATION_REVOCATION_AND_TRANSFER_POLICY.md). Kein Paket-/Manifestformat, kein Installer, keine Update-Technologie.

## 2. Purpose

CorePack ist die Einheit, über die CoreOps-bezogene Inhalte offline zu einer restricted/isolated/air-gapped Umgebung transferiert, importiert, bewertet und aktiviert werden. Das Modell legt fest, warum `CorePack transferred ≠ imported`, `imported ≠ trusted`, `trusted ≠ approved for activation`, `activated ≠ deployment authorised` und warum CorePack-Contents Trust/Compatibility **nicht** automatisch vom Container erben.

## 3. Scope

CorePack Concepts/Boundary/Classes · Identity · Version/Revision/Instances · Manifest · Content Population/Resolution · Assembly · Provenance/Integrity/Validation/Trust/Compatibility · Target Binding · Lifecycle · Transfer/Receipt/Import/Quarantine · Activation · Partial Activation · Updates/Deltas · Revocation/Reinstatement · Rollback/Recovery · Evidence Return · Audit · Security Invariants.

## 4. Non-Goals

- Kein konkretes Archiv-/Paketformat, kein Manifest-/Serialisierungsformat, kein Installer, keine Update-Engine, keine Transfer-/Removable-Media-Technologie, keine Signing-/Hash-Technologie.
- Keine tatsächliche CorePack-Erzeugung, kein Artifact-Resolve/-Download, keine Aktivierung; kein Runtime-Code.
- Keine Air-Gap-/Security-/Compliance-Eignungsbehauptung.

## 5. CorePack Concepts

`CorePack` · `CorePack identity` · `CorePack version` · `CorePack revision` · `assembly instance` · `distribution instance` · `transfer instance` · `receipt instance` · `import instance` · `activation instance` · `rollback instance` · `result-package instance` · `manifest` · `content set` · `content population` · `content resolution` · `trust snapshot` · `revocation snapshot` · `transfer` · `receipt` · `import` · `quarantine` · `assessment` · `approval` · `activation` · `update` · `delta` · `rollback` · `recovery` · `evidence-return package`.

## 6. CorePack Boundary

CorePack ist eine versionierte, governte Distributionseinheit — **keine** andere Foundation-Einheit:

```text
CorePack ≠ Domain Pack
CorePack ≠ single artifact
CorePack ≠ deployment blueprint
CorePack ≠ backup
CorePack ≠ evidence package
```

Ein CorePack kann konzeptionell enthalten: CoreOps application/service artifacts · agent/adapter artifacts · Domain Packs · deployment blueprints · policy packages · schema/migration references · configuration templates · documentation/runbooks · SBOM references · provenance/integrity metadata · compatibility information · revocation snapshot · trust snapshot · known limitations. Ein CorePack enthält **keine** Raw Secrets oder wiederverwendbaren realen Credentials (Bezug THR-019, THR-020).

```text
CorePack contents available ≠ contents individually trusted
                            ≠ contents compatible with target
                            ≠ contents authorised for activation
```

Enthaltene Domain Packs, Artifacts und Blueprints behalten ihre eigenen Governance-Grenzen ([CO-WP-015](DOMAIN_PACK_GOVERNANCE_MODEL.md), [CO-WP-022](ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md), [CO-WP-021](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md)); der Container erweitert sie nicht.

## 7. CorePack Classes

| Klasse | Kurzcharakter |
| ------ | ------------- |
| `full installation pack` | vollständige Erstinstallation |
| `bounded update pack` | begrenztes Update definierter Contents |
| `delta pack` | Differenz gegen exakt gebundene Ausgangsbasis |
| `policy update pack` | ausschließlich Policy-Inhalte |
| `revocation update pack` | ausschließlich Revocation-Snapshot |
| `Domain-Pack update pack` | Domain-Pack-Contents |
| `recovery pack` | Wiederherstellungsinhalte |
| `diagnostic pack` | Diagnose-/Evidence-Sammlung |
| `evidence-return pack` | ausgehende Evidence |
| `result and reconciliation pack` | Ergebnis-/Reconciliation-Daten |
| `unknown pack` | nicht klassifiziert → konservativ, Quarantine |

Je Klasse dokumentiert: Purpose · Permitted contents · Target constraints · Authority expectation · Validation expectation · Activation behavior · Rollback relevance · Offline relevance · Known risks. Ein `delta pack` wird **nie** ohne eindeutig gebundene Ausgangsbasis bewertet oder aktiviert (§27). Ein `unknown pack` bleibt bis zur Klassifikation in Quarantine.

## 8. CorePack Identity

Mindestfelder: stable CorePack identity · canonical name · owner · maintainer · purpose · pack class · version · revision · lifecycle state · assembly identity · content-manifest reference · target profiles · target environment constraints · workspace/scope constraints · compatibility state · trust state · validation state · revocation state · validity boundary · known limitations · audit reference.

```text
file name       ≠ CorePack identity
archive name    ≠ CorePack identity
CorePack version ≠ CoreOps product version
CorePack revision ≠ distribution instance
```

CorePack-IDs werden **nicht** still wiederverwendet; retired IDs bleiben mit historischer Evidenz erhalten.

## 9. Version, Revision and Instances

Getrennte Dimensionen: `CorePack version` · `CorePack revision` · `assembly instance` · `distribution instance` · `transfer instance` · `receipt instance` · `import instance` · `activation instance` · `rollback instance` · `result-package instance`.

```text
same CorePack version   ≠ same revision
same CorePack revision  ≠ same transfer history
same transferred bytes  ≠ same target binding or activation authority
```

## 10. Manifest

Ein konzeptionelles Manifest benötigt: CorePack identity/version/revision · assembly identity · content population · content artifact identities/revisions · Domain-Pack identities/versions · blueprint identities/versions · policy/schema references · dependency relationships · required CoreOps version/compatibility scope · target profiles · target environment binding · workspace/scope constraints · provenance state · integrity state · validation state · trust state · SBOM references · revocation snapshot reference · known exclusions · known limitations · validity boundary · audit reference. Keine Manifest-/Serialisierungstechnologie ausgewählt.

## 11. Content Population

Die Content Population ist die aufgelöste Menge konkreter Inhalte des CorePacks — jeder Eintrag mit Identität, Klasse, gebundener Revision und Herkunft. `content listed in manifest ≠ content present`; `content present ≠ content trusted` (Companion 3 §7). Eine Population kann bekannte Exclusions und Limitations tragen; fehlende Angabe beweist keine Vollständigkeit.

## 12. Content Resolution

Content Resolution bindet jeden Eintrag an eine konkrete Artifact-/Pack-/Blueprint-**Revision** (nicht an einen veränderlichen Alias). Übernommen aus [Artifact Resolution (CO-WP-022)](ARTIFACT_IDENTITY_PROVENANCE_AND_SBOM_MODEL.md): `mutable alias ≠ final binding`. Mutable Aliases dürfen nicht Bestandteil der finalen privilegierten Bindung sein (§13). Unaufgelöste oder mehrdeutige Contents blockieren Assembly-Abschluss.

## 13. Assembly

Assembly benötigt: assembly identity · CorePack identity/revision · assembly owner · content selection decision · content-resolution results · artifact/dependency bindings · source/builder contexts · manifest generation context · trust/revocation snapshot · validation plan · target profiles · time boundary · audit reference.

```text
assembly completed ≠ CorePack validated ≠ CorePack approved for transfer
```

Mutable Artifact Aliases sind **nicht** Bestandteil der finalen privilegierten Bindung.

## 14. Provenance

Getrennt: `content provenance` · `CorePack assembly provenance` · `transfer provenance`. `transfer provenance available ≠ source provenance complete`. Provenance-Lücken bleiben sichtbar; Provenance ≠ Integrity ≠ Trust (Bezug THR-024).

## 15. Integrity

Getrennt: `manifest integrity` · `content integrity`. `CorePack integrity verified ≠ every content item trusted`; `integrity verified ≠ provenance verified ≠ trusted`. Keine Hash-/Signing-Technologie ausgewählt.

## 16. Validation

Validation prüft Manifest-Konsistenz, Content-Resolutions, Compatibility-Scope und Vollständigkeit gegen die Erwartung. `validated ≠ valid for every target`; `validation not performed ≠ invalid`. Validation ist getrennt von Trust und Approval.

## 17. Trust Assessment

Trust ist eine bewusste Bewertung, getrennt von Provenance/Integrity/Validation/Compatibility/Support. `trusted CorePack ≠ activation authorised`. Trust bleibt use-, target-, profile-, version-, revision- und zeitgebunden (Companion 3 §6–§7). Contents erben Trust **nicht** vom Container.

## 18. Compatibility

Getrennt von Trust: Compatibility bindet an CoreOps-Version/Compatibility-Scope, Target Profile und Environment. `compatible ≠ trusted`; `unknown compatibility ≠ compatible`. Unbekannte/konflikthafte Compatibility blockiert automatische privilegierte Activation.

## 19. Target Binding

Vor Transfer bzw. Activation benötigt ein CorePack: target environment identity · workspace/scope · operational connectivity class · CoreOps version/compatibility context · deployment profile · hardware/resource constraints (falls zutreffend) · local authority model · policy version · trust snapshot · revocation snapshot · time/freshness assumptions · known limitations.

```text
CorePack valid for one target ≠ valid for another isolated environment
```

Target Binding bleibt durch Transfer, Import und Activation erhalten und wird nicht durch lokalen Rename ersetzt.

## 20. Lifecycle

`proposed · assembled · validated · trust-assessed · approved-for-transfer · transferred · received · quarantined · assessed · approved-for-activation · activated · partially-activated · superseded · withdrawn · revoked · reinstated · rolled-back`. Zustandsübergänge sind explizit und auditiert; `assembled ≠ approved`, `received ≠ activated`, `withdrawn ≠ revoked`, `reinstated ≠ prior trust automatically restored`.

## 21. Transfer

Transfer erfasst: transfer identity · source/destination environment · transfer owner · CorePack identity/revision · distribution instance · transfer boundary · handling steps · inspection/quarantine expectation · integrity state · target binding · receipt confirmation · known interruptions · audit reference. `handling history ≠ legal chain of custody`. Keine Transfer-/Removable-Media-/Kuriertechnologie ausgewählt (Bezug THR-022, THR-039).

## 22. Receipt

Receipt bestätigt den Empfang einer Distribution Instance im Zielumfeld. Receipt ist ein eigener Zustand vor Import: `received ≠ imported ≠ trusted ≠ activated`. `recent receipt ≠ recent assembly` (§ Freshness, Companion 1 §12).

## 23. Import

Import überführt ein empfangenes CorePack in die kontrollierte Bewertung des Zielumfelds (Quarantine → Assessment). Import erweitert Scope nicht und ersetzt Target Binding nicht. `import failed ≠ no content entered target environment` (fail-closed, Bezug THR-024).

## 24. Quarantine

Quarantine gilt mindestens bei: unknown CorePack identity · unknown revision · target mismatch · manifest missing/conflicted · provenance unknown · integrity failure/unknown · revocation snapshot too stale · unsupported compatibility · unexpected contents · partial transfer · duplicate/replayed transfer (Bezug THR-026).

```text
quarantine release ≠ activation authorization
```

## 25. Activation

Activation benötigt: activation identity · CorePack identity/revision · import instance · target environment · workspace/scope · effective content set · policy decision · approval · activation authority · execution authorization where deployment occurs · validity boundary · preconditions · rollback/recovery plan · verification plan · audit-start capability · audit reference.

```text
activation approved   ≠ every contained deployment authorised
CorePack activated    ≠ every content item installed or executed
activation completed  ≠ desired state verified
```

Deployment innerhalb einer Activation bindet an [Execution Authorization (CO-WP-013)](../security/EXECUTION_AUTHORIZATION_AND_GUARD_POLICY.md) und [Deployment Control Plane (CO-WP-021)](DEPLOYMENT_CONTROL_PLANE_AND_EXECUTION_MODEL.md).

## 26. Partial Activation

Per-Content-Ergebnis: `content activated · content skipped · content rejected · content failed · content outcome unknown · content rollback-pending · content recovery-pending`.

```text
partial activation      ≠ complete success
unknown activation outcome → no automatic retry → reconciliation required
```

Per-Content- und Per-Target-Historie bleiben erhalten (Bezug THR-031, THR-029).

## 27. Updates and Deltas

Ein Update benötigt: base CorePack identity/revision · target CorePack identity/revision · changed contents · removed contents · dependency changes · policy/schema changes · revocation changes · compatibility impact · migration requirement · rollback feasibility · validation evidence.

```text
delta available   ≠ target baseline confirmed
delta applied     ≠ resulting CorePack verified
same base version ≠ same base revision
```

Unbekannte oder konflikthafte Ausgangsbasis blockiert automatische Delta-Aktivierung.

## 28. Revocation

Getrennt: `content-artifact revocation · CorePack revocation · Domain-Pack revocation · policy withdrawal · activation prohibition · distribution prohibition · local emergency suspension`.

```text
one contained artifact revoked ≠ CorePack automatically safe
                               ≠ CorePack automatically wholly revoked
```

Eine Impact-Bewertung bestimmt betroffene Contents, Targets, Activations und Existing Deployments (Companion 3 §17). `no local revocation entry ≠ not revoked centrally`.

## 29. Reinstatement

Reinstatement benötigt: prior suspension/revocation · authority · reason · new evidence · updated CorePack/content revision · updated trust/compatibility · target scope · validity boundary · review trigger · audit reference.

```text
reinstated ≠ historical revocation erased ≠ activation authorised
```

## 30. Rollback and Recovery

Rollback benötigt: prior CorePack identity/revision · prior content population · current revocation/trust state · target compatibility · schema/migration compatibility · rollback authority · recovery plan · verification · remaining differences · audit reference.

```text
previously activated ≠ currently safe rollback target
```

Ein früher aktiviertes CorePack kann inzwischen revoked/inkompatibel sein; Rollback erfordert aktuelle Trust-/Revocation-/Compatibility-Bewertung. Recovery Packs erhalten keine automatische globale Recovery Authority (Bezug THR-032, THR-033).

## 31. Evidence Return

Ein Evidence-Return-/Result-Package trägt: source isolated environment · destination environment · CorePack/activation references · operation/attempt references · per-content/per-target results · unknown outcomes · verification evidence · audit records · local exceptions · policy/trust snapshots · redaction/disclosure state · provenance · integrity · time uncertainty · known gaps.

```text
evidence returned ≠ evidence complete ≠ evidence sufficient
```

Export benötigt eigene Autorität; keine Rohsecrets/wiederverwendbaren Credentials (Companion 3 §27; Bezug THR-020, THR-025).

## 32. Audit and Evidence

Erfasst über den Lebenszyklus: assembly · manifest generation · validation · trust assessment · target binding · transfer · receipt · import · quarantine · assessment · approval · activation · partial results · update/delta · revocation · reinstatement · rollback · evidence return · closure. `CorePack audit record ≠ trust decision ≠ activation authorization`; Historie wird nicht umgeschrieben (Bezug THR-016, THR-017).

## 33. Security Invariants

Als Designanforderungen (keine implementierten Kontrollen):

1. CorePack remains separate from Domain Pack, artifact, deployment blueprint, backup and evidence package.
2. CorePack identity, version, revision, assembly, transfer, import and activation instances remain separate dimensions.
3. CorePack transfer, receipt, import, trust, approval and activation remain separate states.
4. CorePack contents do not inherit trust, validation or compatibility from the container automatically.
5. Content resolution binds concrete revisions; mutable aliases do not serve as final CorePack content binding.
6. Target binding is preserved through transfer, import and activation and is not replaced by local rename.
7. Quarantine release does not imply activation authorization.
8. Delta activation requires a confirmed exact base revision.
9. Partial and unknown activation results remain visible and block unsafe retry.
10. Revocation, reinstatement and rollback preserve history and require current trust, revocation and compatibility assessment.
11. Evidence-return packages are not treated as complete or sufficient automatically and carry no secrets.
12. A CorePack must not contain raw secrets or reusable real credentials.

Keine Invariante ist als implementierte Kontrolle dargestellt.

## 34. Threat References

Reale IDs aus [THREAT_SCENARIO_REGISTER.md](../security/THREAT_SCENARIO_REGISTER.md): manipuliertes Pack THR-021/THR-022/THR-024; kompromittierte Dependency THR-023; falsches Environment THR-022/THR-035; Replay/Duplicate THR-026/THR-004; Identifier Confusion THR-014; Secret Exposure THR-019/THR-020/THR-025; Audit Manipulation THR-016/THR-017; Partial/Unknown THR-031/THR-029/THR-028; Rollback/Backup THR-032/THR-033; Machine-Imitation THR-038. Keine erfundenen IDs.

## 35. Technology Boundary

Nicht ausgewählt: Pack-/Archivformat · Manifest-/Serialisierungsformat · Installer · Update-Engine · Transfer-/Removable-Media-Technologie · Hash/Signing · Trust Anchor. Alle bleiben `deferred`.

## 36. Open Questions

- Genaue Feldschemata für Manifest und Result Package (deferred, formatabhängig).
- Delta-Baseline-Bestätigungsverfahren ohne festgelegte Hash-/Signing-Technologie (Companion 3, deferred).
- Grenzwerte für „revocation snapshot too stale" je Profil (Companion 3 §18).

## 37. Next Decision

`CO-WP-024 – Secrets, Configuration Vault and Key Custody` gemäß lokaler [WORK_PACKAGE_QUEUE.md](../../project-system/WORK_PACKAGE_QUEUE.md). Zuerst Nova Review von CO-WP-023, danach Human-Maintainer-Commit.
