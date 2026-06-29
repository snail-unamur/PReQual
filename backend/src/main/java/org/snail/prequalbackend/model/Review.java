package org.snail.prequalbackend.model;

import org.springframework.data.mongodb.core.mapping.Field;

public record Review(
        Author author,
        String state,
        String body,
        @Field("submitted_at") String submittedAt
) {}
