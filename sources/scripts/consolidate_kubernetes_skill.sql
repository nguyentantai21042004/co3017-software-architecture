-- =====================================================================
-- Consolidate all Kubernetes skill tags into single "kubernetes" tag
-- This reduces dashboard clutter from 6 cards to 1 card
-- =====================================================================

-- Update all kubernetes-related skill tags to just "kubernetes"
UPDATE questions
SET skill_tag = 'kubernetes'
WHERE skill_tag IN (
    'kubernetes_architecture',
    'kubernetes_workloads',
    'kubernetes_networking',
    'kubernetes_configuration',
    'kubernetes_core_concepts',
    'kubectl_commands'
);

-- Verify the consolidation
SELECT
    skill_tag,
    COUNT(*) as total_questions,
    SUM(CASE WHEN is_remedial THEN 1 ELSE 0 END) as remedial_count,
    SUM(CASE WHEN NOT is_remedial THEN 1 ELSE 0 END) as standard_count
FROM questions
WHERE skill_tag = 'kubernetes'
GROUP BY skill_tag;

-- Show sample questions
SELECT id, LEFT(content, 60) as preview, skill_tag, is_remedial
FROM questions
WHERE skill_tag = 'kubernetes'
LIMIT 5;
