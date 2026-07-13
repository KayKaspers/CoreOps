# CoreOps – Repository Governance Standard

> Document Status: Implemented, pending Nova review
> Standard Status: Foundation repository-governance standard
> Human-Maintainer Authority: Preserved
> Automated Enforcement Status: Not implemented
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-005` (docs-only / repository-governance foundation)

## 1. Status

Foundation repository-governance standard consolidating roles/authority, source-of-truth hierarchy, document status and supersession, file naming, encoding/line endings, git boundaries, commit quality, public hygiene, and the PowerShell correction standard. It duplicates no existing rule; it references and extends the existing [FOUNDATION_SCOPE_LOCK.md](FOUNDATION_SCOPE_LOCK.md) Binding Source Hierarchy and the accepted Human-Maintainer gates (Decision Index DEC-G-01…08).

## 2. Purpose

Give CoreOps one clear reference for how the repository is governed so that authoritative sources are not overridden by summaries, Human-Maintainer gates stay intact, and repository hygiene is verifiable — without introducing automated enforcement.

## 3. Scope

Repository-wide governance for documentation-phase CoreOps: roles, source-of-truth, document lifecycle, naming, encoding, git read/write boundaries, commit quality, hygiene, and correction procedure.

## 4. Non-Goals

- No automated enforcement, CI, secret scanning, or link checking (later work).
- No `.gitignore`, `.gitattributes`, `.github/`, README, CONTRIBUTING, SECURITY, or CHANGELOG change.
- No mass file rename or CRLF/LF conversion.
- No application code, technology selection, or ADR.

## 5. Roles and Authority

- **Human Maintainer** — sole authority for git writes and irreversible/publication actions (§6) and final governance decisions (GO / GO WITH NOTES / REWORK / STOP).
- **Nova** — planning, architecture, review, evaluation.
- **Implementation Assistant (AI)** — executes one approved work package within allowed files; prepares changes, reviews diffs, may run tests only if explicitly allowed; never performs an irreversible repository action.

## 6. Human-Maintainer Gates

Human-Maintainer-only: `git add`; `git commit`; `git push`; branch creation/deletion; merge; tag; release; publication; risk acceptance; final GO / GO WITH NOTES / REWORK / STOP. AI systems may analyze, prepare allowed-file changes, review diffs, run tests when explicitly allowed, provide commit suggestions, and provide PowerShell check/correction blocks — but perform no irreversible repository action. Consistent with DEC-G-01/03.

## 7. Source-of-Truth Hierarchy

Extends the [FOUNDATION_SCOPE_LOCK.md](FOUNDATION_SCOPE_LOCK.md) Binding Source Hierarchy (authoritative). Consolidated ordering (highest first), compatible with it and with NDF v1.0.0 as framework base:

```text
1.  Applicable law and regulation
2.  Human-Maintainer decisions
3.  Security and safety invariants
4.  Accepted ADRs and scope locks (Foundation Scope Lock)
5.  Normative NDF v1.0.0 rules (Tag v1.0.0, Commit 9dcadc1)
6.  Project manifest
7.  Project Brief
8.  Decision Index
9.  Risk Register
10. Active Work Package Queue
11. Next-Phase state
12. Normative governance documents (this standard, language standard, neutrality policy, processes)
13. Architecture documents
14. Security baselines
15. Capability matrices
16. Project Profile
17. Project Brain
18. Context Packs
19. README and public summaries
20. Generated or derived views
```

Where this ordering and the Scope Lock hierarchy both apply, the Scope Lock remains authoritative for its listed items; this list extends it to further artifact classes. Ground rule:

```text
A summary must not silently override its authoritative source.
```

On conflict: identify the authoritative source; document the conflict; do not silently harmonize outside scope; escalate normative conflicts to the Human Maintainer.

## 8. Authoritative and Derived Artifacts

Semantic distinction (not all become global status values): `authoritative` · `normative` · `supporting` · `summary` · `derived` · `generated` · `historical` · `superseded` · `archived`. A `summary`/`derived`/`generated` artifact never overrides an `authoritative`/`normative` one.

## 9. Document Status

Uses the existing lifecycle vocabulary; no parallel incompatible list. Minimum covered states: draft; implemented-pending-review; accepted; accepted-with-notes; superseded; historical; archived. Decision Class, Lifecycle Status, and Binding Level remain separate dimensions (consistent with the Decision Index).

## 10. Supersession and Archiving

A superseded document must contain or reference: `superseded by`; supersession date or phase; reason; migration impact; historical retention decision. Superseded documents are not silently deleted where needed for decisions, risks, or historical evidence.

## 11. File Naming

New governance and standard documents use `UPPER_SNAKE_CASE.md` unless an existing directory norm requires otherwise. No spaces in new file names; no personal names; no unclear names (e.g. `final2.md`); no unnecessary parallel documents; prefer existing directories; use relative repository links; no absolute local paths in public documents.

## 12. Directory Governance

No new top-level structure without a dedicated architecture or governance work package. Prefer existing directories. New governance documents live under `docs/governance/`; security under `docs/security/`; system registers under `project-system/`.

## 13. Stable Identifiers

IDs for Work Packages, Decisions, Risks, Capabilities, PSR domains, and Feedback Candidates are not changed or reused without a documented migration. ID stability is a governance invariant (see also RISK-63).

## 14. Encoding

New text artifacts use UTF-8 (UTF-8 without BOM is acceptable unless an existing rule requires otherwise). No mass encoding migration of existing files in this or any WP without a dedicated approved migration.

## 15. Line Endings

Do not change existing line endings unnecessarily; avoid whole-file diffs caused only by CRLF/LF conversion; honor file conventions and `.gitattributes`; CRLF warnings alone are not a technical error; `git diff --check` remains binding; PowerShell write operations must handle encoding and line endings deliberately. No `.gitattributes` change in this WP.

## 16. Git Read and Write Boundaries

Read-only git is always allowed for verification (`status`, `log`, `diff`, `show`, `rev-parse`, `branch`, `remote get-url`, `tag --points-at`). Git write actions (`add`, `commit`, `push`, `pull`, `fetch`, `checkout`, `switch`, `reset`, `clean`, `tag`, `merge`, `rebase`, `stash`, `config`, `submodule`, `lfs`) are Human-Maintainer-only and never performed by AI systems.

## 17. Commit Quality

One work package per logically coherent commit; no mixing of independent changes; commit messages in English (imperative or established Conventional-Commit style); no claim of successful tests without actual execution; no commit with unresolved scope; no artificial commit for a pure blocked/no-change report; no automatic history rewrite; no AI force push. Example: `docs(governance): establish language and repository standards`.

## 18. Public Repository Hygiene

Minimum review for public changes: no secrets; no private paths; no personal data; no private infrastructure; no customer data; no unsupported claims; no certification overclaim; no broken relative references where checked; no accidental binary artifacts; no temporary editor files; no unrelated generated files; no misleading translation-parity claim. Unwanted file examples: `*.tmp`, `*.bak`, `*.orig`, `desktop.ini`, `Thumbs.db`, `.DS_Store`, editor swap files, local environment exports, debug dumps, unredacted logs. No `.gitignore` change here; a detected gap is a later follow-up.

## 19. Source Handoff

External sources must be identifiable; version/retrieval state documented for time-dependent sources; incomplete sources not treated as complete; local availability is not automatically trustworthiness; version-dependent claims need verified sources; missing source ingestion yields `verification-pending`, not invented detail.

## 20. External Content and Licensing

No extensive copyrighted content is reproduced; no proprietary framework text is reconstructed; license and attribution boundaries are honored (consistent with the ITIL/PRINCE2 tailoring §16 and the disclosure policy §15).

## 21. Blocked and No-Change Reports

A fail-closed stop without any file change is a complete, review-ready result when structured (reason, affected preflight check, unblock condition). No artificial intermediate commit is created for it (consistent with the adopted NDF work-package prompt safety baseline).

## 22. PowerShell Correction Standard

When a deterministic local correction is needed, a copyable PowerShell block must: explain the reason; name affected files; include `Set-Location "<COREOPS_REPOSITORY_ROOT>"`; use `Test-Path` / suitable preflight checks; guard error conditions with `throw` / `Write-Error`; provide a positive post-correction validation; and be followed by `git diff --check` and a targeted diff review. Prefer `[System.IO.File]::ReadAllText(...)` / `WriteAllText(...)` with explicit encoding for pure text substitution. A bare instruction like "edit line X manually" is avoided where a safe deterministic PowerShell correction is possible. AI systems still perform no git write.

## 23. Validation Checklist

```text
Only allowed files changed · No forbidden file changed · No secrets/private data ·
Human-Maintainer gates preserved · Source-of-truth not overridden by a summary ·
Stable IDs preserved · UTF-8 for new files · Line endings preserved ·
git diff --check clean · No mass conversion · Commit message in English
```

## 24. Exceptions

Any deviation touching security invariants, accepted ADRs, or NDF rules requires a documented exception with rationale, a risk entry, Nova review, and Human-Maintainer approval. No silent exceptions.

## 25. Compatibility

Additive; references and extends the Foundation Scope Lock Binding Source Hierarchy and accepted Human-Maintainer gates; introduces no automated enforcement; changes no accepted decision; selects no technology; creates no ADR. Breaking-change potential: low.

## 26. Open Questions

- Which document-status values become global vs. remain semantic distinctions?
- When is automated public-hygiene / link / secret-scanning enforcement scheduled (later WP)?
- Whether a `.gitignore`/`.gitattributes` review becomes a dedicated follow-up WP.

## 27. Next Decision

Nova review of this standard, followed by a Human-Maintainer commit. Automated enforcement, README/CONTRIBUTING alignment, and any `.gitignore`/`.gitattributes` review remain separate later work.
