# CoreOps – Language Standard

> Document Status: Implemented, pending Nova review
> Standard Status: Foundation language standard
> Primary Technical Language: English
> Primary Product Languages: German and English
> Translation Completion Claim: None
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-005` (docs-only / repository-governance foundation)

## 1. Status

Foundation language standard for CoreOps. It defines which languages are used for which artifact classes, the canonical language, the translation-status model, semantic parity, and terminology handling. It translates no existing content and introduces no i18n dependency. It formalizes the existing DE/EN product direction (Decision Index DEC-P-03).

## 2. Purpose

Provide clear, project-wide language rules so that machine-facing content stays consistent and unambiguous, product-facing content can be German and English, and no unverified full-translation parity is ever claimed.

## 3. Scope

- Language rules for all CoreOps artifact classes (§5).
- Canonical/technical language (English) and product languages (German and English).
- Translation-status metadata and semantic-parity requirements.
- Terminology and technical-identifier handling.

## 4. Non-Goals

- No full translation of existing documents.
- No UI localization, API internationalization, or i18n dependency.
- No complete glossary database.
- No README overhaul.
- No claim of full bilingual parity.

## 5. Artifact Classes

- **Machine-facing artifacts** — source code; API and schema identifiers; configuration keys; CLI commands and flags; machine-readable status values; Capability/Decision/Risk/PSR/Candidate IDs; commit messages; technical log-event IDs; file names of new technical/governance artifacts.
- **Public user-facing documentation** — setup content, user guides, help, UI content.
- **Internal project governance** — Project Brain, Context Packs, Work Packages, registers, indices.
- **Logs and audit events** — machine event IDs vs human-readable messages.
- **Examples and sample data** — neutral, synthetic (see [PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md](PUBLIC_NEUTRALITY_AND_DISCLOSURE_POLICY.md)).

## 6. Canonical Language Rules

English is canonical and mandatory for machine-facing artifacts: source-code identifiers (classes, functions, variables); API routes and fields; schema fields; configuration keys; CLI commands and flags; machine-readable status values; Capability and Decision IDs; file names of new technical and governance artifacts; Git commit messages; technical log-event IDs. These are **not** translated.

## 7. German and English Product Scope

- CoreOps remains a DE/EN project; German and English are the primary product-facing languages.
- Translation parity may be claimed **only** when actually reviewed.
- Not every internal Project-Brain or Work-Package file must be fully bilingual.
- Public help, setup content, user interfaces, and central user guidance should later support DE/EN.
- This work package does not translate existing content.

## 8. Governance Document Language

Governance documents may be: English with a clear German short introduction; German with a clear English short summary; or fully bilingual. Each must carry an unambiguous language status. (Existing NDF-derived documents use a `Sprachstatus / Language status` line; that convention remains valid.)

## 9. Translation Status Model

Language-status metadata (required where §16 says so, not physically on every file):

```text
Primary Language:        de | en
Secondary Language:      de | en | none
Translation Status:      not-required | source-only | translation-pending |
                         partially-translated | synchronized | review-required
Canonical Language:      de | en | language-independent
Last Translation Review: date | phase | not-reviewed
```

## 10. Semantic Parity

```text
Translated text must preserve meaning, security boundaries, claim boundaries and normative strength.
```

Not permitted: a binding statement in only one language; different security boundaries between DE and EN; translations that soften `must` to a recommendation; outdated translations without a status marker; an automatic claim of full language parity.

## 11. Technical Identifiers

IDs, status values, API fields, configuration keys, and file names are never translated. Vendor and product names are not translated. Legally or normatively fixed terms stay close to their source.

## 12. Terminology Consistency

- Technical key terms preferably in English.
- German explanations may carry the English technical term in parentheses.
- A term is not translated inconsistently within the same document without reason.
- No full glossary database is created here (a later glossary WP may add one).

## 13. User Interface and Documentation Boundary

UI localization and central user documentation should later support DE/EN, tracked via the translation-status model. This standard sets the rules; implementation and actual translation are later work packages, not this one.

## 14. Logs and Audit Events

Machine event identifiers use English and stable IDs. Human-readable log/audit messages may be localized later; the machine-facing identifier remains English and untranslated so evidence stays correlatable.

## 15. Examples

Examples and sample data are neutral and synthetic (see the disclosure policy §12). Example text may be DE or EN; identifiers within examples stay English.

## 16. Exceptions

A document may omit language-status fields when it is purely machine-facing or clearly source-only. Any binding governance document must carry a language status. Exceptions that weaken semantic parity require a documented rationale, Nova review, and Human-Maintainer approval.

## 17. Review Triggers

- A new public user-facing surface (UI, setup, help) that needs DE/EN.
- A governance document whose translation drifts from its source.
- A claim of translation parity (must be evidence-based).
- A later glossary or localization work package.

## 18. Compatibility

Additive; formalizes DE/EN direction (DEC-P-03) without changing accepted decisions. Consistent with the NDF `Sprachstatus / Language status` convention and the repository-governance standard. No technology selected; no ADR; breaking-change potential: low.

## 19. Open Questions

- Which documents must carry full language-status metadata vs. a short status line?
- When is a DE/EN glossary work package scheduled?
- Which public surfaces are first to receive reviewed DE/EN parity?

## 20. Next Decision

Nova review of this language standard, followed by a Human-Maintainer commit. Actual translation, localization, and glossary work remain separate later work packages.
