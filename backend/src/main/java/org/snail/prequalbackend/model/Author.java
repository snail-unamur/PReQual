package org.snail.prequalbackend.model;

import org.springframework.data.mongodb.core.mapping.Field;

public record Author(
        String login,
        @Field("is_bot") boolean isBot
) {
}
