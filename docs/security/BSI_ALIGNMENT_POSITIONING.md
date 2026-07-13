# CoreOps – BSI Alignment Positioning

> Status: **Accepted Product Direction** (Positionierung, keine Anforderungsmatrix)
> Certification Status: No certification claimed
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004A` (docs-only / governance-baseline)

Dieses Dokument ist eine **Positionierung**, keine IT-Grundschutz-Anforderungsmatrix und keine Zertifizierungsbehauptung. Es legt fest, welche BSI-bezogenen Aussagen zulässig und welche verboten sind.

## 1. Purpose

Klare, ehrliche und öffentlich neutrale Positionierung der BSI-Ausrichtung von CoreOps, um Fehlinterpretationen als Konformität, Zertifizierung oder Zulassung zu vermeiden.

## 2. BSI-oriented Development

CoreOps wird BSI-orientiert und an IT-Grundschutz-Prinzipien ausgerichtet entwickelt (Security First, Fail Closed, Least Privilege, Auditierbarkeit, Nachvollziehbarkeit). Diese Ausrichtung ist eine **Entwicklungsrichtung**, kein Nachweis.

## 3. IT-Grundschutz Alignment Goal

Ein konkretes IT-Grundschutz-Mapping (Bausteine, Anforderungen, Umsetzungshinweise, Nachweisartefakte) ist ein **Ziel** für spätere Security-Baseline-Work-Packages. Es existiert in diesem Work Package **kein** vollständiges Mapping und **keine** Anforderungsmatrix.

## 4. Public-Sector Readiness Goal

Public-Sector-Readiness (auditierbare, sicherheitsorientierte, nachweisfähige Betriebsumgebungen; Härtungsbaseline; Rollen- und Funktionstrennung; Protokollierung/Nachweis) ist ein **Produktziel**. Readiness ist kein Nachweis von Konformität und keine behördliche Anerkennung.

## 5. Claims That Are Allowed

```text
BSI-orientiert entwickelt
an IT-Grundschutz ausgerichtet
für auditierbare und sicherheitsorientierte Betriebsumgebungen vorbereitet
Public-Sector-Readiness als Produktziel
```

## 6. Claims That Are Prohibited

```text
BSI-zertifiziert
vollständig BSI-konform
BSI-zugelassen
VS-NfD-zugelassen
behördlich zertifiziert
```

Ebenfalls unzulässig: jede Darstellung einer externen Organisation (einschließlich Behörden) als Partner, Unterstützer oder Zertifizierer.

## 7. Certification Boundary

CoreOps behauptet ohne formale Prüfung **keine** vollständige BSI-Konformität, Zertifizierung oder Zulassung. Ein Government Profile (siehe [COREOPS_CONCEPT_V3_1_AMENDMENT.md](../architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md) §7) ist ein Konfigurations-, Betriebs- und Nachweisprofil, **keine** Zertifizierung.

## 8. VS Boundary

```text
VS-NfD und höhere Einstufungen sind für frühe CoreOps-Releases
keine zugesicherte Produkteigenschaft und kein aktuelles Releaseziel.
```

Die Architektur darf eine spätere Evaluierung nicht absichtlich verhindern; eine Eignung wird jedoch nicht versprochen.

## 9. Follow-up Baseline Work Package

Das konkrete IT-Grundschutz-Mapping, die Härtungsbaseline und die Nachweisartefakte werden in späteren Security-Baseline-Work-Packages (u. a. im Umfeld von `CO-WP-007 Threat Model and Trust Boundaries` sowie dedizierten BSI-/Hardening-Baseline-WPs) definiert. Dieses Work Package erzeugt keine solche Baseline.

## 10. Current Evidence Status

- BSI-Konformität: **nicht nachgewiesen, nicht behauptet**.
- Zertifizierung/Zulassung: **keine**.
- VS-Eignung: **kein aktuelles Ziel**.
- IT-Grundschutz-Mapping: **noch nicht vorhanden** (späteres WP).
- Government/Hardened Profile: **als Zielrichtung registriert**, nicht implementiert.

## 11. Readiness Baseline (CO-WP-004C)

Diese Positionierung wird additiv durch drei Baseline-Dokumente aus `CO-WP-004C` ergänzt (keine Zertifizierung, keine vollständige IT-Grundschutz-Compliance, keine Behörden- oder VS-NfD-Freigabe):

- [BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md](BSI_AND_PUBLIC_SECTOR_READINESS_BASELINE.md) — Claim Boundaries, 18 Readiness-Domänen (PSR-01…PSR-18), Verantwortungs- und Evidenzmodell, Offline/Air-Gap- und Souveränitäts-Grenzen.
- [BSI_REFERENCE_AND_CLAIMS_REGISTER.md](BSI_REFERENCE_AND_CLAIMS_REGISTER.md) — offizielles Referenzset (BSI 200-1…200-4, IT-Grundschutz-Kompendium versionsoffen, Mindeststandards, Protokollierung/Detektion v2.1; C5/C3A bedingt) mit Versions- und Anwendbarkeitsstatus.
- [PUBLIC_SECTOR_READINESS_PROFILE.md](../governance/PUBLIC_SECTOR_READINESS_PROFILE.md) — interne Standard-/Hardened-/Government-Profile.

**Produkt- vs. Deployment-Verantwortung:** Produktreife ≠ Deployment-Compliance; Compliance hängt von Architektur, Betreiberprozessen, Konfiguration, Evidenz, rechtlicher Anwendbarkeit und ggf. externer Bewertung ab. **Versionsgenaue Control-Mappings** (Klausel-zu-Capability) bleiben späteren Work Packages vorbehalten; hier werden **keine** einzelnen BSI-Anforderungen behauptet.
