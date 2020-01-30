import { TestBed, inject } from "@angular/core/testing";

import { ItemService } from "./item.service";
import {
  HttpClientTestingModule,
  HttpTestingController
} from "@angular/common/http/testing";
import { Item } from "../models/item";

describe("ItemService", () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [ItemService]
    });
  });

  it("should be created", () => {
    const service: ItemService = TestBed.get(ItemService);
    expect(service).toBeTruthy();
  });

  it("should get items from api and set it in instance variable", inject(
    [HttpTestingController, ItemService],
    (httpMock: HttpTestingController, service: ItemService) => {
      expect(service.getAllItems().length).toEqual(0);

      service.getItems();

      const itemList: Item[] = [];
      itemList.push({ id: 123, title: "Title", details: "Details" } as Item);

      const request = httpMock.expectOne("http://localhost:8080/items");
      expect(request.request.method).toEqual("GET");
      request.flush([...itemList]);

      expect(service.getAllItems().length).toEqual(1);
    }
  ));

  it("should create item and add it to the list", () => {
    const service: ItemService = TestBed.get(ItemService);

    expect(service.getAllItems().length).toEqual(0);
    service.createItem();

    expect(service.getAllItems().length).toEqual(1);
  });
});
