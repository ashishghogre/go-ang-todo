import { async, ComponentFixture, TestBed } from "@angular/core/testing";

import { SideDrawerComponent } from "./side-drawer.component";
import { ItemService } from "src/app/services/item.service";

describe("SideDrawerComponent", () => {
  let component: SideDrawerComponent;
  let itemService: ItemService;
  let fixture: ComponentFixture<SideDrawerComponent>;

  beforeEach(async(() => {
    itemService = jasmine.createSpyObj("ItemService", ["createItem"]);
    TestBed.configureTestingModule({
      declarations: [SideDrawerComponent],
      providers: [{ provide: ItemService, useValue: itemService }]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SideDrawerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it("should create", () => {
    expect(component).toBeTruthy();
  });

  it("should add new item to item service when create is clicked", () => {
    TestBed.get(ItemService).itemAdded = false;
    const sideDrawer: HTMLElement = fixture.nativeElement;
    sideDrawer.querySelector("button").click();
    expect(itemService.createItem).toHaveBeenCalledTimes(1);
  });

  it("should not add new item to item service, if an item is already added", () => {
    TestBed.get(ItemService).itemAdded = true;
    const sideDrawer: HTMLElement = fixture.nativeElement;
    sideDrawer.querySelector("button").click();
    expect(itemService.createItem).toHaveBeenCalledTimes(0);
  });
});
