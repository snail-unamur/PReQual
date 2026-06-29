package org.snail.prequalbackend.model;

import org.springframework.data.mongodb.core.mapping.Field;

public record Stats(
        @Field("base_size") String baseSize,
        @Field("head_size") String headSize,
        @Field("total_time") String totalTime
) {}
