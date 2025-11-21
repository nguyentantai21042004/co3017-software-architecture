package co3017.microservices.content_service.usecase.service;

import co3017.microservices.content_service.models.ContentVersion;
import co3017.microservices.content_service.repository.ContentVersionRepository;
import co3017.microservices.content_service.usecase.types.CreateContentVersionCommand;
import co3017.microservices.content_service.usecase.types.UpdateContentVersionCommand;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class ContentVersionServiceTest {

    @Mock
    private ContentVersionRepository repository;

    @InjectMocks
    private ContentVersionService service;

    private ObjectMapper objectMapper;

    @BeforeEach
    void setUp() {
        objectMapper = new ObjectMapper();
    }

    @Test
    void create_ValidCommand_ReturnsContentVersion() {
        // Arrange
        UUID unitId = UUID.randomUUID();
        ObjectNode contentData = objectMapper.createObjectNode();
        contentData.put("text", "Sample content");

        CreateContentVersionCommand command = new CreateContentVersionCommand(unitId, "1.0.0", contentData, true);

        ContentVersion savedVersion = new ContentVersion(1L, unitId, "1.0.0", contentData, true, null, null);

        when(repository.save(any(ContentVersion.class))).thenReturn(savedVersion);

        // Act
        ContentVersion result = service.create(command);

        // Assert
        assertNotNull(result);
        assertEquals(1L, result.getVersionId());
        assertEquals("1.0.0", result.getVersionNumber());
        verify(repository).deactivateAllVersionsByUnitId(unitId); // Since active is true
        verify(repository).save(any(ContentVersion.class));
    }

    @Test
    void create_NullUnitId_ThrowsException() {
        CreateContentVersionCommand command = new CreateContentVersionCommand(null, "1.0.0", null, false);

        assertThrows(IllegalArgumentException.class, () -> service.create(command));
    }

    @Test
    void detail_ExistingId_ReturnsOptional() {
        Long id = 1L;
        ContentVersion version = new ContentVersion(id, UUID.randomUUID(), "1.0.0", objectMapper.createObjectNode(),
                true, null, null);

        when(repository.findById(id)).thenReturn(Optional.of(version));

        Optional<ContentVersion> result = service.detail(id);

        assertTrue(result.isPresent());
        assertEquals(id, result.get().getVersionId());
    }

    @Test
    void update_ExistingVersion_UpdatesAndReturns() {
        Long id = 1L;
        UUID unitId = UUID.randomUUID();
        ContentVersion existing = new ContentVersion(id, unitId, "1.0.0", objectMapper.createObjectNode(), false, null,
                null);

        UpdateContentVersionCommand command = new UpdateContentVersionCommand("1.1.0", null, true);

        when(repository.findById(id)).thenReturn(Optional.of(existing));
        when(repository.save(any(ContentVersion.class))).thenAnswer(i -> i.getArguments()[0]);

        Optional<ContentVersion> result = service.update(id, command);

        assertTrue(result.isPresent());
        assertEquals("1.1.0", result.get().getVersionNumber());
        assertTrue(result.get().isActive());
        verify(repository).deactivateAllVersionsByUnitId(unitId);
    }

    @Test
    void setActiveVersion_ValidIdAndUnit_UpdatesStatus() {
        Long id = 1L;
        UUID unitId = UUID.randomUUID();
        ContentVersion version = new ContentVersion(id, unitId, "1.0.0", objectMapper.createObjectNode(), false, null,
                null);

        when(repository.findById(id)).thenReturn(Optional.of(version));
        when(repository.save(any(ContentVersion.class))).thenAnswer(i -> i.getArguments()[0]);

        Optional<ContentVersion> result = service.setActiveVersion(id, unitId);

        assertTrue(result.isPresent());
        assertTrue(result.get().isActive());
        verify(repository).deactivateAllVersionsByUnitId(unitId);
    }
}
