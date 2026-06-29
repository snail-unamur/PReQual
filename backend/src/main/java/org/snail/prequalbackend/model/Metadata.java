package org.snail.prequalbackend.model;

import org.springframework.data.mongodb.core.mapping.Field;

public record Metadata(
        @Field("id") String id,
        int number,
        Author author,
        String title,
        String body,
        String state,
        @Field("created_at") String createdAt,
        @Field("closed_at") String closedAt,
        @Field("merged_at") String mergedAt
) {}
