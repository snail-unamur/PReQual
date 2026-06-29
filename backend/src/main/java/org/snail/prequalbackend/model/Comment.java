package org.snail.prequalbackend.model;

import org.springframework.data.mongodb.core.mapping.Field;

public record Comment (
        String body,
        Author author,
        @Field("created_at") String createdAt
) {
}
