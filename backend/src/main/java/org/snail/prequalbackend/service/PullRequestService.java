package org.snail.prequalbackend.service;

import org.snail.prequalbackend.model.PullRequest;
import org.snail.prequalbackend.repository.PullRequestRepository;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class PullRequestService {

    private final PullRequestRepository repository;

    public PullRequestService(PullRequestRepository repository) {
        this.repository = repository;
    }

    @Cacheable(value = "prCache", key = "#org + ':' + #repo + ':' + #prNumber")
    public PullRequest getPullRequest(String org, String repo, int prNumber) {
        return repository.findByOrgAndRepoAndMetaNumber(org, repo, prNumber)
                .orElseThrow(() -> new RuntimeException("Pull Request not found: " + org + "/" + repo + "#" + prNumber));
    }

    @Cacheable(value = "repoCache", key = "#org + ':' + #repo")
    public List<PullRequest> getAllPullRequests(String org, String repo) {
        List<PullRequest> pullRequests = repository.findByOrgAndRepo(org, repo);
        if (pullRequests.isEmpty()) {
            throw new RuntimeException("No pull requests found for: " + org + "/" + repo);
        }
        return pullRequests;
    }
}

