# CoreOps – Domain Pack Governance Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation Domain-Pack governance
> Implementation Status: Not implemented
> Packaging Format: Not selected
> Plugin Runtime: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch CO-WP-015 (docs-only / modular product-governance and compatibility foundation)

## 1. Status

Technologie- und herstellerunabhängiges Governance-Modell für **CoreOps Domain Packs**. Companion zu [DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md](DOMAIN_PACK_SUPPORT_AND_COMPATIBILITY_MODEL.md) und [DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md](../security/DOMAIN_PACK_TRUST_PROVENANCE_AND_LIFECYCLE_POLICY.md). Es implementiert **kein** Domain Pack, keinen Adapter, kein Plugin-System, kein Paket-/Archivformat und keine Marketplace-/Update-Technologie.

## 2. Purpose

Domain Packs sollen später zusammengehörige Fähigkeiten, Integrationsprofile, Dokumentation, Kompatibilitätsangaben und optionale Implementierungsbausteine einer Betriebsdomäne **bündeln**. Ein Pack ist zuerst eine **Governance- und Produktgrenze**, nicht ausführbarer Code. Dieses Modell legt fest, wie Packs identifiziert, abgegrenzt, versioniert und über ihren Lebenszyklus regiert werden, sodass `pack activation ≠ runtime authority` und `active ≠ implemented ≠ maintained ≠ supported ≠ validated`.

## 3. Scope

Domain-Pack-Begriff · Pack-Klassen · Pack-Identität/stabile IDs · Statusdimensionen · Lifecycle · Ownership/Maintenance · Pack-Inhalte · Dependencies · Composition/Overlap · Community/External Packs · Vendor Neutrality · Offline-Betrachtungen · Audit/Evidence. Support Levels und Compatibility (Companion 2); Trust/Provenance/Security-Response (Companion 3).

## 4. Non-Goals

- Kein Pack-/Archivformat, keine Plugin-Runtime, kein Marketplace, kein Paketmanager.
- Keine Repository-Verzeichnisstruktur, kein Download-/Update-Mechanismus, keine automatische Dependency Resolution.
- Keine Signatur-/Kryptografietechnologie, kein Adapter-/Domain-Pack-Code.
- Keine vollständige Liste zukünftiger Pack-Instanzen, keine Vendor-Zertifizierung.
- Keine Behauptung implementierter/validierter Packs.

## 5. Domain-Pack Definition

Ein **Domain Pack** ist eine versionierte Governance- und Produktgrenze für eine zusammengehörige Betriebsdomäne. Es kann später enthalten oder referenzieren: capability mappings · integration classes · contract profiles · operation profiles · policy requirements · adapter requirements · documentation · validation profiles · compatibility declarations · known limitations · migration guidance.

**Grundregeln:**
```text
domain pack ≠ adapter
domain pack ≠ plugin
domain pack ≠ integration instance
domain pack ≠ deployment unit
domain pack ≠ product certification
domain pack ≠ vendor endorsement
```
Ein Domain Pack darf mehrere Adapter-/Integrationstypen beschreiben. Ein von einem Pack referenzierter Adapter wird **nicht** automatisch durch den Pack-Status validiert.

## 6. Pack Classes

Herstellerneutral; je Klasse: Purpose · Typical capabilities · Typical integration classes · Expected privilege range · Offline relevance · Security considerations · Compatibility dimensions · Validation expectations. **Keine vollständige Liste konkreter Pack-Instanzen; keine Anbieterbindung.**

| Domain-Pack-Klasse | Purpose (kurz) | Typ. Integration Classes | Priv.-Range | Threat refs |
|---|---|---|---|---|
| virtualisation | VM-Domäne | virtualisation platform | read→execution | THR-005, THR-034 |
| container-and-orchestration | Container/Cluster | container platform, orchestrator | read→deployment | THR-022, THR-034 |
| operating-system-management | OS-Verwaltung | operating system | read→execution | THR-006, THR-034 |
| bare-metal-management | Hardware/BMC | bare-metal management | read→execution | THR-034, THR-039 |
| network-management | Netzobjekte | network management | read→controlled write | THR-039, THR-034 |
| printer-and-print-management | Druck/Geräte | printer/print-management | read→controlled write | THR-025, THR-034 |
| monitoring-and-observability | Telemetrie/Health | monitoring/observability | read/observe | THR-012, THR-013 |
| backup-and-recovery | Backup/Restore | backup/recovery | read→execution | THR-033, THR-032 |
| deployment-and-release | Deployment | source-control/CI/CD, artifact repo | read→deployment | THR-021, THR-022, THR-023 |
| identity-and-access | Identität | identity/directory | read/controlled write | THR-001, THR-035 |
| ITSM-and-ticketing | Tickets/Changes | ITSM/ticketing | read→controlled write | THR-030, THR-018 |
| artifact-and-package-management | Artefakte/Pakete | artifact/package repository | read→deployment | THR-021, THR-023 |
| evidence-and-audit | Evidenz/Audit | evidence/audit | read/export | THR-016, THR-017, THR-018 |
| notification-and-communication | Benachrichtigung | notification system | read/controlled write | THR-018 |
| offline-and-isolated-operation | Offline/Air-Gap | offline transfer integration | read→controlled (aktiviert) | THR-024, THR-025 |

