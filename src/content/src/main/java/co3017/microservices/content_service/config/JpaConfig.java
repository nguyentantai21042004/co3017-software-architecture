package co3017.microservices.content_service.config;

import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@Configuration
@EnableJpaRepositories(basePackages = {
        "co3017.microservices.content.repository"
})
@EntityScan(basePackages = {
        "co3017.microservices.content.entity"
})
public class JpaConfig {
}
