package org.snail.prequalbackend.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.time.Instant;
import java.util.List;

@Document(collection = "pull_requests")
public class PullRequest {

    @Id
    public final String id;

    @Field("analysed_at")
    public final Instant analysedAt;

    public final Metrics base;
    public final Metrics head;
    public final List<Comment> comments;
    public final Metadata meta;
    public final Stats stats;
    public final List<Review> reviews;
    public final String org;
    public final String repo;

    public PullRequest(String id, Instant analysedAt, Metrics base, Metrics head,
                       List<Comment> comments, Metadata meta, Stats stats,
                       List<Review> reviews, String org, String repo) {
        this.id = id;
        this.analysedAt = analysedAt;
        this.base = base;
        this.head = head;
        this.comments = comments;
        this.meta = meta;
        this.stats = stats;
        this.reviews = reviews;
        this.org = org;
        this.repo = repo;
    }
}

