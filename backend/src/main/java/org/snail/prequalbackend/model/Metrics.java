package org.snail.prequalbackend.model;

import org.springframework.data.mongodb.core.mapping.Field;

public record Metrics(
        @Field("code_smells") Integer codeSmells,
        @Field("cognitive_complexity") Integer cognitiveComplexity,
        Integer complexity,
        @Field("development_cost") Integer developmentCost,
        @Field("duplicated_lines") Integer duplicatedLines,
        Integer ncloc,
        @Field("software_quality_maintainability_rating") Integer softwareQualityMaintainabilityRating
) {}
