package org.snail.prequalbackend.controller;

import org.snail.prequalbackend.dto.PullRequestDTO;
import org.snail.prequalbackend.service.PullRequestService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collection;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/pr")
public class PullRequestController {

    private final PullRequestService pullRequestService;

    public PullRequestController(PullRequestService pullRequestService) {
        this.pullRequestService = pullRequestService;
    }

    @GetMapping("/{organisation}/{repo}/{id}")
    public PullRequestDTO getPR(
            @PathVariable String organisation,
            @PathVariable String repo,
            @PathVariable int id) {
        return PullRequestDTO.fromPullRequest(pullRequestService.getPullRequest(organisation, repo, id));
    }

    @GetMapping("/{organisation}/{repo}")
    public Collection<PullRequestDTO> getAllPR(
            @PathVariable String organisation,
            @PathVariable String repo) {
        return pullRequestService.getAllPullRequests(organisation, repo)
                .stream()
                .map(PullRequestDTO::fromPullRequest)
                .collect(Collectors.toList());
    }
}