## 7. Pack Identity

Jedes Domain Pack benötigt mindestens: stable pack identity · canonical name · display name · domain class · owner · maintainer · version · lifecycle state · maintenance state · support level · implementation state · validation state · evidence state · security-review state · compatible contract versions · compatible CoreOps versions · target-system scope · dependency references · known limitations · deprecation status · audit reference.

**Pack-IDs müssen stabil sein und dürfen nicht wiederverwendet werden.** Display Name/Übersetzung ersetzen die stabile Pack-ID nicht (konsistent mit [Language Standard](../governance/COREOPS_LANGUAGE_STANDARD.md)). **Keine vollständige Pack-ID-Liste erzeugt.**

## 8. Status Dimensions

Mindestens **getrennt** gehalten (Detail in Companion 2): Pack Lifecycle · Maintenance Responsibility · Support Level · Implementation Status · Validation Status · Evidence Status · Security Review Status · Compatibility Status · Release Status.
```text
active     ≠ implemented
implemented ≠ maintained
maintained ≠ supported
supported  ≠ validated
validated  ≠ universally compatible
released   ≠ supported
deprecated ≠ immediately unusable
```
**Keine kombinierten Pseudostatuswerte.**

## 9. Lifecycle

Konzeptionelle Statuswerte (keine implementierten Enums):
```text
proposed → draft → preview → active → suspended → deprecated → security-maintenance-only → retired → archived
```
- `preview` erzeugt **keinen** stabilen Support-Claim.
- `active` bedeutet **nicht** automatisch implementiert oder supported.
- `deprecated` benötigt Migrations- und End-of-Maintenance-Informationen.
- `retired` darf **keine** neuen Compatibility Claims erhalten.
- Historische IDs und Evidenz bleiben erhalten.

## 10. Ownership and Maintenance

Maintenance Responsibility (getrennt vom Support Level, Detail Companion 2 §7): `unassigned · community-maintained · project-maintained · external-maintainer · joint-maintenance · maintenance-suspended · retired`. Sie beschreibt, **wer** Änderungen/Reviews verantwortet — keine Aussage über Support Level, SLA, Validation, Vendor Support oder Security Certification. Für externe Maintainer sind Owner, Kontaktweg und Governance-Grenze dokumentierbar, **ohne private Kontaktdaten in öffentlichen Beispielen** (konsistent mit [Public Neutrality](../governance/PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md)).

## 11. Pack Contents

Ein Pack kann Capability Mappings, Integration-/Contract-/Operation-Profile, Policy-/Adapter-Requirements, Dokumentation, Validation-Profile, Compatibility Declarations, Known Limitations und Migration Guidance enthalten oder referenzieren. Inhalte sind **begrenzt** auf die deklarierte Domäne (§8 Scope); Inhalte gewähren keine Autorität (§18 Invarianten). Referenzen auf Capabilities folgen der [Capability Matrix Spec](../project-system/CAPABILITY_MATRIX_SPEC.md); referenzierte Integration-Profile folgen dem [Integration Contract](COREOPS_INTEGRATION_CONTRACT_V0_1.md).

## 12. Dependencies

Zusammenfassung (Detail Companion 2 §15). Jede Dependency dokumentiert mindestens: dependency identity/type · required/optional · minimum/bounded compatibility · reason · security relevance · offline availability · fallback behavior · deprecation impact · validation status. Klassen: CoreOps module · Integration Contract · other Domain Pack · adapter/agent · external-system · data/schema · policy · validation-profile dependency. **Optionale Dependencies werden nicht still zur Core-Pflicht**; zirkuläre Dependencies bleiben sichtbar; keine automatische Dependency Resolution.

## 13. Composition and Overlap

Packs können überlappende Capabilities/Zielsysteme betreffen. Dokumentiert: primary authority · shared capability · extension relationship · composition relationship · conflict owner · precedence rule · fallback behavior · compatibility impact. Regeln:
- Overlap erzeugt **keine** automatische doppelte Autorität.
- Ein Pack überschreibt **keine** Policy eines anderen Packs still.
- Mehrere Packs dürfen denselben Adapter referenzieren.
- Pack-Aktivierung erzeugt **keine** globale Capability-Freigabe.
- Konflikte bleiben sichtbar und auditierbar. Keine Pack-Composition-Engine ausgewählt.

## 14. Community and External Packs

Community-/extern gepflegte Packs sind zulässig und müssen mindestens offenlegen (Detail Companion 3 §11): source · owner · maintainer model · governance status · support level · compatibility claims · validation state · known limitations · dependency provenance · security reporting path · deprecation policy.
```text
community pack     ≠ trusted pack
external maintainer ≠ CoreOps endorsement
public repository  ≠ verified provenance
popular pack       ≠ supported pack
```
CoreOps stuft externe Packs **nicht** automatisch als `SUP-2`/`SUP-3` ein.

