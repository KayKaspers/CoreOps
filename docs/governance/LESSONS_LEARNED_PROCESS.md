# CoreOps – Lessons Learned Process

> Status: **Proposed for acceptance** (bindend nach Human-Maintainer-Commit)
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004B – Lessons Learned and NDF Feedback Governance` (docs-only / governance)

## 1. Status

Proposed for acceptance. Wird mit dem Human-Maintainer-Commit bindende Projektgovernance.

## 2. Purpose

Sicherstellen, dass Erkenntnisse aus Work Packages, Reviews, Blockern, Fehlern und Architekturentscheidungen systematisch erfasst, klassifiziert und projektlokal genutzt werden — und dass wiederverwendbare Erkenntnisse kontrolliert als NDF-Feedback-Kandidaten geprüft werden können, ohne das NDF selbst zu verändern.

## 3. Scope

Gilt für alle CoreOps-Work-Packages ab Foundation 0.1. Erfasst werden Prozess-, Prompt-, Tooling-, Security-, Architektur-, Betriebs- und Dokumentations-Erkenntnisse. Nicht erfasst: technische Implementierungsdetails, sobald Anwendungscode existiert (eigene spätere Prozesse).

## 4. Trigger

Eine Lessons-Learned-Prüfung wird ausgelöst durch:

- Abschluss eines Work Packages (Nova Review + Rückmeldung),
- einen `blocked`-Report,
- eine Nova-Bewertung `REWORK` oder `SPLIT`,
- eine erkannte Prompt- oder Scope-Abweichung,
- einen Sicherheits- oder Datenschutzvorfall (auch potenziell),
- eine erkannte fehlerhafte oder zu strenge/zu lockere Guardrail,
- einen Human-Maintainer-Hinweis außerhalb des formalen Reviews.

## 5. Lesson Lifecycle

```text
observed
→ validated
→ (project-action-planned | ndf-candidate | deferred | rejected | duplicate)
→ project-adopted (falls projektlokal umgesetzt)
→ superseded (falls durch neuere Lesson ersetzt)
→ closed
```

Ein Sprung direkt zu `closed` ohne dokumentierte Behandlung ist nicht zulässig.

## 6. Lesson Classification

Genau eine Primärklasse pro Lesson:

```text
LL-PROJECT
LL-NDF-CANDIDATE
LL-SECURITY
LL-PROCESS
LL-PROMPT
LL-TOOLING
LL-ARCHITECTURE
LL-OPERATIONS
LL-DOCUMENTATION
LL-NOT-REUSABLE
```

Sekundärklassen (beliebig viele der obigen, außer der bereits gewählten Primärklasse) sind zulässig, z. B. eine `LL-PROCESS`-Lesson mit Sekundärklasse `LL-NDF-CANDIDATE`.

## 7. Required Evidence

Jede Lesson mit Status `validated` oder höher benötigt nachvollziehbare Evidenz: Bezug auf Source Project, Source Work Package und eine neutralisierte Beschreibung der Beobachtung (kein Rohlog, keine Chain-of-Thought). Ist keine belastbare Evidenz vorhanden, bleibt die Lesson `observed` oder wird als `validation-required` markiert (siehe Register).

## 8. Project-local Treatment

Lessons mit ausschließlich CoreOps-spezifischer Relevanz erhalten `LL-PROJECT` (oder passende Primärklasse) und werden über `project-action-planned` → `project-adopted` in projektlokale Work Packages oder Governance-Dokumente überführt. Sie werden **nicht** automatisch NDF-Kandidaten.

## 9. NDF Candidate Evaluation

Eine Lesson wird nur dann als potenzieller NDF-Kandidat geprüft, wenn sie über CoreOps hinaus wiederverwendbar erscheint (`Reusable Beyond CoreOps: yes`). Die Bewertung und der weitere Prozess sind in [NDF_FEEDBACK_PROCESS.md](NDF_FEEDBACK_PROCESS.md) geregelt. Eine `ndf-candidate`-Markierung einer Lesson bedeutet **keine** NDF-Übernahme.

## 10. Security and Privacy

- Keine Secrets, keine privaten Namen, Pfade, Domains, Standorte oder IP-Adressen in Lessons oder Kandidaten.
- Sicherheitsrelevante Lessons erhalten `Security Relevance: yes` und werden bevorzugt geprüft (siehe Eskalationsregel in [NDF_FEEDBACK_PROCESS.md](NDF_FEEDBACK_PROCESS.md) §17).
- Keine Chain-of-Thought, keine vollständigen internen Logs.

## 11. Review Responsibilities

- **Nova:** prüft Lessons auf Plausibilität, Klassifikation und NDF-Kandidateneignung.
- **Human Maintainer:** entscheidet über projektlokale Adoption, NDF-Transferfreigabe und jede Statusänderung zu `approved-for-transfer` oder höher.
- **Implementation Agent:** erfasst Lessons, schlägt Klassifikation vor, setzt keine Freigabe-Status.

## 12. Closure and Supersession

Ein Eintrag wird nie ohne Begründung gelöscht. `superseded` verweist auf die ersetzende Lesson-ID. `closed` erfordert eine dokumentierte Behandlung (projektlokale Umsetzung, NDF-Transfer, bewusste Ablehnung oder Duplikat-Verweis).

## 13. Relationship to Project Brain

Wesentliche Lessons werden im Project Brain unter „Lessons Learned" kompakt referenziert (nicht dupliziert). Das vollständige Register bleibt [LESSONS_LEARNED_REGISTER.md](../../project-system/LESSONS_LEARNED_REGISTER.md).

## 14. Relationship to Work Packages

Jedes Work Package soll am Ende (Rückmeldung an Nova) erfassen, ob und welche Lessons entstanden sind. Dieser Prozess selbst wird retrospektiv auf frühere Work Packages angewendet (siehe Register).

## 15. Exceptions

Keine stillen Ausnahmen. Abweichungen von diesem Prozess benötigen ein eigenes Work Package, Nova Review und Human-Maintainer-Freigabe.
