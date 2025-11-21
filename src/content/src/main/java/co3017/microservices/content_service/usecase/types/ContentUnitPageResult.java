package co3017.microservices.content_service.usecase.types;

import co3017.microservices.content_service.models.ContentUnit;

import java.util.List;

/**
 * Page result cho ContentUnit
 */
public class ContentUnitPageResult {
    private List<ContentUnit> items;
    private long totalElements;
    private int totalPages;
    private int currentPage;
    private int size;
    private boolean first;
    private boolean last;

    public ContentUnitPageResult(List<ContentUnit> items, long totalElements, int totalPages,
            int currentPage, int size, boolean first, boolean last) {
        this.items = items;
        this.totalElements = totalElements;
        this.totalPages = totalPages;
        this.currentPage = currentPage;
        this.size = size;
        this.first = first;
        this.last = last;
    }

    public List<ContentUnit> getItems() {
        return items;
    }

    public long getTotalElements() {
        return totalElements;
    }

    public int getTotalPages() {
        return totalPages;
    }

    public int getCurrentPage() {
        return currentPage;
    }

    public int getSize() {
        return size;
    }

    public boolean isFirst() {
        return first;
    }

    public boolean isLast() {
        return last;
    }
}
