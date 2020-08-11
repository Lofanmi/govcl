
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TSpinEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewSpinEdit(owner IComponent) *TSpinEdit {
    s := new(TSpinEdit)
    s.instance = SpinEdit_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsSpinEdit(obj interface{}) *TSpinEdit {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TSpinEdit{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsSpinEdit.
func SpinEditFromInst(inst uintptr) *TSpinEdit {
    return AsSpinEdit(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsSpinEdit.
func SpinEditFromObj(obj IObject) *TSpinEdit {
    return AsSpinEdit(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsSpinEdit.
func SpinEditFromUnsafePointer(ptr unsafe.Pointer) *TSpinEdit {
    return AsSpinEdit(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (s *TSpinEdit) Free() {
    if s.instance != 0 {
        SpinEdit_Free(s.instance)
        s.instance, s.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (s *TSpinEdit) Instance() uintptr {
    return s.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (s *TSpinEdit) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (s *TSpinEdit) IsValid() bool {
    return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (s *TSpinEdit) Is() TIs {
    return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (s *TSpinEdit) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TSpinEditClass() TClass {
    return SpinEdit_StaticClassType()
}

// 清除。
func (s *TSpinEdit) Clear() {
    SpinEdit_Clear(s.instance)
}

// 清除选择。
func (s *TSpinEdit) ClearSelection() {
    SpinEdit_ClearSelection(s.instance)
}

// 复制到粘贴板。
func (s *TSpinEdit) CopyToClipboard() {
    SpinEdit_CopyToClipboard(s.instance)
}

// 剪切到粘贴板。
func (s *TSpinEdit) CutToClipboard() {
    SpinEdit_CutToClipboard(s.instance)
}

// 从剪切板粘贴。
func (s *TSpinEdit) PasteFromClipboard() {
    SpinEdit_PasteFromClipboard(s.instance)
}

// 撤销上一次操作。
func (s *TSpinEdit) Undo() {
    SpinEdit_Undo(s.instance)
}

// 全选。
func (s *TSpinEdit) SelectAll() {
    SpinEdit_SelectAll(s.instance)
}

// 是否可以获得焦点。
func (s *TSpinEdit) CanFocus() bool {
    return SpinEdit_CanFocus(s.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (s *TSpinEdit) ContainsControl(Control IControl) bool {
    return SpinEdit_ContainsControl(s.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (s *TSpinEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(SpinEdit_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (s *TSpinEdit) DisableAlign() {
    SpinEdit_DisableAlign(s.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (s *TSpinEdit) EnableAlign() {
    SpinEdit_EnableAlign(s.instance)
}

// 查找子控件。
//
// Find sub controls.
func (s *TSpinEdit) FindChildControl(ControlName string) *TControl {
    return AsControl(SpinEdit_FindChildControl(s.instance, ControlName))
}

func (s *TSpinEdit) FlipChildren(AllLevels bool) {
    SpinEdit_FlipChildren(s.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (s *TSpinEdit) Focused() bool {
    return SpinEdit_Focused(s.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (s *TSpinEdit) HandleAllocated() bool {
    return SpinEdit_HandleAllocated(s.instance)
}

// 插入一个控件。
//
// Insert a control.
func (s *TSpinEdit) InsertControl(AControl IControl) {
    SpinEdit_InsertControl(s.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (s *TSpinEdit) Invalidate() {
    SpinEdit_Invalidate(s.instance)
}

// 移除一个控件。
//
// Remove a control.
func (s *TSpinEdit) RemoveControl(AControl IControl) {
    SpinEdit_RemoveControl(s.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (s *TSpinEdit) Realign() {
    SpinEdit_Realign(s.instance)
}

// 重绘。
//
// Repaint.
func (s *TSpinEdit) Repaint() {
    SpinEdit_Repaint(s.instance)
}

// 按比例缩放。
//
// Scale by.
func (s *TSpinEdit) ScaleBy(M int32, D int32) {
    SpinEdit_ScaleBy(s.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (s *TSpinEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    SpinEdit_ScrollBy(s.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (s *TSpinEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    SpinEdit_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (s *TSpinEdit) SetFocus() {
    SpinEdit_SetFocus(s.instance)
}

// 控件更新。
//
// Update.
func (s *TSpinEdit) Update() {
    SpinEdit_Update(s.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (s *TSpinEdit) BringToFront() {
    SpinEdit_BringToFront(s.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (s *TSpinEdit) ClientToScreen(Point TPoint) TPoint {
    return SpinEdit_ClientToScreen(s.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (s *TSpinEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return SpinEdit_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (s *TSpinEdit) Dragging() bool {
    return SpinEdit_Dragging(s.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (s *TSpinEdit) HasParent() bool {
    return SpinEdit_HasParent(s.instance)
}

// 隐藏控件。
//
// Hidden control.
func (s *TSpinEdit) Hide() {
    SpinEdit_Hide(s.instance)
}

// 发送一个消息。
//
// Send a message.
func (s *TSpinEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return SpinEdit_Perform(s.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (s *TSpinEdit) Refresh() {
    SpinEdit_Refresh(s.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (s *TSpinEdit) ScreenToClient(Point TPoint) TPoint {
    return SpinEdit_ScreenToClient(s.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (s *TSpinEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return SpinEdit_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (s *TSpinEdit) SendToBack() {
    SpinEdit_SendToBack(s.instance)
}

// 显示控件。
//
// Show control.
func (s *TSpinEdit) Show() {
    SpinEdit_Show(s.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (s *TSpinEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return SpinEdit_GetTextBuf(s.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (s *TSpinEdit) GetTextLen() int32 {
    return SpinEdit_GetTextLen(s.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (s *TSpinEdit) SetTextBuf(Buffer string) {
    SpinEdit_SetTextBuf(s.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TSpinEdit) FindComponent(AName string) *TComponent {
    return AsComponent(SpinEdit_FindComponent(s.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TSpinEdit) GetNamePath() string {
    return SpinEdit_GetNamePath(s.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TSpinEdit) Assign(Source IObject) {
    SpinEdit_Assign(s.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TSpinEdit) ClassType() TClass {
    return SpinEdit_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TSpinEdit) ClassName() string {
    return SpinEdit_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TSpinEdit) InstanceSize() int32 {
    return SpinEdit_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TSpinEdit) InheritsFrom(AClass TClass) bool {
    return SpinEdit_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TSpinEdit) Equals(Obj IObject) bool {
    return SpinEdit_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TSpinEdit) GetHashCode() int32 {
    return SpinEdit_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TSpinEdit) ToString() string {
    return SpinEdit_ToString(s.instance)
}

func (s *TSpinEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    SpinEdit_AnchorToNeighbour(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (s *TSpinEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    SpinEdit_AnchorParallel(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (s *TSpinEdit) AnchorHorizontalCenterTo(ASibling IControl) {
    SpinEdit_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (s *TSpinEdit) AnchorVerticalCenterTo(ASibling IControl) {
    SpinEdit_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TSpinEdit) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    SpinEdit_AnchorSame(s.instance, ASide , CheckPtr(ASibling))
}

func (s *TSpinEdit) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    SpinEdit_AnchorAsAlign(s.instance, ATheAlign , ASpace)
}

func (s *TSpinEdit) AnchorClient(ASpace int32) {
    SpinEdit_AnchorClient(s.instance, ASpace)
}

// 获取四个角位置的锚点。
func (s *TSpinEdit) Anchors() TAnchors {
    return SpinEdit_GetAnchors(s.instance)
}

// 设置四个角位置的锚点。
func (s *TSpinEdit) SetAnchors(value TAnchors) {
    SpinEdit_SetAnchors(s.instance, value)
}

// 获取自动选择。
func (s *TSpinEdit) AutoSelect() bool {
    return SpinEdit_GetAutoSelect(s.instance)
}

// 设置自动选择。
func (s *TSpinEdit) SetAutoSelect(value bool) {
    SpinEdit_SetAutoSelect(s.instance, value)
}

// 获取自动调整大小。
func (s *TSpinEdit) AutoSize() bool {
    return SpinEdit_GetAutoSize(s.instance)
}

// 设置自动调整大小。
func (s *TSpinEdit) SetAutoSize(value bool) {
    SpinEdit_SetAutoSize(s.instance, value)
}

// 获取颜色。
//
// Get color.
func (s *TSpinEdit) Color() TColor {
    return SpinEdit_GetColor(s.instance)
}

// 设置颜色。
//
// Set color.
func (s *TSpinEdit) SetColor(value TColor) {
    SpinEdit_SetColor(s.instance, value)
}

// 获取约束控件大小。
func (s *TSpinEdit) Constraints() *TSizeConstraints {
    return AsSizeConstraints(SpinEdit_GetConstraints(s.instance))
}

// 设置约束控件大小。
func (s *TSpinEdit) SetConstraints(value *TSizeConstraints) {
    SpinEdit_SetConstraints(s.instance, CheckPtr(value))
}

// 获取控件启用。
//
// Get the control enabled.
func (s *TSpinEdit) Enabled() bool {
    return SpinEdit_GetEnabled(s.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (s *TSpinEdit) SetEnabled(value bool) {
    SpinEdit_SetEnabled(s.instance, value)
}

// 获取字体。
//
// Get Font.
func (s *TSpinEdit) Font() *TFont {
    return AsFont(SpinEdit_GetFont(s.instance))
}

// 设置字体。
//
// Set Font.
func (s *TSpinEdit) SetFont(value *TFont) {
    SpinEdit_SetFont(s.instance, CheckPtr(value))
}

func (s *TSpinEdit) Increment() int32 {
    return SpinEdit_GetIncrement(s.instance)
}

func (s *TSpinEdit) SetIncrement(value int32) {
    SpinEdit_SetIncrement(s.instance, value)
}

// 获取最大长度。
func (s *TSpinEdit) MaxLength() int32 {
    return SpinEdit_GetMaxLength(s.instance)
}

// 设置最大长度。
func (s *TSpinEdit) SetMaxLength(value int32) {
    SpinEdit_SetMaxLength(s.instance, value)
}

func (s *TSpinEdit) MaxValue() int32 {
    return SpinEdit_GetMaxValue(s.instance)
}

func (s *TSpinEdit) SetMaxValue(value int32) {
    SpinEdit_SetMaxValue(s.instance, value)
}

func (s *TSpinEdit) MinValue() int32 {
    return SpinEdit_GetMinValue(s.instance)
}

func (s *TSpinEdit) SetMinValue(value int32) {
    SpinEdit_SetMinValue(s.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (s *TSpinEdit) ParentColor() bool {
    return SpinEdit_GetParentColor(s.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (s *TSpinEdit) SetParentColor(value bool) {
    SpinEdit_SetParentColor(s.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (s *TSpinEdit) ParentFont() bool {
    return SpinEdit_GetParentFont(s.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (s *TSpinEdit) SetParentFont(value bool) {
    SpinEdit_SetParentFont(s.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (s *TSpinEdit) ParentShowHint() bool {
    return SpinEdit_GetParentShowHint(s.instance)
}

// 设置以父容器的ShowHint属性为准。
func (s *TSpinEdit) SetParentShowHint(value bool) {
    SpinEdit_SetParentShowHint(s.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (s *TSpinEdit) PopupMenu() *TPopupMenu {
    return AsPopupMenu(SpinEdit_GetPopupMenu(s.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (s *TSpinEdit) SetPopupMenu(value IComponent) {
    SpinEdit_SetPopupMenu(s.instance, CheckPtr(value))
}

// 获取只读。
func (s *TSpinEdit) ReadOnly() bool {
    return SpinEdit_GetReadOnly(s.instance)
}

// 设置只读。
func (s *TSpinEdit) SetReadOnly(value bool) {
    SpinEdit_SetReadOnly(s.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (s *TSpinEdit) ShowHint() bool {
    return SpinEdit_GetShowHint(s.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (s *TSpinEdit) SetShowHint(value bool) {
    SpinEdit_SetShowHint(s.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (s *TSpinEdit) TabOrder() TTabOrder {
    return SpinEdit_GetTabOrder(s.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (s *TSpinEdit) SetTabOrder(value TTabOrder) {
    SpinEdit_SetTabOrder(s.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (s *TSpinEdit) TabStop() bool {
    return SpinEdit_GetTabStop(s.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (s *TSpinEdit) SetTabStop(value bool) {
    SpinEdit_SetTabStop(s.instance, value)
}

func (s *TSpinEdit) Value() int32 {
    return SpinEdit_GetValue(s.instance)
}

func (s *TSpinEdit) SetValue(value int32) {
    SpinEdit_SetValue(s.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (s *TSpinEdit) Visible() bool {
    return SpinEdit_GetVisible(s.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (s *TSpinEdit) SetVisible(value bool) {
    SpinEdit_SetVisible(s.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (s *TSpinEdit) SetOnChange(fn TNotifyEvent) {
    SpinEdit_SetOnChange(s.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (s *TSpinEdit) SetOnClick(fn TNotifyEvent) {
    SpinEdit_SetOnClick(s.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (s *TSpinEdit) SetOnEnter(fn TNotifyEvent) {
    SpinEdit_SetOnEnter(s.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (s *TSpinEdit) SetOnExit(fn TNotifyEvent) {
    SpinEdit_SetOnExit(s.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (s *TSpinEdit) SetOnKeyDown(fn TKeyEvent) {
    SpinEdit_SetOnKeyDown(s.instance, fn)
}

// 设置键键下事件。
func (s *TSpinEdit) SetOnKeyPress(fn TKeyPressEvent) {
    SpinEdit_SetOnKeyPress(s.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (s *TSpinEdit) SetOnKeyUp(fn TKeyEvent) {
    SpinEdit_SetOnKeyUp(s.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (s *TSpinEdit) SetOnMouseDown(fn TMouseEvent) {
    SpinEdit_SetOnMouseDown(s.instance, fn)
}

// 设置鼠标移动事件。
func (s *TSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    SpinEdit_SetOnMouseMove(s.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (s *TSpinEdit) SetOnMouseUp(fn TMouseEvent) {
    SpinEdit_SetOnMouseUp(s.instance, fn)
}

// 获取文字对齐。
//
// Get Text alignment.
func (s *TSpinEdit) Alignment() TAlignment {
    return SpinEdit_GetAlignment(s.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (s *TSpinEdit) SetAlignment(value TAlignment) {
    SpinEdit_SetAlignment(s.instance, value)
}

// 获取能否撤销。
func (s *TSpinEdit) CanUndo() bool {
    return SpinEdit_GetCanUndo(s.instance)
}

// 获取修改。
//
// Get modified.
func (s *TSpinEdit) Modified() bool {
    return SpinEdit_GetModified(s.instance)
}

// 设置修改。
//
// Set modified.
func (s *TSpinEdit) SetModified(value bool) {
    SpinEdit_SetModified(s.instance, value)
}

// 获取选择的长度。
func (s *TSpinEdit) SelLength() int32 {
    return SpinEdit_GetSelLength(s.instance)
}

// 设置选择的长度。
func (s *TSpinEdit) SetSelLength(value int32) {
    SpinEdit_SetSelLength(s.instance, value)
}

// 获取选择的启始位置。
func (s *TSpinEdit) SelStart() int32 {
    return SpinEdit_GetSelStart(s.instance)
}

// 设置选择的启始位置。
func (s *TSpinEdit) SetSelStart(value int32) {
    SpinEdit_SetSelStart(s.instance, value)
}

// 获取选择的文本。
func (s *TSpinEdit) SelText() string {
    return SpinEdit_GetSelText(s.instance)
}

// 设置选择的文本。
func (s *TSpinEdit) SetSelText(value string) {
    SpinEdit_SetSelText(s.instance, value)
}

// 获取文本。
func (s *TSpinEdit) Text() string {
    strLen := s.GetTextLen()
    if strLen != 0 {
        var buffStr string
        s.GetTextBuf(&buffStr, strLen + 1)
        return buffStr
    }
    return ""
}

// 设置文本。
func (s *TSpinEdit) SetText(value string) {
    SpinEdit_SetText(s.instance, value)
}

// 获取提示文本。
func (s *TSpinEdit) TextHint() string {
    return SpinEdit_GetTextHint(s.instance)
}

// 设置提示文本。
func (s *TSpinEdit) SetTextHint(value string) {
    SpinEdit_SetTextHint(s.instance, value)
}

// 获取依靠客户端总数。
func (s *TSpinEdit) DockClientCount() int32 {
    return SpinEdit_GetDockClientCount(s.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (s *TSpinEdit) DockSite() bool {
    return SpinEdit_GetDockSite(s.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (s *TSpinEdit) SetDockSite(value bool) {
    SpinEdit_SetDockSite(s.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (s *TSpinEdit) DoubleBuffered() bool {
    return SpinEdit_GetDoubleBuffered(s.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (s *TSpinEdit) SetDoubleBuffered(value bool) {
    SpinEdit_SetDoubleBuffered(s.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (s *TSpinEdit) MouseInClient() bool {
    return SpinEdit_GetMouseInClient(s.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (s *TSpinEdit) VisibleDockClientCount() int32 {
    return SpinEdit_GetVisibleDockClientCount(s.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (s *TSpinEdit) Brush() *TBrush {
    return AsBrush(SpinEdit_GetBrush(s.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (s *TSpinEdit) ControlCount() int32 {
    return SpinEdit_GetControlCount(s.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (s *TSpinEdit) Handle() HWND {
    return SpinEdit_GetHandle(s.instance)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (s *TSpinEdit) ParentDoubleBuffered() bool {
    return SpinEdit_GetParentDoubleBuffered(s.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (s *TSpinEdit) SetParentDoubleBuffered(value bool) {
    SpinEdit_SetParentDoubleBuffered(s.instance, value)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (s *TSpinEdit) ParentWindow() HWND {
    return SpinEdit_GetParentWindow(s.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (s *TSpinEdit) SetParentWindow(value HWND) {
    SpinEdit_SetParentWindow(s.instance, value)
}

func (s *TSpinEdit) Showing() bool {
    return SpinEdit_GetShowing(s.instance)
}

// 获取使用停靠管理。
func (s *TSpinEdit) UseDockManager() bool {
    return SpinEdit_GetUseDockManager(s.instance)
}

// 设置使用停靠管理。
func (s *TSpinEdit) SetUseDockManager(value bool) {
    SpinEdit_SetUseDockManager(s.instance, value)
}

func (s *TSpinEdit) Action() *TAction {
    return AsAction(SpinEdit_GetAction(s.instance))
}

func (s *TSpinEdit) SetAction(value IComponent) {
    SpinEdit_SetAction(s.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (s *TSpinEdit) Align() TAlign {
    return SpinEdit_GetAlign(s.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (s *TSpinEdit) SetAlign(value TAlign) {
    SpinEdit_SetAlign(s.instance, value)
}

func (s *TSpinEdit) BiDiMode() TBiDiMode {
    return SpinEdit_GetBiDiMode(s.instance)
}

func (s *TSpinEdit) SetBiDiMode(value TBiDiMode) {
    SpinEdit_SetBiDiMode(s.instance, value)
}

func (s *TSpinEdit) BoundsRect() TRect {
    return SpinEdit_GetBoundsRect(s.instance)
}

func (s *TSpinEdit) SetBoundsRect(value TRect) {
    SpinEdit_SetBoundsRect(s.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (s *TSpinEdit) ClientHeight() int32 {
    return SpinEdit_GetClientHeight(s.instance)
}

// 设置客户区高度。
//
// Set client height.
func (s *TSpinEdit) SetClientHeight(value int32) {
    SpinEdit_SetClientHeight(s.instance, value)
}

func (s *TSpinEdit) ClientOrigin() TPoint {
    return SpinEdit_GetClientOrigin(s.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (s *TSpinEdit) ClientRect() TRect {
    return SpinEdit_GetClientRect(s.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (s *TSpinEdit) ClientWidth() int32 {
    return SpinEdit_GetClientWidth(s.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (s *TSpinEdit) SetClientWidth(value int32) {
    SpinEdit_SetClientWidth(s.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (s *TSpinEdit) ControlState() TControlState {
    return SpinEdit_GetControlState(s.instance)
}

// 设置控件状态。
//
// Set control state.
func (s *TSpinEdit) SetControlState(value TControlState) {
    SpinEdit_SetControlState(s.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (s *TSpinEdit) ControlStyle() TControlStyle {
    return SpinEdit_GetControlStyle(s.instance)
}

// 设置控件样式。
//
// Set control style.
func (s *TSpinEdit) SetControlStyle(value TControlStyle) {
    SpinEdit_SetControlStyle(s.instance, value)
}

func (s *TSpinEdit) Floating() bool {
    return SpinEdit_GetFloating(s.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (s *TSpinEdit) Parent() *TWinControl {
    return AsWinControl(SpinEdit_GetParent(s.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (s *TSpinEdit) SetParent(value IWinControl) {
    SpinEdit_SetParent(s.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (s *TSpinEdit) Left() int32 {
    return SpinEdit_GetLeft(s.instance)
}

// 设置左边位置。
//
// Set Left position.
func (s *TSpinEdit) SetLeft(value int32) {
    SpinEdit_SetLeft(s.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (s *TSpinEdit) Top() int32 {
    return SpinEdit_GetTop(s.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (s *TSpinEdit) SetTop(value int32) {
    SpinEdit_SetTop(s.instance, value)
}

// 获取宽度。
//
// Get width.
func (s *TSpinEdit) Width() int32 {
    return SpinEdit_GetWidth(s.instance)
}

// 设置宽度。
//
// Set width.
func (s *TSpinEdit) SetWidth(value int32) {
    SpinEdit_SetWidth(s.instance, value)
}

// 获取高度。
//
// Get height.
func (s *TSpinEdit) Height() int32 {
    return SpinEdit_GetHeight(s.instance)
}

// 设置高度。
//
// Set height.
func (s *TSpinEdit) SetHeight(value int32) {
    SpinEdit_SetHeight(s.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (s *TSpinEdit) Cursor() TCursor {
    return SpinEdit_GetCursor(s.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (s *TSpinEdit) SetCursor(value TCursor) {
    SpinEdit_SetCursor(s.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (s *TSpinEdit) Hint() string {
    return SpinEdit_GetHint(s.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (s *TSpinEdit) SetHint(value string) {
    SpinEdit_SetHint(s.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (s *TSpinEdit) ComponentCount() int32 {
    return SpinEdit_GetComponentCount(s.instance)
}

// 获取组件索引。
//
// Get component index.
func (s *TSpinEdit) ComponentIndex() int32 {
    return SpinEdit_GetComponentIndex(s.instance)
}

// 设置组件索引。
//
// Set component index.
func (s *TSpinEdit) SetComponentIndex(value int32) {
    SpinEdit_SetComponentIndex(s.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (s *TSpinEdit) Owner() *TComponent {
    return AsComponent(SpinEdit_GetOwner(s.instance))
}

// 获取组件名称。
//
// Get the component name.
func (s *TSpinEdit) Name() string {
    return SpinEdit_GetName(s.instance)
}

// 设置组件名称。
//
// Set the component name.
func (s *TSpinEdit) SetName(value string) {
    SpinEdit_SetName(s.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (s *TSpinEdit) Tag() int {
    return SpinEdit_GetTag(s.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (s *TSpinEdit) SetTag(value int) {
    SpinEdit_SetTag(s.instance, value)
}

// 获取左边锚点。
func (s *TSpinEdit) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(SpinEdit_GetAnchorSideLeft(s.instance))
}

// 设置左边锚点。
func (s *TSpinEdit) SetAnchorSideLeft(value *TAnchorSide) {
    SpinEdit_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (s *TSpinEdit) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(SpinEdit_GetAnchorSideTop(s.instance))
}

// 设置顶边锚点。
func (s *TSpinEdit) SetAnchorSideTop(value *TAnchorSide) {
    SpinEdit_SetAnchorSideTop(s.instance, CheckPtr(value))
}

// 获取右边锚点。
func (s *TSpinEdit) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(SpinEdit_GetAnchorSideRight(s.instance))
}

// 设置右边锚点。
func (s *TSpinEdit) SetAnchorSideRight(value *TAnchorSide) {
    SpinEdit_SetAnchorSideRight(s.instance, CheckPtr(value))
}

// 获取底边锚点。
func (s *TSpinEdit) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(SpinEdit_GetAnchorSideBottom(s.instance))
}

// 设置底边锚点。
func (s *TSpinEdit) SetAnchorSideBottom(value *TAnchorSide) {
    SpinEdit_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

func (s *TSpinEdit) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(SpinEdit_GetChildSizing(s.instance))
}

func (s *TSpinEdit) SetChildSizing(value *TControlChildSizing) {
    SpinEdit_SetChildSizing(s.instance, CheckPtr(value))
}

// 获取边框间距。
func (s *TSpinEdit) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(SpinEdit_GetBorderSpacing(s.instance))
}

// 设置边框间距。
func (s *TSpinEdit) SetBorderSpacing(value *TControlBorderSpacing) {
    SpinEdit_SetBorderSpacing(s.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (s *TSpinEdit) DockClients(Index int32) *TControl {
    return AsControl(SpinEdit_GetDockClients(s.instance, Index))
}

// 获取指定索引子控件。
func (s *TSpinEdit) Controls(Index int32) *TControl {
    return AsControl(SpinEdit_GetControls(s.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (s *TSpinEdit) Components(AIndex int32) *TComponent {
    return AsComponent(SpinEdit_GetComponents(s.instance, AIndex))
}

// 获取锚侧面。
func (s *TSpinEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(SpinEdit_GetAnchorSide(s.instance, AKind))
}