## 15. Vendor Neutrality

Herstellerbezüge dienen der sachlichen Kompatibilitätsbeschreibung und bedeuten **nicht** Endorsement, Partnerschaft, Zertifizierung, Herstellerfreigabe oder Herstellersupport. Vendor-spezifische Pack-Erweiterungen überschreiben **keine** Core-Sicherheitsinvarianten (konsistent mit [Integration Contract §25](COREOPS_INTEGRATION_CONTRACT_V0_1.md)).

## 16. Offline Considerations

Offline verfügbare Packs/Metadaten benötigen (Detail Companion 3 §16): pack identity · version · target environment · dependency set · contract compatibility · provenance/integrity status · validity/usage boundary · import quarantine · local activation decision · revocation-distribution challenge · reconciliation requirement · audit continuity. Nicht behauptet: implementierte Offline Distribution, beliebige Air-Gap-Unterstützung, Klassifiziertnetz-Eignung, konkrete Signing-/Trust-Anchor-Technologie.

## 17. Audit and Evidence

Erfasst: pack creation · ownership/maintainer change · version · lifecycle transition · support-level change · compatibility claim · validation activity · evidence reference · dependency change · security finding · suspension · deprecation · retirement · offline import · withdrawal/revocation. Trennung:
```text
support claim       ≠ validation evidence
validation evidence ≠ compatibility for every target
compatibility evidence ≠ vendor certification
```

## 18. Security Invariants

Als **Designanforderungen** (nicht implementierte Kontrollen):

1. Pack activation must not grant runtime authority automatically.
2. Support level must not imply implementation or validation.
3. Compatibility claims must remain version-, profile- and evidence-bound.
4. Community or external origin must not imply trust.
5. Pack dependencies must not silently become mandatory Core dependencies.
6. Deprecated or retired packs must not receive new unsupported claims.
7. A compromised pack or maintainer must be suspendable.
8. Offline pack use requires provenance, integrity, target binding and explicit activation.
9. Pack overlap must not create duplicate or conflicting authority silently.
10. Vendor-specific extensions must not override CoreOps security invariants.
11. Validation evidence must not be generalized beyond its tested scope.
12. Historical pack identity and evidence must not be deleted during retirement.

## 19. Threat References

Reale IDs aus dem [Threat Scenario Register](../security/THREAT_SCENARIO_REGISTER.md) (nur reale IDs, keine Duplikation, kein Parallelregister): THR-001, THR-005, THR-006, THR-008, THR-010, THR-011, THR-012, THR-013, THR-016, THR-017, THR-018, THR-021, THR-022, THR-023, THR-024, THR-025, THR-026, THR-030, THR-032, THR-033, THR-034, THR-035, THR-039.

## 20. Technology Boundary

Nicht ausgewählt/implementiert: Pack-/Archivformat, Plugin-Runtime, Marketplace, Paketmanager, Repository-Struktur, Download-/Update-Mechanismus, automatische Dependency Resolution, Signatur-/Kryptografietechnologie, Pack-/Adapter-Code.

## 21. Compatibility

Konsistent mit [Capability Matrix](FOUNDATION_CAPABILITY_MATRIX.md)/[Spec](../project-system/CAPABILITY_MATRIX_SPEC.md), [Module Architecture](COREOPS_LOGICAL_MODULE_ARCHITECTURE.md)/[Catalog](COREOPS_MODULE_CATALOG.md), [Integration Contract](COREOPS_INTEGRATION_CONTRACT_V0_1.md)/[Capability/Operation](INTEGRATION_CAPABILITY_AND_OPERATION_MODEL.md), [Integration Trust/Failure](../security/INTEGRATION_TRUST_FAILURE_AND_RECOVERY_POLICY.md), [Policy/Approval/Execution](../security/POLICY_DECISION_AND_EVALUATION_MODEL.md), [Foundation Scope Lock](../governance/FOUNDATION_SCOPE_LOCK.md). Ergänzt DEC-P-01, DEC-S-16…22 (interne Profile) und die ITIL/PRINCE2-Tailoring-Linie.

## 22. Open Questions

- Verbindliche SemVer-Policy für Pack-Versionen (mit DEC-O-02 verbunden, später).
- Präzedenzregeln bei Pack-Overlap im Detail (spätere Iteration).
- Verhältnis Domain Pack ↔ späteres Deployment-/Blueprint-Modell (CO-WP-021).

## 23. Next Decision

Companion 2 (Support/Compatibility) und Companion 3 (Trust/Provenance/Lifecycle) konkretisieren dieses Gerüst. Deployment Control Plane (CO-WP-021) und Artifact Trust (CO-WP-022) bauen darauf auf. Packaging-/Plugin-/Marketplace-Technologie bleibt einer späteren ADR-Runde vorbehalten.
