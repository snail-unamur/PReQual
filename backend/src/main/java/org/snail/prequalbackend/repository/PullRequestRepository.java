package org.snail.prequalbackend.repository;

import org.snail.prequalbackend.model.PullRequest;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface PullRequestRepository extends MongoRepository<PullRequest, String> {
    List<PullRequest> findByOrgAndRepo(String org, String repo);
    Optional<PullRequest> findByOrgAndRepoAndMetaNumber(String org, String repo, int number);
}